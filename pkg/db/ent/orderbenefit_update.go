// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/orderbenefit"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderBenefitUpdate is the builder for updating OrderBenefit entities.
type OrderBenefitUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderBenefitMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderBenefitUpdate builder.
func (obu *OrderBenefitUpdate) Where(ps ...predicate.OrderBenefit) *OrderBenefitUpdate {
	obu.mutation.Where(ps...)
	return obu
}

// SetCreatedAt sets the "created_at" field.
func (obu *OrderBenefitUpdate) SetCreatedAt(u uint32) *OrderBenefitUpdate {
	obu.mutation.ResetCreatedAt()
	obu.mutation.SetCreatedAt(u)
	return obu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableCreatedAt(u *uint32) *OrderBenefitUpdate {
	if u != nil {
		obu.SetCreatedAt(*u)
	}
	return obu
}

// AddCreatedAt adds u to the "created_at" field.
func (obu *OrderBenefitUpdate) AddCreatedAt(u int32) *OrderBenefitUpdate {
	obu.mutation.AddCreatedAt(u)
	return obu
}

// SetUpdatedAt sets the "updated_at" field.
func (obu *OrderBenefitUpdate) SetUpdatedAt(u uint32) *OrderBenefitUpdate {
	obu.mutation.ResetUpdatedAt()
	obu.mutation.SetUpdatedAt(u)
	return obu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (obu *OrderBenefitUpdate) AddUpdatedAt(u int32) *OrderBenefitUpdate {
	obu.mutation.AddUpdatedAt(u)
	return obu
}

// SetDeletedAt sets the "deleted_at" field.
func (obu *OrderBenefitUpdate) SetDeletedAt(u uint32) *OrderBenefitUpdate {
	obu.mutation.ResetDeletedAt()
	obu.mutation.SetDeletedAt(u)
	return obu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableDeletedAt(u *uint32) *OrderBenefitUpdate {
	if u != nil {
		obu.SetDeletedAt(*u)
	}
	return obu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (obu *OrderBenefitUpdate) AddDeletedAt(u int32) *OrderBenefitUpdate {
	obu.mutation.AddDeletedAt(u)
	return obu
}

// SetEntID sets the "ent_id" field.
func (obu *OrderBenefitUpdate) SetEntID(u uuid.UUID) *OrderBenefitUpdate {
	obu.mutation.SetEntID(u)
	return obu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableEntID(u *uuid.UUID) *OrderBenefitUpdate {
	if u != nil {
		obu.SetEntID(*u)
	}
	return obu
}

// SetAppID sets the "app_id" field.
func (obu *OrderBenefitUpdate) SetAppID(u uuid.UUID) *OrderBenefitUpdate {
	obu.mutation.SetAppID(u)
	return obu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableAppID(u *uuid.UUID) *OrderBenefitUpdate {
	if u != nil {
		obu.SetAppID(*u)
	}
	return obu
}

// ClearAppID clears the value of the "app_id" field.
func (obu *OrderBenefitUpdate) ClearAppID() *OrderBenefitUpdate {
	obu.mutation.ClearAppID()
	return obu
}

// SetUserID sets the "user_id" field.
func (obu *OrderBenefitUpdate) SetUserID(u uuid.UUID) *OrderBenefitUpdate {
	obu.mutation.SetUserID(u)
	return obu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableUserID(u *uuid.UUID) *OrderBenefitUpdate {
	if u != nil {
		obu.SetUserID(*u)
	}
	return obu
}

// ClearUserID clears the value of the "user_id" field.
func (obu *OrderBenefitUpdate) ClearUserID() *OrderBenefitUpdate {
	obu.mutation.ClearUserID()
	return obu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (obu *OrderBenefitUpdate) SetCoinTypeID(u uuid.UUID) *OrderBenefitUpdate {
	obu.mutation.SetCoinTypeID(u)
	return obu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableCoinTypeID(u *uuid.UUID) *OrderBenefitUpdate {
	if u != nil {
		obu.SetCoinTypeID(*u)
	}
	return obu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (obu *OrderBenefitUpdate) ClearCoinTypeID() *OrderBenefitUpdate {
	obu.mutation.ClearCoinTypeID()
	return obu
}

// SetAccountID sets the "account_id" field.
func (obu *OrderBenefitUpdate) SetAccountID(u uuid.UUID) *OrderBenefitUpdate {
	obu.mutation.SetAccountID(u)
	return obu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableAccountID(u *uuid.UUID) *OrderBenefitUpdate {
	if u != nil {
		obu.SetAccountID(*u)
	}
	return obu
}

// ClearAccountID clears the value of the "account_id" field.
func (obu *OrderBenefitUpdate) ClearAccountID() *OrderBenefitUpdate {
	obu.mutation.ClearAccountID()
	return obu
}

// SetOrderID sets the "order_id" field.
func (obu *OrderBenefitUpdate) SetOrderID(u uuid.UUID) *OrderBenefitUpdate {
	obu.mutation.SetOrderID(u)
	return obu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (obu *OrderBenefitUpdate) SetNillableOrderID(u *uuid.UUID) *OrderBenefitUpdate {
	if u != nil {
		obu.SetOrderID(*u)
	}
	return obu
}

// ClearOrderID clears the value of the "order_id" field.
func (obu *OrderBenefitUpdate) ClearOrderID() *OrderBenefitUpdate {
	obu.mutation.ClearOrderID()
	return obu
}

// Mutation returns the OrderBenefitMutation object of the builder.
func (obu *OrderBenefitUpdate) Mutation() *OrderBenefitMutation {
	return obu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (obu *OrderBenefitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := obu.defaults(); err != nil {
		return 0, err
	}
	if len(obu.hooks) == 0 {
		affected, err = obu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			obu.mutation = mutation
			affected, err = obu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(obu.hooks) - 1; i >= 0; i-- {
			if obu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = obu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, obu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (obu *OrderBenefitUpdate) SaveX(ctx context.Context) int {
	affected, err := obu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (obu *OrderBenefitUpdate) Exec(ctx context.Context) error {
	_, err := obu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (obu *OrderBenefitUpdate) ExecX(ctx context.Context) {
	if err := obu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (obu *OrderBenefitUpdate) defaults() error {
	if _, ok := obu.mutation.UpdatedAt(); !ok {
		if orderbenefit.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderbenefit.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderbenefit.UpdateDefaultUpdatedAt()
		obu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (obu *OrderBenefitUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderBenefitUpdate {
	obu.modifiers = append(obu.modifiers, modifiers...)
	return obu
}

func (obu *OrderBenefitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderbenefit.Table,
			Columns: orderbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderbenefit.FieldID,
			},
		},
	}
	if ps := obu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := obu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldCreatedAt,
		})
	}
	if value, ok := obu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldCreatedAt,
		})
	}
	if value, ok := obu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := obu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := obu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldDeletedAt,
		})
	}
	if value, ok := obu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldDeletedAt,
		})
	}
	if value, ok := obu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldEntID,
		})
	}
	if value, ok := obu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldAppID,
		})
	}
	if obu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldAppID,
		})
	}
	if value, ok := obu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldUserID,
		})
	}
	if obu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldUserID,
		})
	}
	if value, ok := obu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldCoinTypeID,
		})
	}
	if obu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldCoinTypeID,
		})
	}
	if value, ok := obu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldAccountID,
		})
	}
	if obu.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldAccountID,
		})
	}
	if value, ok := obu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldOrderID,
		})
	}
	if obu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldOrderID,
		})
	}
	_spec.Modifiers = obu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, obu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderbenefit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderBenefitUpdateOne is the builder for updating a single OrderBenefit entity.
type OrderBenefitUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderBenefitMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (obuo *OrderBenefitUpdateOne) SetCreatedAt(u uint32) *OrderBenefitUpdateOne {
	obuo.mutation.ResetCreatedAt()
	obuo.mutation.SetCreatedAt(u)
	return obuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableCreatedAt(u *uint32) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetCreatedAt(*u)
	}
	return obuo
}

// AddCreatedAt adds u to the "created_at" field.
func (obuo *OrderBenefitUpdateOne) AddCreatedAt(u int32) *OrderBenefitUpdateOne {
	obuo.mutation.AddCreatedAt(u)
	return obuo
}

// SetUpdatedAt sets the "updated_at" field.
func (obuo *OrderBenefitUpdateOne) SetUpdatedAt(u uint32) *OrderBenefitUpdateOne {
	obuo.mutation.ResetUpdatedAt()
	obuo.mutation.SetUpdatedAt(u)
	return obuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (obuo *OrderBenefitUpdateOne) AddUpdatedAt(u int32) *OrderBenefitUpdateOne {
	obuo.mutation.AddUpdatedAt(u)
	return obuo
}

// SetDeletedAt sets the "deleted_at" field.
func (obuo *OrderBenefitUpdateOne) SetDeletedAt(u uint32) *OrderBenefitUpdateOne {
	obuo.mutation.ResetDeletedAt()
	obuo.mutation.SetDeletedAt(u)
	return obuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableDeletedAt(u *uint32) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetDeletedAt(*u)
	}
	return obuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (obuo *OrderBenefitUpdateOne) AddDeletedAt(u int32) *OrderBenefitUpdateOne {
	obuo.mutation.AddDeletedAt(u)
	return obuo
}

// SetEntID sets the "ent_id" field.
func (obuo *OrderBenefitUpdateOne) SetEntID(u uuid.UUID) *OrderBenefitUpdateOne {
	obuo.mutation.SetEntID(u)
	return obuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetEntID(*u)
	}
	return obuo
}

