package account

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
)

type existHandler struct {
	*Handler
	stm   *ent.AccountQuery
	infos []*npool.Account
	total uint32
}

func (h *existHandler) queryAccount(cli *ent.Client) {
	h.stm = cli.Account.
		Query().
		Where(
			entaccount.ID(*h.ID),
			entaccount.DeletedAt(0),
		)
}

func (h *existHandler) queryAccounts(ctx context.Context, cli *ent.Client) error {
	stm, err := accountcrud.SetQueryConds(cli.Account.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistAccount(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAccount(cli)
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

func (h *Handler) ExistAccountConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(_ctx, cli); err != nil {
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
