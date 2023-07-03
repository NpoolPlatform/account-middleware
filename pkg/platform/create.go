package platform

import (
	"context"

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
