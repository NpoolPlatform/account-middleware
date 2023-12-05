package transfer

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"

	transfercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/transfer"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateTransfer(ctx context.Context) (*npool.Transfer, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.AppID, *h.UserID, *h.TargetUserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID.String()},
			TargetUserID: &basetypes.StringVal{Op: cruder.EQ, Value: h.TargetUserID.String()},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistTransferConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("transfer exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := transfercrud.CreateSet(
			tx.Transfer.Create(),
			&transfercrud.Req{
				EntID:        h.EntID,
				AppID:        h.AppID,
				UserID:       h.UserID,
				TargetUserID: h.TargetUserID,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTransfer(ctx)
}
