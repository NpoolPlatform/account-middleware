// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/orderbenefit"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/predicate"
)

// OrderBenefitDelete is the builder for deleting a OrderBenefit entity.
type OrderBenefitDelete struct {
	config
	hooks    []Hook
	mutation *OrderBenefitMutation
}

// Where appends a list predicates to the OrderBenefitDelete builder.
func (obd *OrderBenefitDelete) Where(ps ...predicate.OrderBenefit) *OrderBenefitDelete {
	obd.mutation.Where(ps...)
	return obd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (obd *OrderBenefitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(obd.hooks) == 0 {
		affected, err = obd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			obd.mutation = mutation
			affected, err = obd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(obd.hooks) - 1; i >= 0; i-- {
			if obd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = obd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, obd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (obd *OrderBenefitDelete) ExecX(ctx context.Context) int {
	n, err := obd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (obd *OrderBenefitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderbenefit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderbenefit.FieldID,
			},
		},
	}
	if ps := obd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, obd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrderBenefitDeleteOne is the builder for deleting a single OrderBenefit entity.
type OrderBenefitDeleteOne struct {
	obd *OrderBenefitDelete
}

// Exec executes the deletion query.
func (obdo *OrderBenefitDeleteOne) Exec(ctx context.Context) error {
	n, err := obdo.obd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderbenefit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (obdo *OrderBenefitDeleteOne) ExecX(ctx context.Context) {
	obdo.obd.ExecX(ctx)
}