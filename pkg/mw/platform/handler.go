package platform

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/account-middleware/pkg/const"
	platformcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/platform"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID         *uuid.UUID
	CoinTypeID *uuid.UUID
	UsedFor    *basetypes.AccountUsedFor
	AccountID  *uuid.UUID
	Address    *string
	Backup     *bool
	Active     *bool
	Locked     *bool
	LockedBy   *basetypes.AccountLockedBy
	Blocked    *bool
	Conds      *platformcrud.Conds
	Offset     int32
	Limit      int32
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

func WithUsedFor(usedFor *basetypes.AccountUsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usedFor == nil {
			return nil
		}
		switch *usedFor {
		case basetypes.AccountUsedFor_UserBenefitHot:
		case basetypes.AccountUsedFor_UserBenefitCold:
		case basetypes.AccountUsedFor_PlatformBenefitCold:
		case basetypes.AccountUsedFor_GasProvider:
		case basetypes.AccountUsedFor_PaymentCollector:
		default:
			return fmt.Errorf("invalid usedfor")
		}
		h.UsedFor = usedFor
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
		if len(*addr) == 0 {
			return fmt.Errorf("invalid address")
		}
		h.Address = addr
		return nil
	}
}

func WithBackup(backup *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Backup = backup
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &platformcrud.Conds{}
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
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: id}
		}
		if conds.UsedFor != nil {
			h.Conds.UsedFor = &cruder.Cond{
				Op:  conds.GetUsedFor().GetOp(),
				Val: basetypes.AccountUsedFor(conds.GetUsedFor().GetValue()),
			}
		}
		if conds.Backup != nil {
			h.Conds.Backup = &cruder.Cond{
				Op:  conds.GetBackup().GetOp(),
				Val: conds.GetBackup().GetValue(),
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
