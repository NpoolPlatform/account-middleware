package transfer

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	enttransfer "github.com/NpoolPlatform/account-middleware/pkg/db/ent/transfer"

	transfercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/transfer"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"
)

type queryHandler struct {
	*Handler
	stm   *ent.TransferSelect
	infos []*npool.Transfer
	total uint32
}

func (h *queryHandler) selectTransfer(stm *ent.TransferQuery) {
	h.stm = stm.Select(
		enttransfer.FieldID,
		enttransfer.FieldAppID,
		enttransfer.FieldUserID,
		enttransfer.FieldTargetUserID,
		enttransfer.FieldCreatedAt,
		enttransfer.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryTransfer(cli *ent.Client) {
	h.selectTransfer(
		cli.Transfer.
			Query().
			Where(
				enttransfer.ID(*h.ID),
				enttransfer.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryTransfers(cli *ent.Client) error {
	stm, err := transfercrud.SetQueryConds(cli.Transfer.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.selectTransfer(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTransfer(ctx context.Context) (*npool.Transfer, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryTransfer(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetTransfers(ctx context.Context) ([]*npool.Transfer, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfers(cli); err != nil {
			return err
		}

		_total, err := handler.stm.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(enttransfer.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetTransferOnly(ctx context.Context) (*npool.Transfer, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfers(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.
			Offset(0).
			Limit(singleRowLimit).
			Order(ent.Desc(enttransfer.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}
