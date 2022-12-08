package deposit

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"

	depositcrud "github.com/NpoolPlatform/account-manager/pkg/crud/deposit"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	depositmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func GetAccount(ctx context.Context, id string) (info *npool.Account, err error) {
	infos := []*npool.Account{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "QueryJoin")

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
				deposit.FieldAccountID,
				deposit.FieldCollectingTid,
				deposit.FieldCreatedAt,
				deposit.FieldIncoming,
				deposit.FieldOutcoming,
				deposit.FieldScannableAt,
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
						sql.As(t1.C(account.FieldCoinTypeID), "coin_type_id"),
						sql.As(t1.C(account.FieldAddress), "address"),
						sql.As(t1.C(account.FieldActive), "active"),
						sql.As(t1.C(account.FieldLocked), "locked"),
						sql.As(t1.C(account.FieldLockedBy), "locked_by"),
						sql.As(t1.C(account.FieldBlocked), "blocked"),
					)
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}

	return infos[0], nil
}

//nolint
func GetAccounts(ctx context.Context,
	conds *npool.Conds,
	offset,
	limit int32,
) (infos []*npool.Account, total uint32, err error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccounts")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "QueryJoin")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := depositcrud.SetQueryConds(&depositmgrpb.Conds{
			ID:          conds.ID,
			AppID:       conds.AppID,
			UserID:      conds.UserID,
			AccountID:   conds.AccountID,
			ScannableAt: conds.ScannableAt,
		}, cli)
		if err != nil {
			return err
		}

		stm.Where(deposit.ScannableAtLT(uint32(time.Now().Unix())))

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		return stm.
			Select(
				deposit.FieldID,
				deposit.FieldAppID,
				deposit.FieldUserID,
				deposit.FieldAccountID,
				deposit.FieldCollectingTid,
				deposit.FieldCreatedAt,
				deposit.FieldIncoming,
				deposit.FieldOutcoming,
				deposit.FieldScannableAt,
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
							accountmgrpb.LockedBy(conds.GetLockedBy().GetValue()).String(),
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
						sql.As(t1.C(account.FieldCoinTypeID), "coin_type_id"),
						sql.As(t1.C(account.FieldAddress), "address"),
						sql.As(t1.C(account.FieldActive), "active"),
						sql.As(t1.C(account.FieldLocked), "locked"),
						sql.As(t1.C(account.FieldLockedBy), "locked_by"),
						sql.As(t1.C(account.FieldBlocked), "blocked"),
					)
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, total, err
	}
	for key := range infos {
		if infos[key].CoinTypeID == "" {
			infos[key].CoinTypeID = uuid.NewString()
		}
	}
	return infos, total, nil
}
