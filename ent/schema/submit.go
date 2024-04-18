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
		field.Int8("status").NonNegative(),
		field.Int16("point").NonNegative(),
		field.Time("create_time").Default(time.Now()),
		field.Int("total_time").NonNegative(),
		field.Int("max_memory").NonNegative(),
		field.String("language"),
		field.Int8("case_version").Positive(),
		field.Int("problem_id"),
		field.Int("user_id"),
	}
}

// Edges of the Submit.
func (Submit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submit_cases", SubmitCase.Type),
		edge.From("problems", Problem.Type).Ref("submission").Field("problem_id").Unique().Required(),
		edge.From("users", User.Type).Ref("submission").Field("user_id").Unique().Required(),
	}
}
