package payment

import (
	"context"
	"fmt"
	"time"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	paymentcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entpayment "github.com/NpoolPlatform/account-middleware/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		payment, err := tx.Payment.
			Query().
			Where(
				entpayment.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if payment == nil {
			return fmt.Errorf("invalid payment")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(payment.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if account.Locked && h.Locked != nil && !*h.Locked {
			const coolDown = uint32(60 * 60)
			availableAt := uint32(time.Now().Unix()) + coolDown
			h.AvailableAt = &availableAt
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

		u := paymentcrud.UpdateSet(
			payment.Update(),
			&paymentcrud.Req{
				AccountID:     h.AccountID,
				CollectingTID: h.CollectingTID,
				AvailableAt:   h.AvailableAt,
			},
		)

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
