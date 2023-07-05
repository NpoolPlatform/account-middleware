package deposit

import (
	"context"
	"fmt"
	"time"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	depositcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/deposit"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entdeposit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/deposit"
	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

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

		var _scannableAt *uint32
		if account.Locked && h.Locked != nil && !*h.Locked {
			scannableAt := uint32(time.Now().Unix()) + timedef.SecondsPerHour
			_scannableAt = &scannableAt
		}

		incoming := deposit.Incoming
		if h.Incoming != nil {
			incoming = incoming.Add(*h.Incoming)
		}
		outcoming := deposit.Outcoming
		if h.Outcoming != nil {
			outcoming = outcoming.Add(*h.Outcoming)
		}

		if incoming.Cmp(outcoming) < 0 {
			return fmt.Errorf("incoming (%v) < outcoming (%v)", incoming, outcoming)
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
				Incoming:      &incoming,
				Outcoming:     &outcoming,
				ScannableAt:   _scannableAt,
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
