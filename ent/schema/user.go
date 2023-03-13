package schema

import (
	"entgo.io/ent"
	_ "entgo.io/ent/entc/integration/ent/runtime"
	"entgo.io/ent/schema/field"
	"file_flow/ent/mixin"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("nickname").NotEmpty(),
		field.Text("avatar").NotEmpty(),
		field.Text("email").NotEmpty().Unique(),
		field.Text("password").NotEmpty().Sensitive(),
		field.Time("create_at").Default(time.Now()).Immutable(),
		//field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
