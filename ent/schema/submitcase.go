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
		field.Int("id").Unique(),
		field.Int("submit_id").Positive(),
		field.Int("case_id").Positive(),
		field.Int("state").Positive(),
		field.Int("point"),
		field.Text("message"),
		field.Int("time"),
		field.Int("memory"),
	}
}

// Edges of the SubmitCase.
func (SubmitCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission", Submit.Type).Ref("submit_cases").Unique(),
		edge.From("problem_cases", ProblemCase.Type).Ref("submit_cases").Unique(),
	}
}
