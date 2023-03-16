package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"file_flow/ent/mixin"
	"time"
)

// UserStoragePool holds the schema definition for the UserStoragePool entity.
type UserStoragePool struct {
	ent.Schema
}

// Fields of the UserStoragePool.
func (UserStoragePool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("uid").Comment("用户id"),
		field.Int("repo_id").Comment("中心存储id"),
		field.Int("parent_id").Default(0).Comment("父目录id, 0表示根目录"),
		field.Text("filename").NotEmpty().Comment("文件名"),
		field.Text("ext").NotEmpty().Comment("文件扩展名"),
		field.Time("create_at").Default(time.Now()).Immutable(),
	}
}

func (UserStoragePool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_storage_pool"},
	}
}

// Edges of the UserStoragePool.
func (UserStoragePool) Edges() []ent.Edge {
	return nil
}

func (UserStoragePool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
