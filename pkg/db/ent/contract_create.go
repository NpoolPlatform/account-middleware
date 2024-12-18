// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/contract"
	"github.com/google/uuid"
)

// ContractCreate is the builder for creating a Contract entity.
type ContractCreate struct {
	config
	mutation *ContractMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContractCreate) SetCreatedAt(u uint32) *ContractCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableCreatedAt(u *uint32) *ContractCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContractCreate) SetUpdatedAt(u uint32) *ContractCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableUpdatedAt(u *uint32) *ContractCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ContractCreate) SetDeletedAt(u uint32) *ContractCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableDeletedAt(u *uint32) *ContractCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetEntID sets the "ent_id" field.
func (cc *ContractCreate) SetEntID(u uuid.UUID) *ContractCreate {
	cc.mutation.SetEntID(u)
	return cc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cc *ContractCreate) SetNillableEntID(u *uuid.UUID) *ContractCreate {
	if u != nil {
		cc.SetEntID(*u)
	}
	return cc
}

// SetGoodID sets the "good_id" field.
func (cc *ContractCreate) SetGoodID(u uuid.UUID) *ContractCreate {
	cc.mutation.SetGoodID(u)
	return cc
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (cc *ContractCreate) SetNillableGoodID(u *uuid.UUID) *ContractCreate {
	if u != nil {
		cc.SetGoodID(*u)
	}
	return cc
}

// SetPledgeID sets the "pledge_id" field.
func (cc *ContractCreate) SetPledgeID(u uuid.UUID) *ContractCreate {
	cc.mutation.SetPledgeID(u)
	return cc
}

// SetNillablePledgeID sets the "pledge_id" field if the given value is not nil.
func (cc *ContractCreate) SetNillablePledgeID(u *uuid.UUID) *ContractCreate {
	if u != nil {
		cc.SetPledgeID(*u)
	}
	return cc
}

// SetAccountID sets the "account_id" field.
func (cc *ContractCreate) SetAccountID(u uuid.UUID) *ContractCreate {
	cc.mutation.SetAccountID(u)
	return cc
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (cc *ContractCreate) SetNillableAccountID(u *uuid.UUID) *ContractCreate {
	if u != nil {
		cc.SetAccountID(*u)
	}
	return cc
}

// SetBackup sets the "backup" field.
func (cc *ContractCreate) SetBackup(b bool) *ContractCreate {
	cc.mutation.SetBackup(b)
	return cc
}

// SetNillableBackup sets the "backup" field if the given value is not nil.
func (cc *ContractCreate) SetNillableBackup(b *bool) *ContractCreate {
	if b != nil {
		cc.SetBackup(*b)
	}
	return cc
}

// SetContractType sets the "contract_type" field.
func (cc *ContractCreate) SetContractType(s string) *ContractCreate {
	cc.mutation.SetContractType(s)
	return cc
}

// SetNillableContractType sets the "contract_type" field if the given value is not nil.
func (cc *ContractCreate) SetNillableContractType(s *string) *ContractCreate {
	if s != nil {
		cc.SetContractType(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ContractCreate) SetID(u uint32) *ContractCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the ContractMutation object of the builder.
func (cc *ContractCreate) Mutation() *ContractMutation {
	return cc.mutation
}

// Save creates the Contract in the database.
func (cc *ContractCreate) Save(ctx context.Context) (*Contract, error) {
	var (
		err  error
		node *Contract
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Contract)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContractMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContractCreate) SaveX(ctx context.Context) *Contract {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContractCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContractCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContractCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if contract.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := contract.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if contract.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := contract.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if contract.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := contract.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.EntID(); !ok {
		if contract.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := contract.DefaultEntID()
		cc.mutation.SetEntID(v)
	}
	if _, ok := cc.mutation.GoodID(); !ok {
		if contract.DefaultGoodID == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultGoodID (forgotten import ent/runtime?)")
		}
		v := contract.DefaultGoodID()
		cc.mutation.SetGoodID(v)
	}
	if _, ok := cc.mutation.PledgeID(); !ok {
		if contract.DefaultPledgeID == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultPledgeID (forgotten import ent/runtime?)")
		}
		v := contract.DefaultPledgeID()
		cc.mutation.SetPledgeID(v)
	}
	if _, ok := cc.mutation.AccountID(); !ok {
		if contract.DefaultAccountID == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultAccountID (forgotten import ent/runtime?)")
		}
		v := contract.DefaultAccountID()
		cc.mutation.SetAccountID(v)
	}
	if _, ok := cc.mutation.Backup(); !ok {
		v := contract.DefaultBackup
		cc.mutation.SetBackup(v)
	}
	if _, ok := cc.mutation.ContractType(); !ok {
		v := contract.DefaultContractType
		cc.mutation.SetContractType(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContractCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Contract.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Contract.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Contract.deleted_at"`)}
	}
	if _, ok := cc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Contract.ent_id"`)}
	}
	return nil
}

func (cc *ContractCreate) sqlSave(ctx context.Context) (*Contract, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (cc *ContractCreate) createSpec() (*Contract, *sqlgraph.CreateSpec) {
	var (
		_node = &Contract{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: contract.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: contract.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: contract.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := cc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: contract.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := cc.mutation.PledgeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: contract.FieldPledgeID,
		})
		_node.PledgeID = value
	}
	if value, ok := cc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: contract.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := cc.mutation.Backup(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: contract.FieldBackup,
		})
		_node.Backup = value
	}
	if value, ok := cc.mutation.ContractType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldContractType,
		})
		_node.ContractType = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Contract.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContractUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *ContractCreate) OnConflict(opts ...sql.ConflictOption) *ContractUpsertOne {
	cc.conflict = opts
	return &ContractUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *ContractCreate) OnConflictColumns(columns ...string) *ContractUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ContractUpsertOne{
		create: cc,
	}
}

