package user

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	usercrud "github.com/NpoolPlatform/account-manager/pkg/crud/user"
	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	entuser "github.com/NpoolPlatform/account-manager/pkg/db/ent/user"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	usermgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/user"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"

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

	span = commontracer.TraceInvoker(span, "user", "user", "UpdateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(uuid.MustParse(in.GetAccountID())),
			).Only(ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(account, &accountmgrpb.AccountReq{
			Active:  in.Active,
			Locked:  in.Locked,
			Blocked: in.Blocked,
		}).Save(ctx); err != nil {
			return err
		}

		user, err := tx.User.
			Query().
			Where(
				entuser.ID(uuid.MustParse(in.GetID())),
			).Only(ctx)
		if err != nil {
			return err
		}

		if _, err = usercrud.UpdateSet(user, &usermgrpb.AccountReq{
			Labels: in.Labels,
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
