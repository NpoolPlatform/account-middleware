package orderbenefit

import (
	"context"
	"fmt"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entorderbenefit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/orderbenefit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		orderbenefit, err := tx.OrderBenefit.
			Query().
			Where(
				entorderbenefit.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if orderbenefit == nil {
			return fmt.Errorf("invalid orderbenefit")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(orderbenefit.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				Active:  h.Active,
				Locked:  h.Locked,
				Blocked: h.Blocked,
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