type (
	// ContractUpsertOne is the builder for "upsert"-ing
	//  one Contract node.
	ContractUpsertOne struct {
		create *ContractCreate
	}

	// ContractUpsert is the "OnConflict" setter.
	ContractUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ContractUpsert) SetCreatedAt(v uint32) *ContractUpsert {
	u.Set(contract.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ContractUpsert) UpdateCreatedAt() *ContractUpsert {
	u.SetExcluded(contract.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ContractUpsert) AddCreatedAt(v uint32) *ContractUpsert {
	u.Add(contract.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContractUpsert) SetUpdatedAt(v uint32) *ContractUpsert {
	u.Set(contract.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContractUpsert) UpdateUpdatedAt() *ContractUpsert {
	u.SetExcluded(contract.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ContractUpsert) AddUpdatedAt(v uint32) *ContractUpsert {
	u.Add(contract.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ContractUpsert) SetDeletedAt(v uint32) *ContractUpsert {
	u.Set(contract.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ContractUpsert) UpdateDeletedAt() *ContractUpsert {
	u.SetExcluded(contract.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ContractUpsert) AddDeletedAt(v uint32) *ContractUpsert {
	u.Add(contract.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *ContractUpsert) SetEntID(v uuid.UUID) *ContractUpsert {
	u.Set(contract.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ContractUpsert) UpdateEntID() *ContractUpsert {
	u.SetExcluded(contract.FieldEntID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *ContractUpsert) SetGoodID(v uuid.UUID) *ContractUpsert {
	u.Set(contract.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ContractUpsert) UpdateGoodID() *ContractUpsert {
	u.SetExcluded(contract.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *ContractUpsert) ClearGoodID() *ContractUpsert {
	u.SetNull(contract.FieldGoodID)
	return u
}

// SetPledgeID sets the "pledge_id" field.
func (u *ContractUpsert) SetPledgeID(v uuid.UUID) *ContractUpsert {
	u.Set(contract.FieldPledgeID, v)
	return u
}

// UpdatePledgeID sets the "pledge_id" field to the value that was provided on create.
func (u *ContractUpsert) UpdatePledgeID() *ContractUpsert {
	u.SetExcluded(contract.FieldPledgeID)
	return u
}

// ClearPledgeID clears the value of the "pledge_id" field.
func (u *ContractUpsert) ClearPledgeID() *ContractUpsert {
	u.SetNull(contract.FieldPledgeID)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *ContractUpsert) SetAccountID(v uuid.UUID) *ContractUpsert {
	u.Set(contract.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *ContractUpsert) UpdateAccountID() *ContractUpsert {
	u.SetExcluded(contract.FieldAccountID)
	return u
}

// ClearAccountID clears the value of the "account_id" field.
func (u *ContractUpsert) ClearAccountID() *ContractUpsert {
	u.SetNull(contract.FieldAccountID)
	return u
}

// SetBackup sets the "backup" field.
func (u *ContractUpsert) SetBackup(v bool) *ContractUpsert {
	u.Set(contract.FieldBackup, v)
	return u
}

// UpdateBackup sets the "backup" field to the value that was provided on create.
func (u *ContractUpsert) UpdateBackup() *ContractUpsert {
	u.SetExcluded(contract.FieldBackup)
	return u
}

// ClearBackup clears the value of the "backup" field.
func (u *ContractUpsert) ClearBackup() *ContractUpsert {
	u.SetNull(contract.FieldBackup)
	return u
}

// SetContractType sets the "contract_type" field.
func (u *ContractUpsert) SetContractType(v string) *ContractUpsert {
	u.Set(contract.FieldContractType, v)
	return u
}

// UpdateContractType sets the "contract_type" field to the value that was provided on create.
func (u *ContractUpsert) UpdateContractType() *ContractUpsert {
	u.SetExcluded(contract.FieldContractType)
	return u
}

// ClearContractType clears the value of the "contract_type" field.
func (u *ContractUpsert) ClearContractType() *ContractUpsert {
	u.SetNull(contract.FieldContractType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contract.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ContractUpsertOne) UpdateNewValues() *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contract.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Contract.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ContractUpsertOne) Ignore() *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContractUpsertOne) DoNothing() *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContractCreate.OnConflict
// documentation for more info.
func (u *ContractUpsertOne) Update(set func(*ContractUpsert)) *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContractUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ContractUpsertOne) SetCreatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ContractUpsertOne) AddCreatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateCreatedAt() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContractUpsertOne) SetUpdatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ContractUpsertOne) AddUpdatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateUpdatedAt() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ContractUpsertOne) SetDeletedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ContractUpsertOne) AddDeletedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateDeletedAt() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *ContractUpsertOne) SetEntID(v uuid.UUID) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateEntID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateEntID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *ContractUpsertOne) SetGoodID(v uuid.UUID) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateGoodID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *ContractUpsertOne) ClearGoodID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.ClearGoodID()
	})
}

// SetPledgeID sets the "pledge_id" field.
func (u *ContractUpsertOne) SetPledgeID(v uuid.UUID) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetPledgeID(v)
	})
}

