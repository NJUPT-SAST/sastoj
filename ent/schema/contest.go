package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
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
		field.Int16("status").NonNegative(),
		field.Int16("type").Positive(),
		field.Time("start_time"),
		field.Time("end_time"),
		field.String("language"),
		field.Int16("extra_time").Default(0).NonNegative(),
		field.Time("create_time").Default(time.Now()),
	}
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problems", Problem.Type),
		edge.From("contest", Group.Type).Ref("contestants"),
		edge.From("manage", Group.Type).Ref("admins"),
	}
}
