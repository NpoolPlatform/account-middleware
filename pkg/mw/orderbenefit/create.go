package orderbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/account-middleware/pkg/mw/account"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	pbaccount "github.com/NpoolPlatform/message/npool/account/mw/v1/account"

	"github.com/google/uuid"
)

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

		if baseAccount.UsedFor != *h.accountReq.UsedFor {
			return false, fmt.Errorf("invalid account usedfor")
		}

		if h.CoinTypeID != nil && baseAccount.CoinTypeID != h.CoinTypeID.String() {
			return false, fmt.Errorf("invalid cointypeid")
		}

		if h.accountReq.Address != nil && baseAccount.Address != *h.accountReq.Address {
			return false, fmt.Errorf("invalid address")
		}

		return true, nil
	} else if h.CoinTypeID == nil || h.accountReq.Address == nil {
		return false, fmt.Errorf("invalid cointypeid or address")
	} else {
		id := uuid.New()
		h.AccountID = &id
	}

	return false, nil
}

func (h *Handler) CreateAccountWithTx(ctx context.Context, tx *ent.Tx) error {
	if h.EntID == nil {
		id := uuid.New()
		h.EntID = &id
	}

	accountExist, err := h.checkBaseAccount(ctx)
	if err != nil {
		return err
	}

	sqlH := h.newSQLHandler()

	if !accountExist {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				EntID:                  h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.accountReq.Address,
				UsedFor:                h.accountReq.UsedFor,
				PlatformHoldPrivateKey: h.accountReq.PlatformHoldPrivateKey,
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
}

func (h *Handler) CreateAccount(ctx context.Context) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return h.CreateAccountWithTx(ctx, tx)
	})
}