// UpdatePledgeID sets the "pledge_id" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdatePledgeID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdatePledgeID()
	})
}

// ClearPledgeID clears the value of the "pledge_id" field.
func (u *ContractUpsertOne) ClearPledgeID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.ClearPledgeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *ContractUpsertOne) SetAccountID(v uuid.UUID) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateAccountID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateAccountID()
	})
}

// ClearAccountID clears the value of the "account_id" field.
func (u *ContractUpsertOne) ClearAccountID() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.ClearAccountID()
	})
}

// SetBackup sets the "backup" field.
func (u *ContractUpsertOne) SetBackup(v bool) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetBackup(v)
	})
}

// UpdateBackup sets the "backup" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateBackup() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateBackup()
	})
}

// ClearBackup clears the value of the "backup" field.
func (u *ContractUpsertOne) ClearBackup() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.ClearBackup()
	})
}

// SetContractType sets the "contract_type" field.
func (u *ContractUpsertOne) SetContractType(v string) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetContractType(v)
	})
}

// UpdateContractType sets the "contract_type" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateContractType() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateContractType()
	})
}

// ClearContractType clears the value of the "contract_type" field.
func (u *ContractUpsertOne) ClearContractType() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.ClearContractType()
	})
}

// Exec executes the query.
func (u *ContractUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContractCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContractUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContractUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContractUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContractCreateBulk is the builder for creating many Contract entities in bulk.
type ContractCreateBulk struct {
	config
	builders []*ContractCreate
	conflict []sql.ConflictOption
}

// Save creates the Contract entities in the database.
func (ccb *ContractCreateBulk) Save(ctx context.Context) ([]*Contract, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contract, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContractMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContractCreateBulk) SaveX(ctx context.Context) []*Contract {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContractCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContractCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Contract.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContractUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *ContractCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContractUpsertBulk {
	ccb.conflict = opts
	return &ContractUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *ContractCreateBulk) OnConflictColumns(columns ...string) *ContractUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ContractUpsertBulk{
		create: ccb,
	}
}

// ContractUpsertBulk is the builder for "upsert"-ing
// a bulk of Contract nodes.
type ContractUpsertBulk struct {
	create *ContractCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contract.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ContractUpsertBulk) UpdateNewValues() *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contract.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ContractUpsertBulk) Ignore() *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContractUpsertBulk) DoNothing() *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContractCreateBulk.OnConflict
