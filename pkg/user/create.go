package user

import (
	"context"

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
			ID:         in.ID,
			AppID:      in.AppID,
			UserID:     in.UserID,
			CoinTypeID: in.CoinTypeID,
			AccountID:  &accountID,
			UsedFor:    in.UsedFor,
			Labels:     in.Labels,
			Memo:       in.Memo,
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
