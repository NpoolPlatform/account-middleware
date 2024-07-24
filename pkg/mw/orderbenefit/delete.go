package orderbenefit

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	orderbenefitcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/orderbenefit"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entorderbenefit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/orderbenefit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
)

func (h *Handler) DeleteAccount(ctx context.Context) (*npool.Account, error) {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
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
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := orderbenefitcrud.UpdateSet(
			orderbenefit.Update(),
			&orderbenefitcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
