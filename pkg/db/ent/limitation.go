// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/limitation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Limitation is the model entity for the Limitation schema.
type Limitation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Limitation holds the value of the "limitation" field.
	Limitation string `json:"limitation,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Limitation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case limitation.FieldAmount:
			values[i] = new(decimal.Decimal)
		case limitation.FieldCreatedAt, limitation.FieldUpdatedAt, limitation.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case limitation.FieldLimitation:
			values[i] = new(sql.NullString)
		case limitation.FieldID, limitation.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Limitation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Limitation fields.
func (l *Limitation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case limitation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case limitation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = uint32(value.Int64)
			}
		case limitation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = uint32(value.Int64)
			}
		case limitation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				l.DeletedAt = uint32(value.Int64)
			}
		case limitation.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				l.CoinTypeID = *value
			}
		case limitation.FieldLimitation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field limitation", values[i])
			} else if value.Valid {
				l.Limitation = value.String
			}
		case limitation.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				l.Amount = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Limitation.
// Note that you need to call Limitation.Unwrap() before calling this method if this Limitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Limitation) Update() *LimitationUpdateOne {
	return (&LimitationClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Limitation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Limitation) Unwrap() *Limitation {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Limitation is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Limitation) String() string {
	var builder strings.Builder
	builder.WriteString("Limitation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", l.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", l.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", l.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", l.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("limitation=")
	builder.WriteString(l.Limitation)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", l.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// Limitations is a parsable slice of Limitation.
type Limitations []*Limitation

func (l Limitations) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