// SetAppID sets the "app_id" field.
func (obuo *OrderBenefitUpdateOne) SetAppID(u uuid.UUID) *OrderBenefitUpdateOne {
	obuo.mutation.SetAppID(u)
	return obuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableAppID(u *uuid.UUID) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetAppID(*u)
	}
	return obuo
}

// ClearAppID clears the value of the "app_id" field.
func (obuo *OrderBenefitUpdateOne) ClearAppID() *OrderBenefitUpdateOne {
	obuo.mutation.ClearAppID()
	return obuo
}

// SetUserID sets the "user_id" field.
func (obuo *OrderBenefitUpdateOne) SetUserID(u uuid.UUID) *OrderBenefitUpdateOne {
	obuo.mutation.SetUserID(u)
	return obuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableUserID(u *uuid.UUID) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetUserID(*u)
	}
	return obuo
}

// ClearUserID clears the value of the "user_id" field.
func (obuo *OrderBenefitUpdateOne) ClearUserID() *OrderBenefitUpdateOne {
	obuo.mutation.ClearUserID()
	return obuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (obuo *OrderBenefitUpdateOne) SetCoinTypeID(u uuid.UUID) *OrderBenefitUpdateOne {
	obuo.mutation.SetCoinTypeID(u)
	return obuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetCoinTypeID(*u)
	}
	return obuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (obuo *OrderBenefitUpdateOne) ClearCoinTypeID() *OrderBenefitUpdateOne {
	obuo.mutation.ClearCoinTypeID()
	return obuo
}

// SetAccountID sets the "account_id" field.
func (obuo *OrderBenefitUpdateOne) SetAccountID(u uuid.UUID) *OrderBenefitUpdateOne {
	obuo.mutation.SetAccountID(u)
	return obuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableAccountID(u *uuid.UUID) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetAccountID(*u)
	}
	return obuo
}

// ClearAccountID clears the value of the "account_id" field.
func (obuo *OrderBenefitUpdateOne) ClearAccountID() *OrderBenefitUpdateOne {
	obuo.mutation.ClearAccountID()
	return obuo
}

