package orderbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/account-middleware/pkg/mw/account"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	"github.com/NpoolPlatform/account-middleware/pkg/crud/orderbenefit"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	pbaccount "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) queryBaseAccount(ctx context.Context) (*npool.Account, error) {
	h.Conds = &orderbenefit.Conds{}
	h.Conds.AppID = &cruder.Cond{
		Op:  cruder.EQ,
		Val: *h.AppID,
	}
	h.Conds.UserID = &cruder.Cond{
		Op:  cruder.EQ,
		Val: *h.UserID,
	}
	h.Conds.UsedFor = &cruder.Cond{
		Op:  cruder.EQ,
		Val: *h.usedFor,
	}
	h.Conds.CoinTypeID = &cruder.Cond{
		Op:  cruder.EQ,
		Val: *h.CoinTypeID,
	}
	h.Conds.Address = &cruder.Cond{
		Op:  cruder.EQ,
		Val: *h.Address,
	}

	h.Limit = 1
	h.Offset = 0

	handler := queryHandler{Handler: h}
	accInfos, _, err := handler.getAccounts(ctx)
	if err != nil || len(accInfos) == 0 {
		return nil, err
	}
	return accInfos[0], nil
}

//nolint:gocyclo
func (h *Handler) checkBaseAccount(ctx context.Context) (exist bool, err error) {
	var baseAccount *pbaccount.Account

	if h.AccountID == nil {
		id := uuid.New()
		h.AccountID = &id
	} else {
		accountID := h.AccountID.String()
		accountH, err := account.NewHandler(ctx, account.WithEntID(&accountID, true))
		if err != nil {
			return false, err
		}
		baseAccount, err = accountH.GetAccount(ctx)
		if err != nil {
			return false, err
		}

		if baseAccount != nil && baseAccount.UsedFor != *h.usedFor {
			return false, fmt.Errorf("invalid account usedfor")
		}
	}

	if baseAccount == nil && (h.CoinTypeID == nil || h.Address == nil) {
		return false, fmt.Errorf("invalid cointypeid or address")
	}

	if baseAccount != nil {
		if h.CoinTypeID != nil && baseAccount.CoinTypeID != h.CoinTypeID.String() {
			return false, fmt.Errorf("invalid cointypeid")
		} else if h.CoinTypeID == nil {
			cointypeID, err := uuid.Parse(baseAccount.CoinTypeID)
			if err != nil {
				return false, err
			}
			h.CoinTypeID = &cointypeID
		}

		if h.Address != nil && baseAccount.Address != *h.Address {
			return false, fmt.Errorf("invalid address")
		} else if h.Address == nil {
			h.Address = &baseAccount.Address
		}
		return true, nil
	} else {
		historyAccount, err := h.queryBaseAccount(ctx)
		if err != nil {
			return false, err
		}
		if historyAccount != nil {
			accountID, err := uuid.Parse(historyAccount.AccountID)
			if err != nil {
				return false, err
			}
			h.AccountID = &accountID
			return true, nil
		}
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