// documentation for more info.
func (u *ContractUpsertBulk) Update(set func(*ContractUpsert)) *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContractUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ContractUpsertBulk) SetCreatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ContractUpsertBulk) AddCreatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateCreatedAt() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContractUpsertBulk) SetUpdatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ContractUpsertBulk) AddUpdatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateUpdatedAt() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ContractUpsertBulk) SetDeletedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ContractUpsertBulk) AddDeletedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateDeletedAt() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *ContractUpsertBulk) SetEntID(v uuid.UUID) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateEntID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateEntID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *ContractUpsertBulk) SetGoodID(v uuid.UUID) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateGoodID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *ContractUpsertBulk) ClearGoodID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.ClearGoodID()
	})
}

// SetPledgeID sets the "pledge_id" field.
func (u *ContractUpsertBulk) SetPledgeID(v uuid.UUID) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetPledgeID(v)
	})
}

// UpdatePledgeID sets the "pledge_id" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdatePledgeID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdatePledgeID()
	})
}

// ClearPledgeID clears the value of the "pledge_id" field.
func (u *ContractUpsertBulk) ClearPledgeID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.ClearPledgeID()
	})
}

// SetAccountID sets the "account_id" field.
func (u *ContractUpsertBulk) SetAccountID(v uuid.UUID) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateAccountID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateAccountID()
	})
}

// ClearAccountID clears the value of the "account_id" field.
func (u *ContractUpsertBulk) ClearAccountID() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.ClearAccountID()
	})
}

// SetBackup sets the "backup" field.
func (u *ContractUpsertBulk) SetBackup(v bool) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetBackup(v)
	})
}

// UpdateBackup sets the "backup" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateBackup() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateBackup()
	})
}

// ClearBackup clears the value of the "backup" field.
func (u *ContractUpsertBulk) ClearBackup() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.ClearBackup()
	})
}

// SetContractType sets the "contract_type" field.
func (u *ContractUpsertBulk) SetContractType(v string) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetContractType(v)
	})
}

// UpdateContractType sets the "contract_type" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateContractType() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateContractType()
	})
}

// ClearContractType clears the value of the "contract_type" field.
func (u *ContractUpsertBulk) ClearContractType() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.ClearContractType()
	})
}

// Exec executes the query.
func (u *ContractUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContractCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContractCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContractUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}