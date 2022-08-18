package user

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"

	depositcrud "github.com/NpoolPlatform/account-manager/pkg/crud/deposit"

	depositmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func GetAccount(ctx context.Context, id string) (info *npool.Account, err error) {
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			Deposit.
			Query().
			Where(
				deposit.ID(uuid.MustParse(id)),
			).
			Select(
				deposit.FieldID,
				deposit.FieldAppID,
				deposit.FieldUserID,
				deposit.FieldCoinTypeID,
				deposit.FieldAccountID,
				deposit.FieldCollectingTid,
				deposit.FieldCreatedAt,
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
		stm, err := depositcrud.SetQueryConds(&depositmgrpb.Conds{
			ID:         conds.ID,
			AppID:      conds.AppID,
			UserID:     conds.UserID,
			CoinTypeID: conds.CoinTypeID,
			AccountID:  conds.AccountID,
		}, cli)
		if err != nil {
			return err
		}

		return stm.
			Select(
				deposit.FieldID,
				deposit.FieldAppID,
				deposit.FieldUserID,
				deposit.FieldCoinTypeID,
				deposit.FieldAccountID,
				deposit.FieldCollectingTid,
				deposit.FieldCreatedAt,
			).
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(account.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(deposit.FieldAccountID),
						t1.C(account.FieldID),
					)

				if conds.Address != nil && conds.GetAddress().GetOp() == cruder.EQ {
					s.Where(
						sql.EQ(
							t1.C(account.FieldAddress),
							conds.GetAddress().GetValue(),
						),
					)
				}
				if conds.Active != nil && conds.GetActive().GetOp() == cruder.EQ {
					s.Where(
						sql.EQ(
							t1.C(account.FieldActive),
							conds.GetActive().GetValue(),
						),
					)
				}
				if conds.Locked != nil && conds.GetLocked().GetOp() == cruder.EQ {
					s.Where(
						sql.EQ(
							t1.C(account.FieldLocked),
							conds.GetLocked().GetValue(),
						),
					)
				}
				if conds.LockedBy != nil && conds.GetLockedBy().GetOp() == cruder.EQ {
					s.Where(
						sql.EQ(
							t1.C(account.FieldLockedBy),
							conds.GetLockedBy().GetValue(),
						),
					)
				}
				if conds.Blocked != nil && conds.GetBlocked().GetOp() == cruder.EQ {
					s.Where(
						sql.EQ(
							t1.C(account.FieldBlocked),
							conds.GetBlocked().GetValue(),
						),
					)
				}

				s.
					AppendSelect(
						sql.As(t1.C(account.FieldAddress), "address"),
						sql.As(t1.C(account.FieldActive), "active"),
						sql.As(t1.C(account.FieldLocked), "locked"),
						sql.As(t1.C(account.FieldLockedBy), "locked_by"),
						sql.As(t1.C(account.FieldBlocked), "blocked"),
						sql.As(t1.C(account.FieldUsedFor), "used_for"),
					)
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
