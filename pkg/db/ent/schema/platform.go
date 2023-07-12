package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

// Platform holds the schema definition for the Platform entity.
type Platform struct {
	ent.Schema
}

func (Platform) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Platform.
func (Platform) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("used_for").
			Optional().
			Default(basetypes.AccountUsedFor_DefaultAccountUsedFor.String()),
		field.
			Bool("backup").
			Optional().
			Default(false),
	}
}

// Edges of the Platform.
func (Platform) Edges() []ent.Edge {
	return nil
}
