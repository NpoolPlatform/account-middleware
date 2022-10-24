package platform

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	platformcrud "github.com/NpoolPlatform/account-manager/pkg/crud/platform"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"
	mwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
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

	span = commontracer.TraceInvoker(span, "platform", "platform", "CreateTX")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		privateKey := true

		switch in.GetUsedFor() {
		case accountmgrpb.AccountUsedFor_UserBenefitHot:
		case accountmgrpb.AccountUsedFor_UserBenefitCold:
			privateKey = false
		case accountmgrpb.AccountUsedFor_PlatformBenefitCold:
			privateKey = false
		case accountmgrpb.AccountUsedFor_GasProvider:
		case accountmgrpb.AccountUsedFor_PaymentCollector:
			privateKey = false
		}

		info1, err := accountcrud.CreateSet(tx.Account.Create(), &accountmgrpb.AccountReq{
			CoinTypeID:             in.CoinTypeID,
			Address:                in.Address,
			UsedFor:                in.UsedFor,
			PlatformHoldPrivateKey: &privateKey,
		}).Save(ctx)
		if err != nil {
			return err
		}

		accountID := info1.ID.String()
		info2, err := platformcrud.CreateSet(tx.Platform.Create(), &mgrpb.AccountReq{
			ID:        in.ID,
			UsedFor:   in.UsedFor,
			AccountID: &accountID,
			Backup:    in.Backup,
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
