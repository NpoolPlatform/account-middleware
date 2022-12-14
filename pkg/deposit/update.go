package deposit

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	depositcrud "github.com/NpoolPlatform/account-manager/pkg/crud/deposit"
	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	entdeposit "github.com/NpoolPlatform/account-manager/pkg/db/ent/deposit"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	depositmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"github.com/google/uuid"
)

func UpdateAccount(ctx context.Context, in *npool.AccountReq) (info *npool.Account, err error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "UpdateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		deposit, err := tx.Deposit.
			Query().
			Where(
				entdeposit.ID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}
		if deposit == nil {
			return fmt.Errorf("invalid deposit")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(deposit.AccountID),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(account, &accountmgrpb.AccountReq{
			Active:   in.Active,
			Locked:   in.Locked,
			LockedBy: in.LockedBy,
			Blocked:  in.Blocked,
		}).Save(ctx); err != nil {
			return err
		}

		u, err := depositcrud.UpdateSet(deposit, &depositmgrpb.AccountReq{
			CollectingTID: in.CollectingTID,
			Incoming:      in.Incoming,
			Outcoming:     in.Outcoming,
			ScannableAt:   in.ScannableAt,
		})
		if err != nil {
			return err
		}

		if _, err = u.Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetAccount(ctx, in.GetID())
}
