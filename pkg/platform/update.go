package platform

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	platformcrud "github.com/NpoolPlatform/account-manager/pkg/crud/platform"
	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	entplatform "github.com/NpoolPlatform/account-manager/pkg/db/ent/platform"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	platformmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

	"github.com/google/uuid"
)

func UpdateAccount(ctx context.Context, in *npool.AccountReq) (info *npool.Account, err error) { //nolint
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "platform", "platform", "UpdateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		platform, err := tx.Platform.
			Query().
			Where(
				entplatform.ID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if !in.GetBackup() {
			platformAccount, err := tx.Platform.
				Query().
				Where(
					entplatform.UsedFor(platform.UsedFor),
					entplatform.Backup(false),
				).
				ForUpdate().
				Only(ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return err
				}
			}

			if platformAccount != nil {
				backup := true
				if _, err = platformcrud.UpdateSet(platformAccount, &platformmgrpb.AccountReq{
					Backup: &backup,
				}).Save(ctx); err != nil {
					return err
				}
			}
		}

		if _, err := platformcrud.UpdateSet(platform, &platformmgrpb.AccountReq{
			Backup: in.Backup,
		}).Save(ctx); err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(platform.AccountID),
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
