package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
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
		field.Text("code"),
		field.Int("status").NonNegative(),
		field.Int("point").NonNegative(),
		field.Time("create_time").Default(time.Now()),
		field.Int("total_time").NonNegative(),
		field.Int("max_memory").NonNegative(),
		field.String("language"),
		field.Int("case_version").Positive(),
	}
}

// Edges of the Submit.
func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submit_cases", SubmitCase.Type),
		edge.From("problems", Problem.Type).Ref("submission").Unique(),
		edge.From("users", User.Type).Ref("submission").Unique(),
	}
}
