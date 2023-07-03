// Code generated by ent, DO NOT EDIT.

package account

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// UsedFor applies equality check predicate on the "used_for" field. It's identical to UsedForEQ.
func UsedFor(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// PlatformHoldPrivateKey applies equality check predicate on the "platform_hold_private_key" field. It's identical to PlatformHoldPrivateKeyEQ.
func PlatformHoldPrivateKey(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformHoldPrivateKey), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// LockedBy applies equality check predicate on the "locked_by" field. It's identical to LockedByEQ.
func LockedBy(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockedBy), v))
	})
}

// Blocked applies equality check predicate on the "blocked" field. It's identical to BlockedEQ.
func Blocked(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlocked), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAddress)))
	})
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAddress)))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// UsedForEQ applies the EQ predicate on the "used_for" field.
func UsedForEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// UsedForNEQ applies the NEQ predicate on the "used_for" field.
func UsedForNEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedFor), v))
	})
}

// UsedForIn applies the In predicate on the "used_for" field.
func UsedForIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsedFor), v...))
	})
}

// UsedForNotIn applies the NotIn predicate on the "used_for" field.
func UsedForNotIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsedFor), v...))
	})
}

// UsedForGT applies the GT predicate on the "used_for" field.
func UsedForGT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedFor), v))
	})
}

// UsedForGTE applies the GTE predicate on the "used_for" field.
func UsedForGTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedFor), v))
	})
}

// UsedForLT applies the LT predicate on the "used_for" field.
func UsedForLT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedFor), v))
	})
}

// UsedForLTE applies the LTE predicate on the "used_for" field.
func UsedForLTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedFor), v))
	})
}

// UsedForContains applies the Contains predicate on the "used_for" field.
func UsedForContains(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsedFor), v))
	})
}

// UsedForHasPrefix applies the HasPrefix predicate on the "used_for" field.
func UsedForHasPrefix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsedFor), v))
	})
}

// UsedForHasSuffix applies the HasSuffix predicate on the "used_for" field.
func UsedForHasSuffix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsedFor), v))
	})
}

// UsedForIsNil applies the IsNil predicate on the "used_for" field.
func UsedForIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsedFor)))
	})
}

// UsedForNotNil applies the NotNil predicate on the "used_for" field.
func UsedForNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsedFor)))
	})
}

// UsedForEqualFold applies the EqualFold predicate on the "used_for" field.
func UsedForEqualFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsedFor), v))
	})
}

// UsedForContainsFold applies the ContainsFold predicate on the "used_for" field.
func UsedForContainsFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsedFor), v))
	})
}

// PlatformHoldPrivateKeyEQ applies the EQ predicate on the "platform_hold_private_key" field.
func PlatformHoldPrivateKeyEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformHoldPrivateKey), v))
	})
}

// PlatformHoldPrivateKeyNEQ applies the NEQ predicate on the "platform_hold_private_key" field.
func PlatformHoldPrivateKeyNEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlatformHoldPrivateKey), v))
	})
}

// PlatformHoldPrivateKeyIsNil applies the IsNil predicate on the "platform_hold_private_key" field.
func PlatformHoldPrivateKeyIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPlatformHoldPrivateKey)))
	})
}

// PlatformHoldPrivateKeyNotNil applies the NotNil predicate on the "platform_hold_private_key" field.
func PlatformHoldPrivateKeyNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPlatformHoldPrivateKey)))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// ActiveIsNil applies the IsNil predicate on the "active" field.
func ActiveIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldActive)))
	})
}

// ActiveNotNil applies the NotNil predicate on the "active" field.
func ActiveNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldActive)))
	})
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocked), v))
	})
}

// LockedIsNil applies the IsNil predicate on the "locked" field.
func LockedIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLocked)))
	})
}

// LockedNotNil applies the NotNil predicate on the "locked" field.
func LockedNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLocked)))
	})
}

// LockedByEQ applies the EQ predicate on the "locked_by" field.
func LockedByEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockedBy), v))
	})
}

// LockedByNEQ applies the NEQ predicate on the "locked_by" field.
func LockedByNEQ(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLockedBy), v))
	})
}

// LockedByIn applies the In predicate on the "locked_by" field.
func LockedByIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLockedBy), v...))
	})
}

// LockedByNotIn applies the NotIn predicate on the "locked_by" field.
func LockedByNotIn(vs ...string) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLockedBy), v...))
	})
}

// LockedByGT applies the GT predicate on the "locked_by" field.
func LockedByGT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLockedBy), v))
	})
}

// LockedByGTE applies the GTE predicate on the "locked_by" field.
func LockedByGTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLockedBy), v))
	})
}

// LockedByLT applies the LT predicate on the "locked_by" field.
func LockedByLT(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLockedBy), v))
	})
}

// LockedByLTE applies the LTE predicate on the "locked_by" field.
func LockedByLTE(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLockedBy), v))
	})
}

// LockedByContains applies the Contains predicate on the "locked_by" field.
func LockedByContains(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLockedBy), v))
	})
}

// LockedByHasPrefix applies the HasPrefix predicate on the "locked_by" field.
func LockedByHasPrefix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLockedBy), v))
	})
}

// LockedByHasSuffix applies the HasSuffix predicate on the "locked_by" field.
func LockedByHasSuffix(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLockedBy), v))
	})
}

// LockedByIsNil applies the IsNil predicate on the "locked_by" field.
func LockedByIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLockedBy)))
	})
}

// LockedByNotNil applies the NotNil predicate on the "locked_by" field.
func LockedByNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLockedBy)))
	})
}

// LockedByEqualFold applies the EqualFold predicate on the "locked_by" field.
func LockedByEqualFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLockedBy), v))
	})
}

// LockedByContainsFold applies the ContainsFold predicate on the "locked_by" field.
func LockedByContainsFold(v string) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLockedBy), v))
	})
}

// BlockedEQ applies the EQ predicate on the "blocked" field.
func BlockedEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlocked), v))
	})
}

// BlockedNEQ applies the NEQ predicate on the "blocked" field.
func BlockedNEQ(v bool) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlocked), v))
	})
}

// BlockedIsNil applies the IsNil predicate on the "blocked" field.
func BlockedIsNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBlocked)))
	})
}

// BlockedNotNil applies the NotNil predicate on the "blocked" field.
func BlockedNotNil() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBlocked)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
