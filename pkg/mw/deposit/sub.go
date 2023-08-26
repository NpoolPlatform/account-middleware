package deposit

import (
	"context"
	"fmt"

	depositcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/deposit"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entdeposit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/deposit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
)

func (h *Handler) SubBalance(ctx context.Context) (*npool.Account, error) {
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

		if _, err := depositcrud.UpdateSet(
			deposit.Update(),
			&depositcrud.Req{
				Incoming:  &incoming,
				Outcoming: &outcoming,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
