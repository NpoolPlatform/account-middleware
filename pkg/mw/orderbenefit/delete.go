package orderbenefit

import (
	"context"
	"fmt"
	"time"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	"github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"

	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAccountBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	account, err := tx.Account.
		Query().
		Where(
			entaccount.EntID(*h.AccountID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err := accountcrud.UpdateSet(
		account.Update(),
		&accountcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return err
}

func (h *Handler) checkAccountInfo(info *orderbenefit.Account) error {
	if h.AppID != nil && h.AppID.String() != info.AppID {
		return fmt.Errorf("invalid appid")
	}
	if h.UserID != nil && h.UserID.String() != info.UserID {
		return fmt.Errorf("invalid userid")
	}
	if h.OrderID != nil && h.OrderID.String() != info.OrderID {
		return fmt.Errorf("invalid orderid")
	}
	if h.AccountID != nil && h.AccountID.String() != info.AccountID {
		return fmt.Errorf("invalid accountid")
	}
	return nil
}

func (h *Handler) DeleteAccountWithTx(ctx context.Context, tx *ent.Tx) error {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return nil
	}

	if err := h.checkAccountInfo(info); err != nil {
		return err
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
	}

	if err := handler.deleteAccountBase(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteAccount(ctx context.Context) error {
	return db.WithTx(ctx, h.DeleteAccountWithTx)
}
