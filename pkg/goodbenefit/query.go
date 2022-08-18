package goodbenefit

import (
	"context"
	goodbenefitcurlcrud "github.com/NpoolPlatform/account-manager/pkg/crud/goodbenefit"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/goodbenefit"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

	"github.com/google/uuid"
)

func GetAccount(ctx context.Context, id string) (info *npool.Account, err error) {
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			GoodBenefit.
			Query().
			Where(
				goodbenefit.ID(uuid.MustParse(id)),
			).
			Select(
				goodbenefit.FieldID,
				goodbenefit.FieldGoodID,
				goodbenefit.FieldBackup,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(account.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(deposit.FieldAccountID),
						t1.C(account.FieldID),
					).
					AppendSelect(
						sql.As(t1.C(account.FieldID), "account_id"),
						sql.As(t1.C(account.FieldAddress), "address"),
						sql.As(t1.C(account.FieldActive), "active"),
						sql.As(t1.C(account.FieldLocked), "locked"),
						sql.As(t1.C(account.FieldLockedBy), "locked_by"),
						sql.As(t1.C(account.FieldBlocked), "blocked"),
						sql.As(t1.C(account.FieldUsedFor), "used_for"),
					)
			}).
			Scan(ctx, &info)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetAccounts(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Account, err error) {
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := goodbenefitcurlcrud.SetQueryConds(&mgrpb.Conds{
			ID:        conds.ID,
			GoodID:    conds.GoodID,
			AccountID: conds.AccountID,
			Backup:    conds.Backup,
		}, cli)
		if err != nil {
			return err
		}

		stm.Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func join(stm *ent.GoodBenefitQuery) *ent.GoodBenefitSelect {
	return stm.Select(
		goodbenefit.FieldID,
		goodbenefit.FieldGoodID,
		goodbenefit.FieldBackup,
	).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(account.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(deposit.FieldAccountID),
					t1.C(account.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(account.FieldID), "account_id"),
					sql.As(t1.C(account.FieldAddress), "address"),
					sql.As(t1.C(account.FieldActive), "active"),
					sql.As(t1.C(account.FieldLocked), "locked"),
					sql.As(t1.C(account.FieldLockedBy), "locked_by"),
					sql.As(t1.C(account.FieldBlocked), "blocked"),
					sql.As(t1.C(account.FieldUsedFor), "used_for"),
				)
		})
}
