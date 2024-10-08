package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProblemType holds the schema definition for the ProblemType entity.
type ProblemType struct {
	ent.Schema
}

func (ProblemType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "problem_types"},
	}
}

// Fields of the ProblemType.
func (ProblemType) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("slug_name").Unique(),
		field.String("display_name"),
		field.String("description").Default(""),
		field.String("submission_channel_name").Default("submission"),
		field.String("self_test_channel_name").Default("self-test"),
		field.String("judge").Default("default"),
	}
}

// Edges of the ProblemType.
func (ProblemType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problems", Problem.Type),
	}
}
