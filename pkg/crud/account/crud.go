package account

import (
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID                     *uuid.UUID
	CoinTypeID             *uuid.UUID
	Address                *string
	UsedFor                *basetypes.AccountUsedFor
	PlatformHoldPrivateKey *bool
	Active                 *bool
	Locked                 *bool
	LockedBy               *basetypes.AccountLockedBy
	Blocked                *bool
	DeletedAt              *uint32
}

func CreateSet(c *ent.AccountCreate, req *Req) *ent.AccountCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Address != nil {
		c.SetAddress(*req.Address)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.PlatformHoldPrivateKey != nil {
		c.SetPlatformHoldPrivateKey(*req.PlatformHoldPrivateKey)
	}
	if req.Active != nil {
		c.SetActive(*req.Active)
	}
	if req.Locked != nil {
		c.SetLocked(*req.Locked)
	}
	if req.LockedBy != nil {
		c.SetLockedBy(req.LockedBy.String())
	}
	if req.Blocked != nil {
		c.SetBlocked(*req.Blocked)
	}
	return c
}

func UpdateSet(u *ent.AccountUpdateOne, req *Req) *ent.AccountUpdateOne {
	if req.Active != nil {
		u.SetActive(*req.Active)
	}
	if req.Locked != nil {
		u.SetLocked(*req.Locked)
	}
	if req.LockedBy != nil {
		u.SetLockedBy(req.LockedBy.String())
	}
	if req.Blocked != nil {
		u.SetBlocked(*req.Blocked)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID                     *cruder.Cond
	IDs                    *cruder.Cond
	CoinTypeID             *cruder.Cond
	Address                *cruder.Cond
	UsedFor                *cruder.Cond
	PlatformHoldPrivateKey *cruder.Cond
	Active                 *cruder.Cond
	Locked                 *cruder.Cond
	LockedBy               *cruder.Cond
	Blocked                *cruder.Cond
	CreatedAt              *cruder.Cond
}

func SetQueryConds(q *ent.AccountQuery, conds *Conds) (*ent.AccountQuery, error) { // nolint
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid account id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entaccount.ID(id))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid account cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entaccount.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Address != nil {
		addr, ok := conds.Address.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid account address")
		}
		switch conds.Address.Op {
		case cruder.EQ:
			q.Where(entaccount.Address(addr))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.AccountUsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid account usedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entaccount.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.PlatformHoldPrivateKey != nil {
		hold, ok := conds.PlatformHoldPrivateKey.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid account platformholdprivatekey")
		}
		switch conds.PlatformHoldPrivateKey.Op {
		case cruder.EQ:
			q.Where(entaccount.PlatformHoldPrivateKey(hold))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Active != nil {
		active, ok := conds.Active.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid account active")
		}
		switch conds.Active.Op {
		case cruder.EQ:
			q.Where(entaccount.Active(active))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Locked != nil {
		locked, ok := conds.Locked.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid account locked")
		}
		switch conds.Locked.Op {
		case cruder.EQ:
			q.Where(entaccount.Locked(locked))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.LockedBy != nil {
		lockedBy, ok := conds.LockedBy.Val.(basetypes.AccountLockedBy)
		if !ok {
			return nil, fmt.Errorf("invalid account lockedby")
		}
		switch conds.LockedBy.Op {
		case cruder.EQ:
			q.Where(entaccount.LockedBy(lockedBy.String()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.Blocked != nil {
		blocked, ok := conds.Blocked.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid account blocked")
		}
		switch conds.Blocked.Op {
		case cruder.EQ:
			q.Where(entaccount.Blocked(blocked))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid account ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entaccount.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.CreatedAt != nil {
		at, ok := conds.CreatedAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid account createdat")
		}
		switch conds.CreatedAt.Op {
		case cruder.GTE:
			q.Where(entaccount.CreatedAtGTE(at))
		case cruder.LTE:
			q.Where(entaccount.CreatedAtLTE(at))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	q.Where(entaccount.DeletedAt(0))
	return q, nil
}
