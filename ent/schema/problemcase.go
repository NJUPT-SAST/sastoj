package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	gen "sastoj/ent"
	"sastoj/ent/hook"
)

// ProblemCase holds the schema definition for the ProblemCase entity.
type ProblemCase struct {
	ent.Schema
}

// Fields of the ProblemCase.
func (ProblemCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int16("point").NonNegative(),
		field.Int16("index").Positive(),
		field.Bool("is_auto").Default(false).Comment("是否自动均分"),
		field.Bool("is_deleted").Default(false),
		field.Int64("problem_id"),
	}
}

// Edges of the ProblemCase.
func (ProblemCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission_cases", SubmissionCase.Type),
		edge.From("problem", Problem.Type).Ref("problem_cases").Field("problem_id").Unique().Required(),
	}
}
func (ProblemCase) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ProblemCaseFunc(func(ctx context.Context, m *gen.ProblemCaseMutation) (ent.Value, error) {
					isAuto, isAutoOk := m.IsAuto()
					point, pointOk := m.Point()
					if isAutoOk && pointOk && !isAuto && point == 0 {
						return nil, fmt.Errorf("refuse to set point to 0 and isAuto to false at the same time")
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
