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

// LfCompare holds the schema definition for the LfCompare entity.
type LfCompare struct {
	IgnoreLineEndSpace      bool
	IgnoreLineEndEnter      bool
	IgnoreTextTrailingSpace bool
	IgnoreTextTrailingEnter bool
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("problem_type_id"),
		field.String("title"),
		field.String("content"),
		field.Int16("score").NonNegative(),
		field.Int16("case_version").Default(1),
		field.Int16("index").Positive(),
		field.JSON("lf_compare", LfCompare{}).Default(LfCompare{
			IgnoreLineEndSpace:      false,
			IgnoreLineEndEnter:      false,
			IgnoreTextTrailingSpace: false,
			IgnoreTextTrailingEnter: false,
		}),
		field.Bool("is_deleted").Default(false),
		field.Int64("contest_id"),
		field.Int64("user_id"),
		field.Enum("visibility").Values("PRIVATE", "PUBLIC", "CONTEST").Default("PRIVATE"),
		field.JSON("metadata", map[string]string{}).Default(make(map[string]string)).Comment("Metadata like 'allowed languages' of the problem."),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission", Submission.Type),
		edge.From("contest", Contest.Type).Ref("problems").Field("contest_id").Unique().Required(),
		edge.From("owner", User.Type).Ref("owned_problems").Field("user_id").Unique().Required(),
		edge.From("problem_type", ProblemType.Type).Ref("problems").Field("problem_type_id").Unique().Required(),
		edge.To("judgers", Group.Type),
	}
}
