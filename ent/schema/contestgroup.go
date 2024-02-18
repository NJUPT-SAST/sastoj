package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ContestGroup holds the schema definition for the ContestGroup entity.
type ContestGroup struct {
	ent.Schema
}

func (ContestGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contest_group"},
	}
}

// Fields of the ContestGroup.
func (ContestGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("contest_id"),
		field.Int("group_id"),
	}
}

// Edges of the ContestGroup.
func (ContestGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contest", Contest.Type).Unique(),
		edge.To("group", Group.Type).Unique(),
	}
}
