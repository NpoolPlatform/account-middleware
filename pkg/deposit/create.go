package user

import (
	"context"

	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"

	account "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"github.com/google/uuid"
)

func CreateAccount(ctx context.Context, in *npool.AccountReq) (info *npool.Account, err error) {
	var id string

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info1, err := tx.
			Account.
			Create().
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetAddress(in.GetAddress()).
			SetUsedFor(account.AccountUsedFor_UserDeposit.String()).
			Save(ctx)
		if err != nil {
			return err
		}

		info2, err := tx.
			Deposit.
			Create().
			SetAppID(uuid.MustParse(in.GetAppID())).
			SetUserID(uuid.MustParse(in.GetUserID())).
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetAccountID(info1.ID).
			Save(ctx)
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
