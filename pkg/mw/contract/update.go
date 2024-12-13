package contract

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	contractcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/contract"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entcontract "github.com/NpoolPlatform/account-middleware/pkg/db/ent/contract"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/contract"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) { //nolint
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		contract, err := tx.Contract.
			Query().
			Where(
				entcontract.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if contract == nil {
			return fmt.Errorf("invalid contract")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(contract.AccountID),
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

		if _, err := contractcrud.UpdateSet(
			contract.Update(),
			&contractcrud.Req{
				Backup:        h.Backup,
				TransactionID: h.TransactionID,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if h.Backup != nil && *h.Backup {
			return nil
		}

		ids, err := tx.
			Contract.
			Query().
			Select().
			Modify(func(s *sql.Selector) {
				t := sql.Table(entaccount.Table)
				s.LeftJoin(t).
					On(
						t.C(entaccount.FieldEntID),
						s.C(entcontract.FieldAccountID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldCoinTypeID), account.CoinTypeID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
					)
				s.Where(
					sql.EQ(t.C(entaccount.FieldCoinTypeID), account.CoinTypeID),
				)
			}).
			Where(
				entcontract.GoodID(contract.GoodID),
				entcontract.IDNEQ(*h.ID),
				entcontract.Backup(false),
			).
			IDs(_ctx)
		if err != nil {
			return err
		}

		if _, err := tx.
			Contract.
			Update().
			Where(
				entcontract.IDIn(ids...),
			).
			SetBackup(true).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
