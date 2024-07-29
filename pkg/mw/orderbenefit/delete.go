package orderbenefit

import (
	"context"
	"fmt"
	"time"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	"github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"

	crud "github.com/NpoolPlatform/account-middleware/pkg/crud/orderbenefit"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAccountBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	account, err := tx.Account.
		Query().
		Where(
			entaccount.EntID(*h.AccountID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err := accountcrud.UpdateSet(
		account.Update(),
		&accountcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return err
}

func (h *deleteHandler) deleteOrderBenefit(ctx context.Context, tx *ent.Tx) error {
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

func (h *Handler) DeleteAccounts(ctx context.Context) ([]*orderbenefit.Account, error) {
	entIDs := []uuid.UUID{}
	accountIDMap := make(map[string]*uuid.UUID)
	for _, req := range h.Reqs {
		if req.EntID == nil {
			return nil, fmt.Errorf("invaild entid")
		}
		entIDs = append(entIDs, *req.EntID)
		if req.AccountID != nil {
			accountIDMap[req.EntID.String()] = req.AccountID
		}
	}

	h.Conds = &crud.Conds{
		EntIDs: &cruder.Cond{
			Op:  cruder.IN,
			Val: entIDs,
		},
	}

	h.Limit = int32(len(entIDs))
	h.Offset = 0

	infos, _, err := h.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, info := range infos {
			h.ID = &info.ID
			handler := &deleteHandler{
				Handler: h,
			}
			if err := handler.deleteOrderBenefit(_ctx, tx); err != nil {
				return err
			}

			_, ok := accountIDMap[info.EntID]
			if !ok {
				accountID, err := uuid.Parse(info.AccountID)
				if err != nil {
					return err
				}
				handler.AccountID = &accountID
				err = handler.deleteAccountBase(ctx, tx)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
