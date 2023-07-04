package transfer

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	transfercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/transfer"
	enttransfer "github.com/NpoolPlatform/account-middleware/pkg/db/ent/transfer"
)

type existHandler struct {
	*Handler
	stm *ent.TransferQuery
}

func (h *existHandler) queryTransfer(cli *ent.Client) {
	h.stm = cli.Transfer.
		Query().
		Where(
			enttransfer.ID(*h.ID),
			enttransfer.DeletedAt(0),
		)
}

func (h *existHandler) queryTransfers(cli *ent.Client) error {
	stm, err := transfercrud.SetQueryConds(cli.Transfer.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistTransfer(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryTransfer(cli)
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistTransferConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfers(cli); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
