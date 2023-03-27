package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Share holds the schema definition for the Share entity.
type Share struct {
	ent.Schema
}

// Fields of the Share.
func (Share) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_file_id").Comment("用户文件id"),
		field.Int("expiration").Comment("过期时间"),
		field.Int("click_number").Comment("点击量"),
		field.Time("create_at").Default(time.Now()).Immutable(),
	}
}

// Edges of the Share.
func (Share) Edges() []ent.Edge {
	return nil
}
