package goodbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/goodbenefit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) {
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}
	if h.Address == nil {
		return nil, fmt.Errorf("invalid address")
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateGoodBenefitAccount, *h.CoinTypeID, *h.Address)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	id = uuid.New()
	if h.AccountID == nil {
		h.AccountID = &id
	}

	usedFor := basetypes.AccountUsedFor_GoodBenefit
	privateKey := true

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				ID:                     h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                &usedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(ctx); err != nil {
			return err
		}

		if _, err := goodbenefitcrud.CreateSet(
			tx.GoodBenefit.Create(),
			&goodbenefitcrud.Req{
				ID:        h.ID,
				GoodID:    h.GoodID,
				AccountID: h.AccountID,
				Backup:    h.Backup,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}