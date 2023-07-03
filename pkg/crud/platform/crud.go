package platform

import (
	"fmt"

	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entplatform "github.com/NpoolPlatform/account-manager/pkg/db/ent/platform"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	accountcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AccountID *uuid.UUID
	UsedFor   *basetypes.AccountUsedFor
	Backup    *bool
}

func CreateSet(c *ent.PlatformCreate, req *Req) *ent.PlatformCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Backup != nil {
		c.SetBackup(*req.Backup)
	}
	return c
}

func UpdateSet(u *ent.PlatformUpdateOne, req *Req) *ent.PlatformUpdateOne {
	if req.Backup != nil {
		u.SetBackup(*req.Backup)
	}
	return u
}

type Conds struct {
	*accountcrud.Conds
	AccountID *cruder.Cond
	UsedFor   *cruder.Cond
	Backup    *cruder.Cond
}

func SetQueryConds(q *ent.PlatformQuery, conds *Conds) (*ent.PlatformQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid platform id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entplatform.ID(id))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid platform accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entplatform.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.AccountUsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid platform accountusedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entplatform.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	if conds.Backup != nil {
		backup, ok := conds.Backup.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid platform backup")
		}
		switch conds.Backup.Op {
		case cruder.EQ:
			q.Where(entplatform.Backup(backup))
		default:
			return nil, fmt.Errorf("invalid platform field")
		}
	}
	q.Where(entplatform.DeletedAt(0))
	return q, nil
}
