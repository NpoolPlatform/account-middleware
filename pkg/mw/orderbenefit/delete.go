package orderbenefit

import (
	"context"
	"time"

	crud "github.com/NpoolPlatform/account-middleware/pkg/crud/orderbenefit"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAccountBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne := crud.UpdateSet(tx.OrderBenefit.UpdateOneID(*h.ID), &crud.Req{DeletedAt: &now})
	_, err := updateOne.Save(ctx)
	return err
}

func (h *Handler) DeleteAccount(ctx context.Context) error {
	info, err := h.GetAccount(ctx)
	if err != nil {
		return err
	}

	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAccountBase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
