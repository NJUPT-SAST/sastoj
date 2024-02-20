package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Submit holds the schema definition for the Submit entity.
type Submit struct {
	ent.Schema
}

func (Submit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "submit"},
	}
}

// Fields of the Submit.
func (Submit) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_id").Positive(),
		field.Int("problem_id").Positive(),
		field.Text("code"),
		field.Int("state"),
		field.Int("point"),
		field.Time("create_time"),
		field.Time("total_time"),
		field.Int("max_memory"),
		field.String("language"),
		field.Int("case_version"),
	}
}

// Edges of the Submit.
func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
		edge.To("problem", Problem.Type).Unique(),
	}
}
