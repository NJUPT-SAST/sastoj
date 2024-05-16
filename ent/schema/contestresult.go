package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ContestResult holds the schema definition for the ContestResult entity.
type ContestResult struct {
	ent.Schema
}

// Fields of the ContestResult.
func (ContestResult) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("score"),
		field.Int32("rank"),
		field.Int32("score_time"),
		field.Int32("penalty"),
		field.Int64("contest_id"),
		field.Int64("user_id"),
	}
}

// Edges of the ContestResult.
func (ContestResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contest", Contest.Type).Ref("contest_results").Field("contest_id").Unique().Required(),
		edge.From("user", User.Type).Ref("contest_results").Field("user_id").Unique().Required(),
		edge.From("submissions", Submission.Type).Ref("contest_results"),
	}
}
