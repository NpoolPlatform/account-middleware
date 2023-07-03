package deposit

import (
	"context"
	"fmt"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	depositcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/deposit"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entdeposit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/deposit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		deposit, err := tx.Deposit.
			Query().
			Where(
				entdeposit.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if deposit == nil {
			return fmt.Errorf("invalid deposit")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(deposit.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				Active:   h.Active,
				Locked:   h.Locked,
				LockedBy: h.LockedBy,
				Blocked:  h.Blocked,
			},
		).Save(_ctx); err != nil {
			return err
		}

		u, err := depositcrud.UpdateSet(
			deposit.Update(),
			&depositcrud.Req{
				CollectingTID: h.CollectingTID,
				Incoming:      h.Incoming,
				Outcoming:     h.Outcoming,
				ScannableAt:   h.ScannableAt,
			},
		)
		if err != nil {
			return err
		}

		if _, err := u.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
