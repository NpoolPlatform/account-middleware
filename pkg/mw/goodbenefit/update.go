package goodbenefit

import (
	"context"
	"fmt"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/goodbenefit"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entgoodbenefit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/goodbenefit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		goodbenefit, err := tx.GoodBenefit.
			Query().
			Where(
				entgoodbenefit.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if goodbenefit == nil {
			return fmt.Errorf("invalid goodbenefit")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(goodbenefit.AccountID),
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

		if _, err := goodbenefitcrud.UpdateSet(
			goodbenefit.Update(),
			&goodbenefitcrud.Req{
				Backup:        h.Backup,
				TransactionID: h.TransactionID,
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