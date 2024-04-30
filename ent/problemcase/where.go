// Code generated by ent, DO NOT EDIT.

package problemcase

import (
	"sastoj/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldLTE(FieldID, id))
}

// Point applies equality check predicate on the "point" field. It's identical to PointEQ.
func Point(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldPoint, v))
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldIndex, v))
}

// IsAuto applies equality check predicate on the "is_auto" field. It's identical to IsAutoEQ.
func IsAuto(v bool) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldIsAuto, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldIsDeleted, v))
}

// ProblemID applies equality check predicate on the "problem_id" field. It's identical to ProblemIDEQ.
func ProblemID(v int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldProblemID, v))
}

// PointEQ applies the EQ predicate on the "point" field.
func PointEQ(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldPoint, v))
}

// PointNEQ applies the NEQ predicate on the "point" field.
func PointNEQ(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNEQ(FieldPoint, v))
}

// PointIn applies the In predicate on the "point" field.
func PointIn(vs ...int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldIn(FieldPoint, vs...))
}

// PointNotIn applies the NotIn predicate on the "point" field.
func PointNotIn(vs ...int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNotIn(FieldPoint, vs...))
}

// PointGT applies the GT predicate on the "point" field.
func PointGT(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldGT(FieldPoint, v))
}

// PointGTE applies the GTE predicate on the "point" field.
func PointGTE(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldGTE(FieldPoint, v))
}

// PointLT applies the LT predicate on the "point" field.
func PointLT(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldLT(FieldPoint, v))
}

// PointLTE applies the LTE predicate on the "point" field.
func PointLTE(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldLTE(FieldPoint, v))
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldIndex, v))
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNEQ(FieldIndex, v))
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldIn(FieldIndex, vs...))
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNotIn(FieldIndex, vs...))
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldGT(FieldIndex, v))
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldGTE(FieldIndex, v))
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldLT(FieldIndex, v))
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int16) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldLTE(FieldIndex, v))
}

// IsAutoEQ applies the EQ predicate on the "is_auto" field.
func IsAutoEQ(v bool) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldIsAuto, v))
}

// IsAutoNEQ applies the NEQ predicate on the "is_auto" field.
func IsAutoNEQ(v bool) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNEQ(FieldIsAuto, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNEQ(FieldIsDeleted, v))
}

// ProblemIDEQ applies the EQ predicate on the "problem_id" field.
func ProblemIDEQ(v int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldEQ(FieldProblemID, v))
}

// ProblemIDNEQ applies the NEQ predicate on the "problem_id" field.
func ProblemIDNEQ(v int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNEQ(FieldProblemID, v))
}

// ProblemIDIn applies the In predicate on the "problem_id" field.
func ProblemIDIn(vs ...int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldIn(FieldProblemID, vs...))
}

// ProblemIDNotIn applies the NotIn predicate on the "problem_id" field.
func ProblemIDNotIn(vs ...int64) predicate.ProblemCase {
	return predicate.ProblemCase(sql.FieldNotIn(FieldProblemID, vs...))
}

// HasSubmissionCases applies the HasEdge predicate on the "submission_cases" edge.
func HasSubmissionCases() predicate.ProblemCase {
	return predicate.ProblemCase(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubmissionCasesTable, SubmissionCasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionCasesWith applies the HasEdge predicate on the "submission_cases" edge with a given conditions (other predicates).
func HasSubmissionCasesWith(preds ...predicate.SubmissionCase) predicate.ProblemCase {
	return predicate.ProblemCase(func(s *sql.Selector) {
		step := newSubmissionCasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProblem applies the HasEdge predicate on the "problem" edge.
func HasProblem() predicate.ProblemCase {
	return predicate.ProblemCase(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProblemTable, ProblemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemWith applies the HasEdge predicate on the "problem" edge with a given conditions (other predicates).
func HasProblemWith(preds ...predicate.Problem) predicate.ProblemCase {
	return predicate.ProblemCase(func(s *sql.Selector) {
		step := newProblemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProblemCase) predicate.ProblemCase {
	return predicate.ProblemCase(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProblemCase) predicate.ProblemCase {
	return predicate.ProblemCase(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProblemCase) predicate.ProblemCase {
	return predicate.ProblemCase(sql.NotPredicates(p))
}