// SetOrderID sets the "order_id" field.
func (obuo *OrderBenefitUpdateOne) SetOrderID(u uuid.UUID) *OrderBenefitUpdateOne {
	obuo.mutation.SetOrderID(u)
	return obuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (obuo *OrderBenefitUpdateOne) SetNillableOrderID(u *uuid.UUID) *OrderBenefitUpdateOne {
	if u != nil {
		obuo.SetOrderID(*u)
	}
	return obuo
}

// ClearOrderID clears the value of the "order_id" field.
func (obuo *OrderBenefitUpdateOne) ClearOrderID() *OrderBenefitUpdateOne {
	obuo.mutation.ClearOrderID()
	return obuo
}

// Mutation returns the OrderBenefitMutation object of the builder.
func (obuo *OrderBenefitUpdateOne) Mutation() *OrderBenefitMutation {
	return obuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (obuo *OrderBenefitUpdateOne) Select(field string, fields ...string) *OrderBenefitUpdateOne {
	obuo.fields = append([]string{field}, fields...)
	return obuo
}

// Save executes the query and returns the updated OrderBenefit entity.
func (obuo *OrderBenefitUpdateOne) Save(ctx context.Context) (*OrderBenefit, error) {
	var (
		err  error
		node *OrderBenefit
	)
	if err := obuo.defaults(); err != nil {
		return nil, err
	}
	if len(obuo.hooks) == 0 {
		node, err = obuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderBenefitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			obuo.mutation = mutation
			node, err = obuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(obuo.hooks) - 1; i >= 0; i-- {
			if obuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = obuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, obuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderBenefit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderBenefitMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (obuo *OrderBenefitUpdateOne) SaveX(ctx context.Context) *OrderBenefit {
	node, err := obuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (obuo *OrderBenefitUpdateOne) Exec(ctx context.Context) error {
	_, err := obuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (obuo *OrderBenefitUpdateOne) ExecX(ctx context.Context) {
	if err := obuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (obuo *OrderBenefitUpdateOne) defaults() error {
	if _, ok := obuo.mutation.UpdatedAt(); !ok {
		if orderbenefit.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderbenefit.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderbenefit.UpdateDefaultUpdatedAt()
		obuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (obuo *OrderBenefitUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderBenefitUpdateOne {
	obuo.modifiers = append(obuo.modifiers, modifiers...)
	return obuo
}

func (obuo *OrderBenefitUpdateOne) sqlSave(ctx context.Context) (_node *OrderBenefit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderbenefit.Table,
			Columns: orderbenefit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderbenefit.FieldID,
			},
		},
	}
	id, ok := obuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderBenefit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := obuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderbenefit.FieldID)
		for _, f := range fields {
			if !orderbenefit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderbenefit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := obuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := obuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldCreatedAt,
		})
	}
	if value, ok := obuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldCreatedAt,
		})
	}
	if value, ok := obuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := obuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldUpdatedAt,
		})
	}
	if value, ok := obuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldDeletedAt,
		})
	}
	if value, ok := obuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderbenefit.FieldDeletedAt,
		})
	}
	if value, ok := obuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldEntID,
		})
	}
	if value, ok := obuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldAppID,
		})
	}
	if obuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldAppID,
		})
	}
	if value, ok := obuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldUserID,
		})
	}
	if obuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldUserID,
		})
	}
	if value, ok := obuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldCoinTypeID,
		})
	}
	if obuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldCoinTypeID,
		})
	}
	if value, ok := obuo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldAccountID,
		})
	}
	if obuo.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldAccountID,
		})
	}
	if value, ok := obuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderbenefit.FieldOrderID,
		})
	}
	if obuo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderbenefit.FieldOrderID,
		})
	}
	_spec.Modifiers = obuo.modifiers
	_node = &OrderBenefit{config: obuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, obuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderbenefit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
