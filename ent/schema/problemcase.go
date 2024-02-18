package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProblemCase holds the schema definition for the ProblemCase entity.
type ProblemCase struct {
	ent.Schema
}

// Fields of the ProblemCase.
func (ProblemCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("problem_id").Positive(),
		field.Int("point"),
		field.Int("index").Positive(),
		field.Bool("is_auto").Default(false),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the ProblemCase.
func (ProblemCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("problems", Problem.Type).Ref("problem_cases").Unique(),
		edge.To("submit_cases", SubmitCase.Type),
	}
}
