package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubmissionSubtask holds the schema definition for the SubmissionSubtask entity.
type SubmissionSubtask struct {
	ent.Schema
}

func (SubmissionSubtask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "submission_subtasks"},
	}
}

// Fields of the SubmissionSubtask.
func (SubmissionSubtask) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int16("state"),
		field.Int16("point"),
		field.Uint64("total_time"),
		field.Uint64("max_memory"),
		field.Int64("submission_id"),
	}
}

// Edges of the SubmissionSubtask.
func (SubmissionSubtask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission_cases", SubmissionCase.Type),
		edge.From("submissions", Submission.Type).Ref("submission_subtasks").Field("submission_id").Unique().Required(),
	}
}
