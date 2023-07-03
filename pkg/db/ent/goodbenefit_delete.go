// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/predicate"
)

// GoodBenefitDelete is the builder for deleting a GoodBenefit entity.
type GoodBenefitDelete struct {
	config
	hooks    []Hook
	mutation *GoodBenefitMutation
}

// Where appends a list predicates to the GoodBenefitDelete builder.
func (gbd *GoodBenefitDelete) Where(ps ...predicate.GoodBenefit) *GoodBenefitDelete {
	gbd.mutation.Where(ps...)
	return gbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gbd *GoodBenefitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gbd.hooks) == 0 {
		affected, err = gbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gbd.mutation = mutation
			affected, err = gbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gbd.hooks) - 1; i >= 0; i-- {
			if gbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbd *GoodBenefitDelete) ExecX(ctx context.Context) int {
	n, err := gbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gbd *GoodBenefitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodbenefit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodbenefit.FieldID,
			},
		},
	}
	if ps := gbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GoodBenefitDeleteOne is the builder for deleting a single GoodBenefit entity.
type GoodBenefitDeleteOne struct {
	gbd *GoodBenefitDelete
}

// Exec executes the deletion query.
func (gbdo *GoodBenefitDeleteOne) Exec(ctx context.Context) error {
	n, err := gbdo.gbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodbenefit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gbdo *GoodBenefitDeleteOne) ExecX(ctx context.Context) {
	gbdo.gbd.ExecX(ctx)
}
