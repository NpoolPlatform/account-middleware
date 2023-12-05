package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

func (Payment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("collecting_tid", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("available_at").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return nil
}
