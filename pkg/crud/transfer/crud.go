package transfer

import (
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	enttransfer "github.com/NpoolPlatform/account-middleware/pkg/db/ent/transfer"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	ID           *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	TargetUserID *uuid.UUID
	DeletedAt    *uint32
}

func CreateSet(c *ent.TransferCreate, req *Req) *ent.TransferCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.TargetUserID != nil {
		c.SetTargetUserID(*req.TargetUserID)
	}
	return c
}

func UpdateSet(u *ent.TransferUpdateOne, req *Req) *ent.TransferUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	*accountcrud.Conds
	AppID        *cruder.Cond
	UserID       *cruder.Cond
	TargetUserID *cruder.Cond
}

func SetQueryConds(q *ent.TransferQuery, conds *Conds) (*ent.TransferQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enttransfer.ID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttransfer.AppID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(enttransfer.UserID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TargetUserID != nil {
		id, ok := conds.TargetUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid transfer targetuserid")
		}
		switch conds.TargetUserID.Op {
		case cruder.EQ:
			q.Where(enttransfer.TargetUserID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	q.Where(enttransfer.DeletedAt(0))
	return q, nil
}
