package payment

import (
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/account-middleware/pkg/db/ent/payment"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	ID            *uuid.UUID
	AccountID     *uuid.UUID
	CollectingTID *uuid.UUID
	AvailableAt   *uint32
}

func CreateSet(c *ent.PaymentCreate, req *Req) *ent.PaymentCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.CollectingTID != nil {
		c.SetCollectingTid(*req.CollectingTID)
	}
	return c
}

func UpdateSet(u *ent.PaymentUpdateOne, req *Req) *ent.PaymentUpdateOne {
	if req.CollectingTID != nil {
		u.SetCollectingTid(*req.CollectingTID)
	}
	if req.AvailableAt != nil {
		u.SetAvailableAt(*req.AvailableAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	AccountID   *cruder.Cond
	AccountIDs  *cruder.Cond
	AvailableAt *cruder.Cond
}

func SetQueryConds(q *ent.PaymentQuery, conds *Conds) (*ent.PaymentQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid payment id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entpayment.ID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid payment accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entpayment.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AccountIDs != nil {
		ids, ok := conds.AccountIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid payment accountids")
		}
		switch conds.AccountIDs.Op {
		case cruder.IN:
			q.Where(entpayment.AccountIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AvailableAt != nil {
		at, ok := conds.AvailableAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid payment availableat")
		}
		switch conds.AvailableAt.Op {
		case cruder.GTE:
			q.Where(entpayment.AvailableAtGTE(at))
		case cruder.LTE:
			q.Where(entpayment.AvailableAtLTE(at))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	q.Where(entpayment.DeletedAt(0))
	return q, nil
}
