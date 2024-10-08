package orderbenefit

import (
	"context"
	"fmt"
	"time"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	"github.com/google/uuid"

	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
)

func (h *Handler) DeleteAccountWithTx(ctx context.Context, tx *ent.Tx) error {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return nil
	}

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
	accountID, err := uuid.Parse(info.AccountID)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	now := uint32(time.Now().Unix())
	account, err := tx.Account.
		Query().
		Where(
			entaccount.EntID(accountID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	_, err = accountcrud.UpdateSet(
		account.Update(),
		&accountcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx)

	return err
}

func (h *Handler) DeleteAccount(ctx context.Context) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return h.DeleteAccountWithTx(ctx, tx)
	})
}
