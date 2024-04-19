package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubmitCase holds the schema definition for the SubmitCase entity.
type SubmitCase struct {
	ent.Schema
}

// Fields of the SubmitCase.
func (SubmitCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int16("state").Positive(),
		field.Int16("point").NonNegative(),
		field.Text("message"),
		field.Int32("time").NonNegative(),
		field.Int32("memory").NonNegative(),
		field.Int64("submit_id"),
		field.Int64("problem_case_id"),
	}
}

// Edges of the SubmitCase.
func (SubmitCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission", Submit.Type).Ref("submit_cases").Field("submit_id").Unique().Required(),
		edge.From("problem_cases", ProblemCase.Type).Ref("submit_cases").Field("problem_case_id").Unique().Required(),
	}
}
