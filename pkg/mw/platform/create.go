package platform

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	platformcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/platform"
	account1 "github.com/NpoolPlatform/account-middleware/pkg/mw/account"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	accountmwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
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
	if h.UsedFor == nil {
		return nil, fmt.Errorf("invalid usedfor")
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreatePlatformAccount, *h.CoinTypeID, *h.Address)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler, err := account1.NewHandler(
		ctx,
		account1.WithConds(&accountmwpb.Conds{
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID.String()},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.Address},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistAccountConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("address exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	id = uuid.New()
	if h.AccountID == nil {
		h.AccountID = &id
	}

	privateKey := true
	switch *h.UsedFor {
	case basetypes.AccountUsedFor_UserBenefitHot:
	case basetypes.AccountUsedFor_UserBenefitCold:
		privateKey = false
	case basetypes.AccountUsedFor_PlatformBenefitCold:
		privateKey = false
	case basetypes.AccountUsedFor_GasProvider:
	case basetypes.AccountUsedFor_PaymentCollector:
		privateKey = false
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				ID:                     h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                h.UsedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(ctx); err != nil {
			return err
		}

		if _, err := platformcrud.CreateSet(
			tx.Platform.Create(),
			&platformcrud.Req{
				ID:        h.ID,
				UsedFor:   h.UsedFor,
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
