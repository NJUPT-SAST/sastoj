package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Submission holds the schema definition for the Submission entity.
type Submission struct {
	ent.Schema
}

// Fields of the Submission.
func (Submission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Text("code"),
		field.Int16("status").NonNegative(),
		field.Int16("point").NonNegative(),
		field.Time("create_time").Default(time.Now()),
		field.Int32("total_time").NonNegative(),
		field.Int32("max_memory").NonNegative(),
		field.String("language"),
		field.Int8("case_version").Positive(),
		field.Int64("problem_id"),
		field.Int64("user_id"),
	}
}

// Edges of the Submission.
func (Submission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission_cases", SubmissionCase.Type),
		edge.From("problems", Problem.Type).Ref("submission").Field("problem_id").Unique().Required(),
		edge.From("users", User.Type).Ref("submission").Field("user_id").Unique().Required(),
		edge.To("contest_results", ContestResult.Type),
	}
}
