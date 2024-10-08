// Code generated by ent, DO NOT EDIT.

package submissionsubtask

import (
	"sastoj/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLTE(FieldID, id))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldState, v))
}

// Point applies equality check predicate on the "point" field. It's identical to PointEQ.
func Point(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldPoint, v))
}

// TotalTime applies equality check predicate on the "total_time" field. It's identical to TotalTimeEQ.
func TotalTime(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldTotalTime, v))
}

// MaxMemory applies equality check predicate on the "max_memory" field. It's identical to MaxMemoryEQ.
func MaxMemory(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldMaxMemory, v))
}

// SubmissionID applies equality check predicate on the "submission_id" field. It's identical to SubmissionIDEQ.
func SubmissionID(v int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldSubmissionID, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLTE(FieldState, v))
}

// PointEQ applies the EQ predicate on the "point" field.
func PointEQ(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldPoint, v))
}

// PointNEQ applies the NEQ predicate on the "point" field.
func PointNEQ(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNEQ(FieldPoint, v))
}

// PointIn applies the In predicate on the "point" field.
func PointIn(vs ...int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldIn(FieldPoint, vs...))
}

// PointNotIn applies the NotIn predicate on the "point" field.
func PointNotIn(vs ...int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNotIn(FieldPoint, vs...))
}

// PointGT applies the GT predicate on the "point" field.
func PointGT(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGT(FieldPoint, v))
}

// PointGTE applies the GTE predicate on the "point" field.
func PointGTE(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGTE(FieldPoint, v))
}

// PointLT applies the LT predicate on the "point" field.
func PointLT(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLT(FieldPoint, v))
}

// PointLTE applies the LTE predicate on the "point" field.
func PointLTE(v int16) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLTE(FieldPoint, v))
}

// TotalTimeEQ applies the EQ predicate on the "total_time" field.
func TotalTimeEQ(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldTotalTime, v))
}

// TotalTimeNEQ applies the NEQ predicate on the "total_time" field.
func TotalTimeNEQ(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNEQ(FieldTotalTime, v))
}

// TotalTimeIn applies the In predicate on the "total_time" field.
func TotalTimeIn(vs ...uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldIn(FieldTotalTime, vs...))
}

// TotalTimeNotIn applies the NotIn predicate on the "total_time" field.
func TotalTimeNotIn(vs ...uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNotIn(FieldTotalTime, vs...))
}

// TotalTimeGT applies the GT predicate on the "total_time" field.
func TotalTimeGT(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGT(FieldTotalTime, v))
}

// TotalTimeGTE applies the GTE predicate on the "total_time" field.
func TotalTimeGTE(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGTE(FieldTotalTime, v))
}

// TotalTimeLT applies the LT predicate on the "total_time" field.
func TotalTimeLT(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLT(FieldTotalTime, v))
}

// TotalTimeLTE applies the LTE predicate on the "total_time" field.
func TotalTimeLTE(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLTE(FieldTotalTime, v))
}

// MaxMemoryEQ applies the EQ predicate on the "max_memory" field.
func MaxMemoryEQ(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldMaxMemory, v))
}

// MaxMemoryNEQ applies the NEQ predicate on the "max_memory" field.
func MaxMemoryNEQ(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNEQ(FieldMaxMemory, v))
}

// MaxMemoryIn applies the In predicate on the "max_memory" field.
func MaxMemoryIn(vs ...uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldIn(FieldMaxMemory, vs...))
}

// MaxMemoryNotIn applies the NotIn predicate on the "max_memory" field.
func MaxMemoryNotIn(vs ...uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNotIn(FieldMaxMemory, vs...))
}

// MaxMemoryGT applies the GT predicate on the "max_memory" field.
func MaxMemoryGT(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGT(FieldMaxMemory, v))
}

// MaxMemoryGTE applies the GTE predicate on the "max_memory" field.
func MaxMemoryGTE(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldGTE(FieldMaxMemory, v))
}

// MaxMemoryLT applies the LT predicate on the "max_memory" field.
func MaxMemoryLT(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLT(FieldMaxMemory, v))
}

// MaxMemoryLTE applies the LTE predicate on the "max_memory" field.
func MaxMemoryLTE(v uint64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldLTE(FieldMaxMemory, v))
}

// SubmissionIDEQ applies the EQ predicate on the "submission_id" field.
func SubmissionIDEQ(v int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldEQ(FieldSubmissionID, v))
}

// SubmissionIDNEQ applies the NEQ predicate on the "submission_id" field.
func SubmissionIDNEQ(v int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNEQ(FieldSubmissionID, v))
}

// SubmissionIDIn applies the In predicate on the "submission_id" field.
func SubmissionIDIn(vs ...int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldIn(FieldSubmissionID, vs...))
}

// SubmissionIDNotIn applies the NotIn predicate on the "submission_id" field.
func SubmissionIDNotIn(vs ...int64) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.FieldNotIn(FieldSubmissionID, vs...))
}

// HasSubmissionCases applies the HasEdge predicate on the "submission_cases" edge.
func HasSubmissionCases() predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubmissionCasesTable, SubmissionCasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionCasesWith applies the HasEdge predicate on the "submission_cases" edge with a given conditions (other predicates).
func HasSubmissionCasesWith(preds ...predicate.SubmissionCase) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(func(s *sql.Selector) {
		step := newSubmissionCasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubmissions applies the HasEdge predicate on the "submissions" edge.
func HasSubmissions() predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubmissionsTable, SubmissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionsWith applies the HasEdge predicate on the "submissions" edge with a given conditions (other predicates).
func HasSubmissionsWith(preds ...predicate.Submission) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(func(s *sql.Selector) {
		step := newSubmissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SubmissionSubtask) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SubmissionSubtask) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SubmissionSubtask) predicate.SubmissionSubtask {
	return predicate.SubmissionSubtask(sql.NotPredicates(p))
}
