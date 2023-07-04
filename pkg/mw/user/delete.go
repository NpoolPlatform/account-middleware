package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	usercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/user"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entuser "github.com/NpoolPlatform/account-middleware/pkg/db/ent/user"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
)

func (h *Handler) DeleteAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		user, err := tx.User.
			Query().
			Where(
				entuser.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("invalid user")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(user.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := usercrud.UpdateSet(
			user.Update(),
			&usercrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
