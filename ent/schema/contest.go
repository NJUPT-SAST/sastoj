package schema

import (
	"time"

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
		field.Int64("id").Unique(),
		field.String("title"),
		field.String("description"),
		field.Enum("state").Values("NORMAL", "CANCELLED", "HIDDEN", "DELETED").Default("HIDDEN"),
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
		edge.To("contestants", Group.Type),
		edge.To("managers", Group.Type),
		edge.To("contest_results", ContestResult.Type),
	}
}
