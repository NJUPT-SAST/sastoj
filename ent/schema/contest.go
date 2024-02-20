package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contest holds the schema definition for the Contest entity.
type Contest struct {
	ent.Schema
}

// Fields of the Contest.
func (Contest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("title"),
		field.String("description"),
		field.Int("state").Positive(),
		field.Int("type").Positive(),
		field.Time("start_time"),
		field.Time("end_time"),
		field.String("language"),
		field.Int("extra_time").Positive(),
		field.Time("create_time"),
	}
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contest_group", ContestGroup.Type),
		edge.To("problems", Problem.Type),
	}
}
