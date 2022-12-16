package goodbenefit

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/account-manager/pkg/crud/goodbenefit"
	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	entgoodbenefit "github.com/NpoolPlatform/account-manager/pkg/db/ent/goodbenefit"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	goodbenefitpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

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

	span = commontracer.TraceInvoker(span, "goodbenefit", "goodbenefit", "UpdateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		goodBenefit, err := tx.GoodBenefit.
			Query().
			Where(
				entgoodbenefit.ID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if !in.GetBackup() && in.Backup != nil {
			gb, err := tx.GoodBenefit.
				Query().
				Where(
					entgoodbenefit.GoodID(goodBenefit.GoodID),
					entgoodbenefit.Backup(false),
				).
				ForUpdate().
				Only(ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return err
				}
			}

			if gb != nil {
				backup := true
				if _, err = goodbenefitcrud.UpdateSet(gb, &goodbenefitpb.AccountReq{
					Backup: &backup,
				}).Save(ctx); err != nil {
					return err
				}
			}
		}

		if _, err = goodbenefitcrud.UpdateSet(goodBenefit, &goodbenefitpb.AccountReq{
			TransactionID: in.TransactionID,
			Backup:        in.Backup,
			IntervalHours: in.IntervalHours,
		}).Save(ctx); err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(goodBenefit.AccountID),
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

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetAccount(ctx, in.GetID())
}
