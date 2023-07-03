package goodbenefit

import (
	"context"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	goodbenefitcrud "github.com/NpoolPlatform/account-manager/pkg/crud/goodbenefit"
	accountpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"
	mwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
)

func CreateAccount(ctx context.Context, in *mwpb.AccountReq) (info *mwpb.Account, err error) {
	var id string

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		usedFor := accountpb.AccountUsedFor_GoodBenefit
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
		info2, err := goodbenefitcrud.CreateSet(tx.GoodBenefit.Create(), &mgrpb.AccountReq{
			ID:        in.ID,
			GoodID:    in.GoodID,
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
