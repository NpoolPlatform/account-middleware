package user

import (
	"context"
	"time"

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

func DeleteAccount(ctx context.Context, id string) (info *npool.Account, err error) {
	info, err = GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		user, err := tx.User.
			Query().
			Where(
				entuser.ID(uuid.MustParse(id)),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = usercrud.UpdateSet(user, &usermgrpb.AccountReq{
			DeletedAt: &now,
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
			DeletedAt: &now,
		}).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
