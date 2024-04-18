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
		field.Int("id").Unique(),
		field.String("title"),
		field.String("content"),
		field.Int("point").NonNegative(),
		field.Int("case_version").Positive().Default(1),
		field.Int("index").Positive(),
		field.Bool("is_deleted").Default(false),
		field.String("config"),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problem_cases", ProblemCase.Type),
		edge.To("submission", Submit.Type),
		edge.From("contests", Contest.Type).Ref("problems"),
		edge.From("groups", Group.Type).Ref("problems"),
	}
}
