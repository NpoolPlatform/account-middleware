// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/goodbenefit"
	"github.com/google/uuid"
)

// GoodBenefit is the model entity for the GoodBenefit schema.
type GoodBenefit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// Backup holds the value of the "backup" field.
	Backup bool `json:"backup,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID uuid.UUID `json:"transaction_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodBenefit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodbenefit.FieldBackup:
			values[i] = new(sql.NullBool)
		case goodbenefit.FieldCreatedAt, goodbenefit.FieldUpdatedAt, goodbenefit.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case goodbenefit.FieldID, goodbenefit.FieldGoodID, goodbenefit.FieldAccountID, goodbenefit.FieldTransactionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodBenefit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodBenefit fields.
func (gb *GoodBenefit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodbenefit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gb.ID = *value
			}
		case goodbenefit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gb.CreatedAt = uint32(value.Int64)
			}
		case goodbenefit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gb.UpdatedAt = uint32(value.Int64)
			}
		case goodbenefit.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gb.DeletedAt = uint32(value.Int64)
			}
		case goodbenefit.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gb.GoodID = *value
			}
		case goodbenefit.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				gb.AccountID = *value
			}
		case goodbenefit.FieldBackup:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field backup", values[i])
			} else if value.Valid {
				gb.Backup = value.Bool
			}
		case goodbenefit.FieldTransactionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value != nil {
				gb.TransactionID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodBenefit.
// Note that you need to call GoodBenefit.Unwrap() before calling this method if this GoodBenefit
// was returned from a transaction, and the transaction was committed or rolled back.
func (gb *GoodBenefit) Update() *GoodBenefitUpdateOne {
	return (&GoodBenefitClient{config: gb.config}).UpdateOne(gb)
}

// Unwrap unwraps the GoodBenefit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gb *GoodBenefit) Unwrap() *GoodBenefit {
	_tx, ok := gb.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodBenefit is not a transactional entity")
	}
	gb.config.driver = _tx.drv
	return gb
}

// String implements the fmt.Stringer.
func (gb *GoodBenefit) String() string {
	var builder strings.Builder
	builder.WriteString("GoodBenefit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gb.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", gb.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", gb.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", gb.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.GoodID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.AccountID))
	builder.WriteString(", ")
	builder.WriteString("backup=")
	builder.WriteString(fmt.Sprintf("%v", gb.Backup))
	builder.WriteString(", ")
	builder.WriteString("transaction_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.TransactionID))
	builder.WriteByte(')')
	return builder.String()
}

// GoodBenefits is a parsable slice of GoodBenefit.
type GoodBenefits []*GoodBenefit

func (gb GoodBenefits) config(cfg config) {
	for _i := range gb {
		gb[_i].config = cfg
	}
}
