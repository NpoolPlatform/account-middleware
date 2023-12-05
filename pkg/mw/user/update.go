package user

import (
	"context"
	"fmt"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	usercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/user"
	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entuser "github.com/NpoolPlatform/account-middleware/pkg/db/ent/user"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
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
				entaccount.EntID(user.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				Active:  h.Active,
				Locked:  h.Locked,
				Blocked: h.Blocked,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := usercrud.UpdateSet(
			user.Update(),
			&usercrud.Req{
				Labels: h.Labels,
				Memo:   h.Memo,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
