package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/mixin"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("address").
			Optional().
			Default(""),
		field.
			String("used_for").
			Optional().
			Default(npool.AccountUsedFor_DefaultAccountUsedFor.String()),
		field.
			Bool("platform_hold_private_key").
			Optional().
			Default(false),
		field.
			Bool("active").
			Optional().
			Default(true),
		field.
			Bool("locked").
			Optional().
			Default(false),
		field.
			String("locked_by").
			Optional().
			Default(npool.LockedBy_DefaultLockedBy.String()),
		field.
			Bool("blocked").
			Optional().
			Default(false),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
