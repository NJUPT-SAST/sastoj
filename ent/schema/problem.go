package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("title"),
		field.String("content"),
		field.Int16("point").NonNegative(),
		field.Int16("case_version").Default(1),
		field.Int16("index").Positive(),
		field.Bool("is_deleted").Default(false),
		field.String("config"),
		field.Int64("contest_id"),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problem_cases", ProblemCase.Type),
		edge.To("submission", Submission.Type),
		edge.From("contests", Contest.Type).Ref("problems").Field("contest_id").Unique().Required(),
		edge.To("judgers", Group.Type),
	}
}
