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

func (SubmissionCase) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "submission_cases"},
	}
}

// Fields of the SubmissionCase.
func (SubmissionCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int16("state"),
		field.Int16("point"),
		field.Uint64("time"),
		field.Uint64("memory"),
		field.String("stdout").Default(""),
		field.String("stderr").Default(""),
		field.Int64("submission_subtask_id"),
	}
}

// Edges of the SubmissionCase.
func (SubmissionCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission_subtasks", SubmissionSubtask.Type).Ref("submission_cases").Field("submission_subtask_id").Unique().Required(),
	}
}
