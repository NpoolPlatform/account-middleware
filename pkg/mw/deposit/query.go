package deposit

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entdeposit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/deposit"

	depositcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/deposit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.DepositSelect
	infos []*npool.Account
	total uint32
}

func (h *queryHandler) selectAccount(stm *ent.DepositQuery) {
	h.stm = stm.Select(
		entdeposit.FieldID,
		entdeposit.FieldAppID,
		entdeposit.FieldUserID,
		entdeposit.FieldAccountID,
		entdeposit.FieldCollectingTid,
		entdeposit.FieldCreatedAt,
		entdeposit.FieldIncoming,
		entdeposit.FieldIncoming,
		entdeposit.FieldOutcoming,
		entdeposit.FieldScannableAt,
		entdeposit.FieldCreatedAt,
		entdeposit.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAccount(cli *ent.Client) {
	h.selectAccount(
		cli.Deposit.
			Query().
			Where(
				entdeposit.ID(*h.ID),
				entdeposit.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryAccounts(ctx context.Context, cli *ent.Client) error {
	stm, err := depositcrud.SetQueryConds(cli.Deposit.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.selectAccount(stm)
	return nil
}

func (h *queryHandler) queryJoinAccount(s *sql.Selector) error {
	t := sql.Table(entaccount.Table)
	s.LeftJoin(t).
		On(
			s.C(entdeposit.FieldAccountID),
			t.C(entaccount.FieldID),
		)

	if h.Conds != nil && h.Conds.CoinTypeID != nil {
		id, ok := h.Conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid deposit cointypeid")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldCoinTypeID), id),
		)
	}
	if h.Conds != nil && h.Conds.Address != nil {
		addr, ok := h.Conds.Address.Val.(string)
		if !ok {
			return fmt.Errorf("invalid deposit address")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldAddress), addr),
		)
	}
	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid deposit active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Locked != nil {
		locked, ok := h.Conds.Locked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid deposit locked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLocked), locked),
		)
	}
	if h.Conds != nil && h.Conds.LockedBy != nil {
		lockedBy, ok := h.Conds.LockedBy.Val.(basetypes.AccountLockedBy)
		if !ok {
			return fmt.Errorf("invalid deposit lockedby")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldLockedBy), lockedBy.String()),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid deposit blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entaccount.FieldCoinTypeID), "coin_type_id"),
		sql.As(t.C(entaccount.FieldAddress), "address"),
		sql.As(t.C(entaccount.FieldActive), "active"),
		sql.As(t.C(entaccount.FieldLocked), "locked"),
		sql.As(t.C(entaccount.FieldLockedBy), "locked_by"),
		sql.As(t.C(entaccount.FieldBlocked), "blocked"),
	)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAccount(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		if _, err := uuid.Parse(info.CoinTypeID); err != nil {
			info.CoinTypeID = uuid.Nil.String()
		}
	}
}

func (h *Handler) GetAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAccount(cli)
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetAccounts(ctx context.Context) ([]*npool.Account, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()

		_total, err := handler.stm.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entdeposit.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}