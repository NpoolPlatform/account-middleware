package payment

import (
	"context"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	paymentcrud "github.com/NpoolPlatform/account-manager/pkg/crud/payment"
	accountpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	mgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/payment"
	mwpb "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func CreateAccount(ctx context.Context, in *mwpb.AccountReq) (info *mwpb.Account, err error) {
	var id string

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		usedFor := accountpb.AccountUsedFor_GoodPayment
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
		info2, err := paymentcrud.CreateSet(tx.Payment.Create(), &mgrpb.AccountReq{
			ID:            in.ID,
			AccountID:     &accountID,
			CollectingTID: in.CollectingTID,
			AvailableAt:   in.AvailableAt,
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
