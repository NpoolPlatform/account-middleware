package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/account-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/account/v1"
	"github.com/google/uuid"
)

// Contract holds the schema definition for the Contract entity.
type Contract struct {
	ent.Schema
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Contract.
func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("pledge_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("backup").
			Optional().
			Default(false),
		field.
			String("contract_type").
			Optional().
			Default(basetypes.ContractType_DefaultContractType.String()),
	}
}

// Edges of the Contract.
func (Contract) Edges() []ent.Edge {
	return nil
}
