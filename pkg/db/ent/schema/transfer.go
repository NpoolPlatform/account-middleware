package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// Transfer holds the schema definition for the Account entity.
type Transfer struct {
	ent.Schema
}

func (Transfer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Transfer.
func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("target_user_id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}
