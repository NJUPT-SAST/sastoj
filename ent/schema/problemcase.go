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
		field.Int64("id").Unique(),
		field.Int16("point").NonNegative(),
		field.Int16("index").Positive(),
		field.Bool("is_auto").Default(false).Comment("是否自动均分"),
		field.Bool("is_deleted").Default(false),
		field.Int64("problem_id"),
	}
}

// Edges of the ProblemCase.
func (ProblemCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission_cases", SubmissionCase.Type),
		edge.From("problem", Problem.Type).Ref("problem_cases").Field("problem_id").Unique().Required(),
	}
}
