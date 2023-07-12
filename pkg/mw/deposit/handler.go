package deposit

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/account-middleware/pkg/const"
	depositcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/deposit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	CoinTypeID    *uuid.UUID
	AccountID     *uuid.UUID
	Address       *string
	Active        *bool
	Locked        *bool
	LockedBy      *basetypes.AccountLockedBy
	Blocked       *bool
	CollectingTID *uuid.UUID
	Incoming      *decimal.Decimal
	Outcoming     *decimal.Decimal
	Conds         *depositcrud.Conds
	Offset        int32
	Limit         int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithAccountID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AccountID = &_id
		return nil
	}
}

func WithAddress(addr *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if addr == nil {
			return nil
		}
		if *addr == "" {
			return fmt.Errorf("invalid address")
		}
		h.Address = addr
		return nil
	}
}

func WithActive(active *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Active = active
		return nil
	}
}

func WithLocked(locked *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Locked = locked
		return nil
	}
}

func WithLockedBy(lockedBy *basetypes.AccountLockedBy) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lockedBy == nil {
			return nil
		}
		switch *lockedBy {
		case basetypes.AccountLockedBy_Payment:
		case basetypes.AccountLockedBy_Collecting:
		default:
			return fmt.Errorf("invalid lockedby")
		}
		h.LockedBy = lockedBy
		return nil
	}
}

func WithBlocked(blocked *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Blocked = blocked
		return nil
	}
}

func WithCollectingTID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CollectingTID = &_id
		return nil
	}
}

func WithIncoming(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.Incoming = &amount
		return nil
	}
}

func WithOutcoming(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.Outcoming = &amount
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &depositcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: id}
		}
		if conds.AccountID != nil {
			id, err := uuid.Parse(conds.GetAccountID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AccountID = &cruder.Cond{Op: conds.GetAccountID().GetOp(), Val: id}
		}
		if conds.Address != nil {
			h.Conds.Address = &cruder.Cond{
				Op:  conds.GetAddress().GetOp(),
				Val: conds.GetAddress().GetValue(),
			}
		}
		if conds.Active != nil {
			h.Conds.Active = &cruder.Cond{
				Op:  conds.GetActive().GetOp(),
				Val: conds.GetActive().GetValue(),
			}
		}
		if conds.Locked != nil {
			h.Conds.Locked = &cruder.Cond{
				Op:  conds.GetLocked().GetOp(),
				Val: conds.GetLocked().GetValue(),
			}
		}
		if conds.LockedBy != nil {
			h.Conds.LockedBy = &cruder.Cond{
				Op:  conds.GetLockedBy().GetOp(),
				Val: basetypes.AccountLockedBy(conds.GetLockedBy().GetValue()),
			}
		}
		if conds.Blocked != nil {
			h.Conds.Blocked = &cruder.Cond{
				Op:  conds.GetBlocked().GetOp(),
				Val: conds.GetBlocked().GetValue(),
			}
		}
		if conds.ScannableAt != nil {
			h.Conds.ScannableAt = &cruder.Cond{
				Op:  conds.GetScannableAt().GetOp(),
				Val: conds.GetScannableAt().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
