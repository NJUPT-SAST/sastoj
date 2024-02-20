package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubmitJudge holds the schema definition for the SubmitJudge entity.
type SubmitJudge struct {
	ent.Schema
}

func (SubmitJudge) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "submit_judge"},
	}
}

// Fields of the SubmitJudge.
func (SubmitJudge) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_id"),
		field.Int("submit_id").Positive(),
	}
}

// Edges of the SubmitJudge.
func (SubmitJudge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission", Submit.Type).Ref("submit_judge").Unique(),
		edge.From("users", User.Type).Ref("submit_judge").Unique(),
	}
}
