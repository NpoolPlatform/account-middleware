package orderbenefit

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entorderbenefit "github.com/NpoolPlatform/account-middleware/pkg/db/ent/orderbenefit"

	orderbenefitcrud "github.com/NpoolPlatform/account-middleware/pkg/crud/orderbenefit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.OrderBenefitSelect
	stmCount  *ent.OrderBenefitSelect
	infos     []*npool.Account
	total     uint32
}

func (h *queryHandler) selectAccount(stm *ent.OrderBenefitQuery) *ent.OrderBenefitSelect {
	return stm.Select(entorderbenefit.FieldEntID)
}

func (h *queryHandler) queryAccount(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.OrderBenefit.Query().Where(entorderbenefit.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entorderbenefit.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderbenefit.EntID(*h.EntID))
	}
	h.stmSelect = h.selectAccount(stm)
	return nil
}

func (h *queryHandler) queryAccounts(cli *ent.Client) (*ent.OrderBenefitSelect, error) {
	stm, err := orderbenefitcrud.SetQueryConds(cli.OrderBenefit.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAccount(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderbenefit.Table)
	s.AppendSelect(
		t.C(entorderbenefit.FieldID),
		t.C(entorderbenefit.FieldEntID),
		t.C(entorderbenefit.FieldAppID),
		t.C(entorderbenefit.FieldUserID),
		t.C(entorderbenefit.FieldCoinTypeID),
		t.C(entorderbenefit.FieldAccountID),
		t.C(entorderbenefit.FieldOrderID),
		t.C(entorderbenefit.FieldCreatedAt),
		t.C(entorderbenefit.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoinAccount(s *sql.Selector) error {
	t := sql.Table(entaccount.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbenefit.FieldAccountID),
			t.C(entaccount.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entaccount.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid orderbenefit active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid orderbenefit blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
		)
	}
	if h.Conds != nil && h.Conds.Address != nil {
		addr, ok := h.Conds.Address.Val.(string)
		if !ok {
			return fmt.Errorf("invalid orderbenefit address")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldAddress), addr),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entaccount.FieldAddress), "address"),
		sql.As(t.C(entaccount.FieldActive), "active"),
		sql.As(t.C(entaccount.FieldBlocked), "blocked"),
		sql.As(t.C(entaccount.FieldUsedFor), "used_for"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinAccount(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinAccount(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		if _, err := uuid.Parse(info.CoinTypeID); err != nil {
			info.CoinTypeID = uuid.Nil.String()
		}
	}
}

func (h *Handler) GetAccount(ctx context.Context) (*npool.Account, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccount(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
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

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAccounts(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryAccounts(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entorderbenefit.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
