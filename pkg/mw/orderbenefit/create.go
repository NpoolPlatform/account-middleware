package orderbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/account-middleware/pkg/mw/account"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	pbaccount "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

//nolint:gocyclo
func (h *Handler) checkBaseAccount(ctx context.Context) (exist bool, err error) {
	var baseAccount *pbaccount.Account

	if h.AccountID != nil {
		accountID := h.AccountID.String()
		accountH, err := account.NewHandler(ctx, account.WithEntID(&accountID, true))
		if err != nil {
			return false, err
		}
		baseAccount, err = accountH.GetAccount(ctx)
		if err != nil {
			return false, err
		}

		if baseAccount == nil {
			return false, fmt.Errorf("invalid accountid")
		}

		if baseAccount.UsedFor != *h.usedFor {
			return false, fmt.Errorf("invalid account usedfor")
		}

		if h.CoinTypeID != nil && baseAccount.CoinTypeID != h.CoinTypeID.String() {
			return false, fmt.Errorf("invalid cointypeid")
		}

		if h.Address != nil && baseAccount.Address != *h.Address {
			return false, fmt.Errorf("invalid address")
		}

		return true, nil
	} else if h.CoinTypeID == nil || h.Address == nil {
		return false, fmt.Errorf("invalid cointypeid or address")
	} else {
		id := uuid.New()
		h.AccountID = &id
	}

	return false, nil
}

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateUserAccount, *h.AppID, *h.UserID, *h.OrderID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	if h.EntID == nil {
		id := uuid.New()
		h.EntID = &id
	}

	createBaseAccount, err := h.checkBaseAccount(ctx)
	if err != nil {
		return nil, err
	}

	sqlH := h.newSQLHandler()

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if !createBaseAccount {
			if _, err := accountcrud.CreateSet(
				tx.Account.Create(),
				&accountcrud.Req{
					EntID:                  h.AccountID,
					CoinTypeID:             h.CoinTypeID,
					Address:                h.Address,
					UsedFor:                h.usedFor,
					PlatformHoldPrivateKey: h.platformHoldPrivateKey,
				},
			).Save(ctx); err != nil {
				return err
			}
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
			return fmt.Errorf("fail create account: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
