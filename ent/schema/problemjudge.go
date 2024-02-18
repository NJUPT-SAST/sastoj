package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProblemJudge holds the schema definition for the ProblemJudge entity.
type ProblemJudge struct {
	ent.Schema
}

// Fields of the ProblemJudge.
func (ProblemJudge) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("group_id").Positive(),
		field.Int("problem_id").Positive(),
	}
}

// Edges of the ProblemJudge.
func (ProblemJudge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).Ref("problem_judges").Unique(),
		edge.From("problems", Problem.Type).Ref("problem_judges").Unique(),
	}
}
