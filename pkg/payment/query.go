package payment

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/account-manager/pkg/crud/payment"
	entpayment "github.com/NpoolPlatform/account-manager/pkg/db/ent/payment"
	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/payment"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"

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

	span = commontracer.TraceInvoker(span, "payment", "payment", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Payment.
			Query().
			Where(
				entpayment.ID(uuid.MustParse(id)),
			)
		return join(stm).Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	infos = expand(infos)

	return infos[0], nil
}

func GetAccounts(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Account, total uint32, err error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "payment", "payment", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:         conds.ID,
			AccountID:  conds.AccountID,
			AccountIDs: conds.AccountIDs,
		}, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		stm.Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
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

	span = commontracer.TraceInvoker(span, "payment", "payment", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&mgrpb.Conds{
			ID:        conds.ID,
			AccountID: conds.AccountID,
		}, cli)
		if err != nil {
			return err
		}

		return join(stm).
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

func join(stm *ent.PaymentQuery) *ent.PaymentSelect {
	return stm.Select(
		entpayment.FieldID,
		entpayment.FieldCollectingTid,
		entpayment.FieldAvailableAt,
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
					sql.As(t1.C(account.FieldCoinTypeID), "coin_type_id"),
				)
		})
}

func expand(infos []*npool.Account) []*npool.Account {
	for _, info := range infos {
		info.LockedBy = accountmgrpb.LockedBy(accountmgrpb.LockedBy_value[info.LockedByStr])
	}
	return infos
}
