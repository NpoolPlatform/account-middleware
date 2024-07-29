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
	pbaccount "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"

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

		if baseAccount.UsedFor != *h.UsedFor {
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
					UsedFor:                h.UsedFor,
					PlatformHoldPrivateKey: h.PlatformHoldPrivateKey,
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

func (h *Handler) CreateAccounts(ctx context.Context) ([]*npool.Account, error) {
	entIDs := []uuid.UUID{}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			h.baseReq = *req

			if h.EntID == nil {
				id := uuid.New()
				h.EntID = &id
			}

			createBaseAccount, err := h.checkBaseAccount(ctx)
			if err != nil {
				return err
			}

			sqlH := h.newSQLHandler()
			if !createBaseAccount {
				if _, err := accountcrud.CreateSet(
					tx.Account.Create(),
					&accountcrud.Req{
						EntID:                  h.AccountID,
						CoinTypeID:             h.CoinTypeID,
						Address:                h.Address,
						UsedFor:                h.UsedFor,
						PlatformHoldPrivateKey: h.PlatformHoldPrivateKey,
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
				return fmt.Errorf("fail create accounts: %v", err)
			}

			entIDs = append(entIDs, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &orderbenefit.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: entIDs},
	}
	h.Offset = 0
	h.Limit = int32(len(entIDs))

	infos, _, err := h.GetAccounts(ctx)
	return infos, err
}
