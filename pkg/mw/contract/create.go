package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	accountSQL string
	pledgeSQL  string
}

//nolint:goconst
func (h *createHandler) constructCreatepledgeSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into pledges "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "pledge_id"
	_sql += comma + "account_id"
	_sql += comma + "backup"
	_sql += comma + "contract_type"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as pledge_id", comma, *h.PledgeID)
	_sql += fmt.Sprintf("%v'%v' as account_id", comma, *h.AccountID)
	_sql += fmt.Sprintf("%v'%v' as backup", comma, *h.Backup)
	_sql += fmt.Sprintf("%v'%v' as contract_type", comma, *h.ContractType)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ")"
	h.pledgeSQL = _sql
}

func (h *createHandler) constructCreateaccountSQL() {
	usedFor := basetypes.AccountUsedFor_GoodBenefit
	privateKey := true
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into accounts "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "coin_type_id"
	comma = ", "
	_sql += comma + "address"
	_sql += comma + "used_for"
	_sql += comma + "backup"
	_sql += comma + "platform_hold_private_key"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.AccountID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as address", comma, *h.Address)
	_sql += fmt.Sprintf("%v'%v' as account_id", comma, *h.AccountID)
	_sql += fmt.Sprintf("%v'%v' as used_for", comma, usedFor)
	_sql += fmt.Sprintf("%v'%v' as platform_hold_private_key", comma, privateKey)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from accounts "
	_sql += fmt.Sprintf(
		"where coin_type_id = '%v' and address = '%v' and deleted_at = 0",
		*h.CoinTypeID,
		*h.Address,
	)
	_sql += " limit 1)"
	h.accountSQL = _sql
}

func (h *createHandler) createAccount(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.accountSQL)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create account: %v", err)
	}
	return nil
}

func (h *createHandler) createPledge(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.pledgeSQL)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create pledge: %v", err)
	}
	return nil
}

func (h *Handler) CreateAccount(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructCreateaccountSQL()
	handler.constructCreatepledgeSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAccount(_ctx, tx); err != nil {
			return err
		}
		return handler.createPledge(_ctx, tx)
	})
}
