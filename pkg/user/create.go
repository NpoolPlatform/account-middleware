package user

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	usercrud "github.com/NpoolPlatform/account-manager/pkg/crud/user"
	accountpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/user"
	mwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
)

func CreateAccount(ctx context.Context, in *mwpb.AccountReq) (info *mwpb.Account, err error) {
	var id string

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "user", "CreateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		privateKey := true

		info1, err := accountcrud.CreateSet(tx.Account.Create(), &accountpb.AccountReq{
			CoinTypeID:             in.CoinTypeID,
			Address:                in.Address,
			UsedFor:                in.UsedFor,
			PlatformHoldPrivateKey: &privateKey,
		}).Save(ctx)
		if err != nil {
			return err
		}

		accountID := info1.ID.String()
		info2, err := usercrud.CreateSet(tx.User.Create(), &mgrpb.AccountReq{
			AppID:      in.AppID,
			UserID:     in.UserID,
			CoinTypeID: in.CoinTypeID,
			AccountID:  &accountID,
			UsedFor:    in.UsedFor,
			Labels:     in.Labels,
		}).Save(ctx)
		if err != nil {
			return err
		}

		id = info2.ID.String()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetAccount(ctx, id)
}
