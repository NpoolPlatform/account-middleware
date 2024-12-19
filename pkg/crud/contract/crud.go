package contract

import (
	"fmt"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entcontract "github.com/NpoolPlatform/account-middleware/pkg/db/ent/contract"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/account/v1"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	GoodID       *uuid.UUID
	PledgeID     *uuid.UUID
	AccountID    *uuid.UUID
	Backup       *bool
	ContractType *basetypes.ContractType
	DeletedAt    *uint32
}

func CreateSet(c *ent.ContractCreate, req *Req) *ent.ContractCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.PledgeID != nil {
		c.SetPledgeID(*req.PledgeID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.Backup != nil {
		c.SetBackup(*req.Backup)
	}
	if req.ContractType != nil {
		c.SetContractType(req.ContractType.String())
	}
	return c
}

func UpdateSet(u *ent.ContractUpdateOne, req *Req) *ent.ContractUpdateOne {
	if req.Backup != nil {
		u.SetBackup(*req.Backup)
	}
	if req.AccountID != nil {
		u.SetAccountID(*req.AccountID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	GoodID       *cruder.Cond
	AccountID    *cruder.Cond
	Backup       *cruder.Cond
	ContractType *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.ContractQuery, conds *Conds) (*ent.ContractQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcontract.EntID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid contract id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcontract.ID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcontract.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entcontract.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.Backup != nil {
		backup, ok := conds.Backup.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid contract backup")
		}
		switch conds.Backup.Op {
		case cruder.EQ:
			q.Where(entcontract.Backup(backup))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.ContractType != nil {
		contractType, ok := conds.ContractType.Val.(basetypes.ContractType)
		if !ok {
			return nil, fmt.Errorf("invalid account contracttype")
		}
		switch conds.ContractType.Op {
		case cruder.EQ:
			q.Where(entcontract.ContractType(contractType.String()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	q.Where(entcontract.DeletedAt(0))
	return q, nil
}
