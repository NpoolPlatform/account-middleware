package orderbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) {
	key := fmt.Sprintf("%v:%v:%v:%v:%v", basetypes.Prefix_PrefixCreateUserAccount, *h.AppID, *h.UserID, *h.CoinTypeID, *h.Address)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id1 := uuid.New()
	if h.EntID == nil {
		h.EntID = &id1
	}

	id2 := uuid.New()
	if h.AccountID == nil {
		h.AccountID = &id2
	}

	privateKey := false
	usedFor := basetypes.AccountUsedFor_OrderBenefit

	sqlH := h.newSQLHandler()

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				EntID:                  h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                &usedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(ctx); err != nil {
			return err
		}

		sql, err := sqlH.genCreateSQL()
		if err != nil {
			return err
		}

		rc, err := tx.ExecContext(ctx, sql)
		if err != nil {
			return err
		}
		if n, err := rc.RowsAffected(); err != nil || n != 1 {
			return fmt.Errorf("fail create gooduser: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
