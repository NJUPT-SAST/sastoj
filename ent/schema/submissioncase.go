package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubmissionCase holds the schema definition for the SubmissionCase entity.
type SubmissionCase struct {
	ent.Schema
}

func (SubmissionCase) Annitations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "submission_case"},
	}
}

// Fields of the SubmissionCase.
func (SubmissionCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int16("state"),
		field.Int16("point"),
		field.Text("message"),
		field.Int32("time"),
		field.Int32("memory"),
		field.Int64("submission_id"),
		field.Int64("problem_case_id"),
	}
}

// Edges of the SubmissionCase.
func (SubmissionCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission", Submission.Type).Ref("submission_cases").Field("submission_id").Unique().Required(),
		edge.From("problem_cases", ProblemCase.Type).Ref("submission_cases").Field("problem_case_id").Unique().Required(),
	}
}
