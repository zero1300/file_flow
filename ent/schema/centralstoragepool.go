package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"file_flow/ent/mixin"
	"time"
)

// CentralStoragePool holds the schema definition for the CentralStoragePool entity.
type CentralStoragePool struct {
	ent.Schema
}

// Fields of the CentralStoragePool.
func (CentralStoragePool) Fields() []ent.Field {
	return []ent.Field{
		field.Text("filename").NotEmpty().Comment("文件名"),
		field.Text("ext").NotEmpty().Comment("文件扩展名"),
		field.Float("size").Comment("文件大小"),
		field.Text("path").NotEmpty().Comment("文件路径"),
		field.Text("hash").NotEmpty().Comment("文件哈希"),
		field.Time("create_at").Default(time.Now()).Immutable(),
	}
}

func (CentralStoragePool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "central_storage_pool"},
	}
}

// Edges of the CentralStoragePool.
func (CentralStoragePool) Edges() []ent.Edge {
	return nil
}

func (CentralStoragePool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
