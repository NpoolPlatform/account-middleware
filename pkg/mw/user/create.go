package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	usercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/user"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid userid")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid cointypeid")
	}
	if h.Address == nil {
		return fmt.Errorf("invalid address")
	}
	if h.UsedFor == nil {
		return fmt.Errorf("invalid usedfor")
	}
	return nil
}

func (h *Handler) CreateAccount(ctx context.Context) (*npool.Account, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.validate(); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v:%v:%v", basetypes.Prefix_PrefixCreateUserAccount, *h.AppID, *h.UserID, *h.CoinTypeID, *h.Address)
	if h.Memo != nil {
		key = fmt.Sprintf("%v:%v", key, *h.Memo)
	}
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id1 := uuid.New()
	if h.ID == nil {
		h.ID = &id1
	}

	id2 := uuid.New()
	if h.AccountID == nil {
		h.AccountID = &id2
	}

	privateKey := true

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := accountcrud.CreateSet(
			tx.Account.Create(),
			&accountcrud.Req{
				ID:                     h.AccountID,
				CoinTypeID:             h.CoinTypeID,
				Address:                h.Address,
				UsedFor:                h.UsedFor,
				PlatformHoldPrivateKey: &privateKey,
			},
		).Save(ctx); err != nil {
			return err
		}

		if _, err := usercrud.CreateSet(
			tx.User.Create(),
			&usercrud.Req{
				ID:         h.ID,
				AppID:      h.AppID,
				UserID:     h.UserID,
				CoinTypeID: h.CoinTypeID,
				AccountID:  h.AccountID,
				UsedFor:    h.UsedFor,
				Labels:     h.Labels,
				Memo:       h.Memo,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
