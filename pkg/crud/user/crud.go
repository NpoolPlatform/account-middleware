package user

import (
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entuser "github.com/NpoolPlatform/account-middleware/pkg/db/ent/user"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	AccountID  *uuid.UUID
	UsedFor    *basetypes.AccountUsedFor
	Labels     []string
	Memo       *string
	DeletedAt  *uint32
}

func CreateSet(c *ent.UserCreate, req *Req) *ent.UserCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if len(req.Labels) > 0 {
		c.SetLabels(req.Labels)
	}
	if req.Memo != nil {
		c.SetMemo(*req.Memo)
	}
	return c
}

func UpdateSet(u *ent.UserUpdateOne, req *Req) *ent.UserUpdateOne {
	if len(req.Labels) > 0 {
		u.SetLabels(req.Labels)
	}
	if req.Memo != nil {
		u.SetMemo(*req.Memo)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	CoinTypeID *cruder.Cond
	AccountID  *cruder.Cond
	UsedFor    *cruder.Cond
	IDs        *cruder.Cond
	AccountIDs *cruder.Cond
}

func SetQueryConds(q *ent.UserQuery, conds *Conds) (*ent.UserQuery, error) { //nolint
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entuser.ID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entuser.UserID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entuser.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entuser.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.AccountUsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid user accountusedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entuser.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entuser.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.AccountIDs != nil {
		ids, ok := conds.AccountIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user accountids")
		}
		switch conds.AccountIDs.Op {
		case cruder.IN:
			q.Where(entuser.AccountIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	q.Where(entuser.DeletedAt(0))
	return q, nil
}
