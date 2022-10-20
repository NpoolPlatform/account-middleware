package deposit

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/deposit"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	depositcrud "github.com/NpoolPlatform/account-manager/pkg/crud/deposit"
	accountpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
)

func CreateAccount(ctx context.Context, in *npool.AccountReq) (info *npool.Account, err error) {
	var id string

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "CreateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		usedFor := accountpb.AccountUsedFor_UserDeposit
		privateKey := true
		info1, err := accountcrud.CreateSet(tx.Account.Create(), &accountpb.AccountReq{
			CoinTypeID:             in.CoinTypeID,
			Address:                in.Address,
			UsedFor:                &usedFor,
			PlatformHoldPrivateKey: &privateKey,
		}).Save(ctx)

		if err != nil {
			return err
		}

		accountID := info1.ID.String()

		info2, err := depositcrud.CreateSet(tx.Deposit.Create(), &mgrpb.AccountReq{
			ID:         in.ID,
			AppID:      in.AppID,
			UserID:     in.UserID,
			CoinTypeID: in.CoinTypeID,
			AccountID:  &accountID,
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
