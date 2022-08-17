package user

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

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

func GetAccounts(ctx context.Context, conds *npool.Conds) (infos []*npool.Account, err error) {
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {

	})
}
