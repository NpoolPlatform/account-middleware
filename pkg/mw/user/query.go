package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/account-middleware/pkg/db"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-middleware/pkg/db/ent/account"
	entuser "github.com/NpoolPlatform/account-middleware/pkg/db/ent/user"

	usercrud "github.com/NpoolPlatform/account-middleware/pkg/crud/user"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.UserSelect
	infos []*npool.Account
	total uint32
}

func (h *queryHandler) selectAccount(stm *ent.UserQuery) {
	h.stm = stm.Select(
		entuser.FieldID,
		entuser.FieldAppID,
		entuser.FieldUserID,
		entuser.FieldCoinTypeID,
		entuser.FieldAccountID,
		entuser.FieldUsedFor,
		entuser.FieldLabels,
		entuser.FieldMemo,
		entuser.FieldCreatedAt,
		entuser.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAccount(cli *ent.Client) {
	h.selectAccount(
		cli.User.
			Query().
			Where(
				entuser.ID(*h.ID),
				entuser.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryAccounts(cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.User.Query(), h.Conds)
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
			s.C(entuser.FieldAccountID),
			t.C(entaccount.FieldID),
		)

	if h.Conds != nil && h.Conds.Active != nil {
		active, ok := h.Conds.Active.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid user active")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldActive), active),
		)
	}
	if h.Conds != nil && h.Conds.Blocked != nil {
		blocked, ok := h.Conds.Blocked.Val.(bool)
		if !ok {
			return fmt.Errorf("invalid user blocked")
		}
		s.Where(
			sql.EQ(t.C(entaccount.FieldBlocked), blocked),
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
	h.stm.Modify(func(s *sql.Selector) {
		err = h.queryJoinAccount(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		if _, err := uuid.Parse(info.CoinTypeID); err != nil {
			info.UsedFor = basetypes.AccountUsedFor(basetypes.AccountUsedFor_value[info.UsedForStr])
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

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stm.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entuser.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetAccountOnly(ctx context.Context) (*npool.Account, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.
			Offset(0).
			Limit(singleRowLimit).
			Order(ent.Desc(entuser.FieldCreatedAt))

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
