package platform

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	platformcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/platform"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entplatform "github.com/NpoolPlatform/account-middleware/pkg/db/ent/platform"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		platform, err := tx.Platform.
			Query().
			Where(
				entplatform.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if platform == nil {
			return fmt.Errorf("invalid platform")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(platform.AccountID),
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

		if h.Backup == nil || !*h.Backup {
			return nil
		}

		ids, err := tx.
			Platform.
			Query().
			Select().
			Modify(func(s *sql.Selector) {
				t := sql.Table(entaccount.Table)
				s.LeftJoin(t).
					On(
						t.C(entaccount.FieldID),
						s.C(entplatform.FieldAccountID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldCoinTypeID), account.CoinTypeID),
					).
					OnP(
						sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
					)
				s.Where(
					sql.EQ(t.C(entaccount.FieldCoinTypeID), *h.CoinTypeID),
				)
			}).
			Where(
				entplatform.IDNEQ(*h.ID),
				entplatform.UsedFor(platform.UsedFor),
				entplatform.Backup(false),
			).
			IDs(_ctx)
		if err != nil {
			return err
		}

		if _, err := tx.
			Platform.
			Update().
			Where(
				entplatform.IDIn(ids...),
			).
			SetBackup(true).
			Save(_ctx); err != nil {
			return err
		}

		if _, err := platformcrud.UpdateSet(
			platform.Update(),
			&platformcrud.Req{
				Backup: h.Backup,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
