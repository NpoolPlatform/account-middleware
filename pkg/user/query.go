package user

import (
	"context"
	"encoding/json"
	"fmt"

	crud "github.com/NpoolPlatform/account-manager/pkg/crud/user"
	entuser "github.com/NpoolPlatform/account-manager/pkg/db/ent/user"
	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/user"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"

	"github.com/google/uuid"
)

func GetAccount(ctx context.Context, id string) (info *npool.Account, err error) {
	var infos []*npool.Account

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "user", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			User.
			Query().
			Where(
				entuser.ID(uuid.MustParse(id)),
			)
		return join(stm, &npool.Conds{}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	infos = expand(infos)
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}

	return infos[0], nil
}

func GetAccounts(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Account, total uint32, err error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "user", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:         conds.ID,
			AppID:      conds.AppID,
			UserID:     conds.UserID,
			UsedFor:    conds.UsedFor,
			CoinTypeID: conds.CoinTypeID,
			AccountID:  conds.AccountID,
			IDs:        conds.IDs,
			AccountIDs: conds.AccountIDs,
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

func GetAccountOnly(ctx context.Context, conds *npool.Conds) (*npool.Account, error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccountOnly")
	defer span.End()

	var err error

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	infos := []*npool.Account{}

	span = commontracer.TraceInvoker(span, "user", "user", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:         conds.ID,
			AppID:      conds.AppID,
			UserID:     conds.UserID,
			UsedFor:    conds.UsedFor,
			CoinTypeID: conds.CoinTypeID,
			AccountID:  conds.AccountID,
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

func join(stm *ent.UserQuery, conds *npool.Conds) *ent.UserSelect {
	return stm.
		Modify(func(s *sql.Selector) {
			s.
				Select(
					s.C(entuser.FieldID),
					s.C(entuser.FieldAppID),
					s.C(entuser.FieldUserID),
					s.C(entuser.FieldCoinTypeID),
					s.C(entuser.FieldLabels),
				)

			t1 := sql.Table(account.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entuser.FieldAccountID),
					t1.C(account.FieldID),
				)

			if conds.Active != nil {
				s.Where(
					sql.EQ(
						t1.C(account.FieldActive),
						conds.GetActive().GetValue(),
					),
				)
			}
			if conds.Blocked != nil {
				s.Where(
					sql.EQ(
						t1.C(account.FieldBlocked),
						conds.GetBlocked().GetValue(),
					),
				)
			}

			s.
				AppendSelect(
					sql.As(t1.C(account.FieldID), "account_id"),
					sql.As(t1.C(account.FieldAddress), "address"),
					sql.As(t1.C(account.FieldActive), "active"),
					sql.As(t1.C(account.FieldBlocked), "blocked"),
					sql.As(t1.C(account.FieldUsedFor), "used_for"),
					sql.As(t1.C(account.FieldCreatedAt), "created_at"),
					sql.As(t1.C(account.FieldUpdatedAt), "updated_at"),
					sql.As(t1.C(account.FieldDeletedAt), "deleted_at"),
				)
		})
}

func expand(infos []*npool.Account) []*npool.Account {
	for _, info := range infos {
		info.UsedFor = accountmgrpb.AccountUsedFor(accountmgrpb.AccountUsedFor_value[info.UsedForStr])
		_ = json.Unmarshal([]byte(info.LabelsStr), &info.Labels)
	}
	return infos
}
