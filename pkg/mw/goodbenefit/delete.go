package goodbenefit

import (
	"context"
	"time"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/goodbenefit"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entgoodbenefit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/goodbenefit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
)

func (h *Handler) DeleteAccount(ctx context.Context) (*npool.Account, error) {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	now := uint32(time.Now().Unix())
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
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

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(goodbenefit.AccountID),
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

		if _, err := goodbenefitcrud.UpdateSet(
			goodbenefit.Update(),
			&goodbenefitcrud.Req{
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
