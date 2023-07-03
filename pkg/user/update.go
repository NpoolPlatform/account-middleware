package user

import (
	"context"

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
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		user, err := tx.User.
			Query().
			Where(
				entuser.ID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = usercrud.UpdateSet(user, &usermgrpb.AccountReq{
			Labels: in.Labels,
			Memo:   in.Memo,
		}).Save(ctx); err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(user.AccountID),
			).
			ForUpdate().
			Only(ctx)
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

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetAccount(ctx, in.GetID())
}
