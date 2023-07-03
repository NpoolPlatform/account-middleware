//nolint:dupl
package deposit

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entdeposit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/deposit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	ID            *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	AccountID     *uuid.UUID
	CollectingTID *uuid.UUID
	Incoming      *decimal.Decimal
	Outcoming     *decimal.Decimal
	ScannableAt   *uint32
	DeletedAt     *uint32
}

func CreateSet(c *ent.DepositCreate, req *Req) *ent.DepositCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}

	c.SetIncoming(decimal.NewFromInt(0))
	c.SetOutcoming(decimal.NewFromInt(0))
	c.SetScannableAt(uint32(time.Now().Unix()))

	return c
}

func UpdateSet(u *ent.DepositUpdateOne, req *Req) (*ent.DepositUpdateOne, error) {
	if req.CollectingTID != nil {
		u.SetCollectingTid(*req.CollectingTID)
	}

	incoming, ok := u.Mutation().Incoming()
	if !ok {
		return nil, fmt.Errorf("invalid incoming")
	}
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
	}
	outcoming, ok := u.Mutation().Outcoming()
	if !ok {
		return nil, fmt.Errorf("invalid outcoming")
	}
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
	}

	if incoming.Cmp(outcoming) < 0 {
		return nil, fmt.Errorf("incoming (%v) < outcoming (%v)", incoming, outcoming)
	}

	if req.Incoming != nil {
		u.SetIncoming(incoming)
	}
	if req.Outcoming != nil {
		u.SetOutcoming(outcoming)
	}

	if req.ScannableAt != nil {
		u.SetScannableAt(*req.ScannableAt)
	}

	return u, nil
}

type Conds struct {
	*accountcrud.Conds
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	AccountID   *cruder.Cond
	ScannableAt *cruder.Cond
}

func SetQueryConds(q *ent.DepositQuery, conds *Conds) (*ent.DepositQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entdeposit.ID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entdeposit.AppID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entdeposit.UserID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entdeposit.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.ScannableAt != nil {
		at, ok := conds.ScannableAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid deposit scannableat")
		}
		switch conds.ScannableAt.Op {
		case cruder.LT:
			q.Where(entdeposit.ScannableAtLT(at))
		case cruder.GT:
			q.Where(entdeposit.ScannableAtGT(at))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	q.Where(entdeposit.DeletedAt(0))
	return q, nil
}