package platform

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/account-manager/pkg/crud/platform"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	entplatform "github.com/NpoolPlatform/account-manager/pkg/db/ent/platform"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

	"github.com/google/uuid"
)

func GetAccount(ctx context.Context, id string) (info *npool.Account, err error) {
	var infos []*npool.Account

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Platform.
			Query().
			Where(
				entplatform.ID(uuid.MustParse(id)),
			)
		return join(stm, &npool.Conds{}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}

	infos = expand(infos)

	return infos[0], nil
}

func GetAccounts(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Account, total uint32, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:        conds.ID,
			AccountID: conds.AccountID,
			UsedFor:   conds.UsedFor,
			Backup:    conds.Backup,
		}, cli)
		if err != nil {
			return err
		}

		sel := join(stm, conds)

		_total, err := sel.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		return sel.
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, total, err
	}

	infos = expand(infos)

	return infos, total, nil
}

func GetAccountOnly(ctx context.Context, conds *npool.Conds) (info *npool.Account, err error) {
	infos := []*npool.Account{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:        conds.ID,
			AccountID: conds.AccountID,
			UsedFor:   conds.UsedFor,
			Backup:    conds.Backup,
		}, cli)
		if err != nil {
			return err
		}

		return join(stm, conds).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	infos = expand(infos)

	return infos[0], nil
}

func join(stm *ent.PlatformQuery, conds *npool.Conds) *ent.PlatformSelect {
	return stm.
		Modify(func(s *sql.Selector) {
			s.
				Select(
					s.C(entplatform.FieldID),
					s.C(entplatform.FieldBackup),
				)

			t1 := sql.Table(entaccount.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entplatform.FieldAccountID),
					t1.C(entaccount.FieldID),
				)

			if conds.CoinTypeID != nil {
				s.Where(
					sql.EQ(
						t1.C(entaccount.FieldCoinTypeID),
						conds.GetCoinTypeID().GetValue(),
					),
				)
			}
			if conds.Active != nil {
				s.Where(
					sql.EQ(
						t1.C(entaccount.FieldActive),
						conds.GetActive().GetValue(),
					),
				)
			}
			if conds.Locked != nil {
				s.Where(
					sql.EQ(
						t1.C(entaccount.FieldLocked),
						conds.GetLocked().GetValue(),
					),
				)
			}
			if conds.LockedBy != nil {
				s.Where(
					sql.EQ(
						t1.C(entaccount.FieldLockedBy),
						conds.GetLockedBy().GetValue(),
					),
				)
			}
			if conds.Blocked != nil {
				s.Where(
					sql.EQ(
						t1.C(entaccount.FieldBlocked),
						conds.GetBlocked().GetValue(),
					),
				)
			}

			s.
				AppendSelect(
					sql.As(t1.C(entaccount.FieldID), "account_id"),
					sql.As(t1.C(entaccount.FieldAddress), "address"),
					sql.As(t1.C(entaccount.FieldActive), "active"),
					sql.As(t1.C(entaccount.FieldLocked), "locked"),
					sql.As(t1.C(entaccount.FieldLockedBy), "locked_by"),
					sql.As(t1.C(entaccount.FieldBlocked), "blocked"),
					sql.As(t1.C(entaccount.FieldUsedFor), "used_for"),
					sql.As(t1.C(entaccount.FieldCoinTypeID), "coin_type_id"),
				)
		})
}

func expand(infos []*npool.Account) []*npool.Account {
	for _, info := range infos {
		info.UsedFor = accountmgrpb.AccountUsedFor(accountmgrpb.AccountUsedFor_value[info.UsedForStr])
		info.LockedBy = accountmgrpb.LockedBy(accountmgrpb.LockedBy_value[info.LockedByStr])
	}
	return infos
}
