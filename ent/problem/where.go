// Code generated by ent, DO NOT EDIT.

package problem

import (
	"sastoj/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldID, id))
}

// ProblemTypeID applies equality check predicate on the "problem_type_id" field. It's identical to ProblemTypeIDEQ.
func ProblemTypeID(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldProblemTypeID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldContent, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldScore, v))
}

// CaseVersion applies equality check predicate on the "case_version" field. It's identical to CaseVersionEQ.
func CaseVersion(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldCaseVersion, v))
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIndex, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIsDeleted, v))
}

// ContestID applies equality check predicate on the "contest_id" field. It's identical to ContestIDEQ.
func ContestID(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldContestID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldUserID, v))
}

// ProblemTypeIDEQ applies the EQ predicate on the "problem_type_id" field.
func ProblemTypeIDEQ(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldProblemTypeID, v))
}

// ProblemTypeIDNEQ applies the NEQ predicate on the "problem_type_id" field.
func ProblemTypeIDNEQ(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldProblemTypeID, v))
}

// ProblemTypeIDIn applies the In predicate on the "problem_type_id" field.
func ProblemTypeIDIn(vs ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldProblemTypeID, vs...))
}

// ProblemTypeIDNotIn applies the NotIn predicate on the "problem_type_id" field.
func ProblemTypeIDNotIn(vs ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldProblemTypeID, vs...))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContainsFold(FieldContent, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldScore, v))
}

// CaseVersionEQ applies the EQ predicate on the "case_version" field.
func CaseVersionEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldCaseVersion, v))
}

// CaseVersionNEQ applies the NEQ predicate on the "case_version" field.
func CaseVersionNEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldCaseVersion, v))
}

// CaseVersionIn applies the In predicate on the "case_version" field.
func CaseVersionIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldCaseVersion, vs...))
}

// CaseVersionNotIn applies the NotIn predicate on the "case_version" field.
func CaseVersionNotIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldCaseVersion, vs...))
}

// CaseVersionGT applies the GT predicate on the "case_version" field.
func CaseVersionGT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldCaseVersion, v))
}

// CaseVersionGTE applies the GTE predicate on the "case_version" field.
func CaseVersionGTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldCaseVersion, v))
}

// CaseVersionLT applies the LT predicate on the "case_version" field.
func CaseVersionLT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldCaseVersion, v))
}

// CaseVersionLTE applies the LTE predicate on the "case_version" field.
func CaseVersionLTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldCaseVersion, v))
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIndex, v))
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldIndex, v))
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldIndex, vs...))
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldIndex, vs...))
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldIndex, v))
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldIndex, v))
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldIndex, v))
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldIndex, v))
}

// CompareTypeEQ applies the EQ predicate on the "compare_type" field.
func CompareTypeEQ(v CompareType) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldCompareType, v))
}

// CompareTypeNEQ applies the NEQ predicate on the "compare_type" field.
func CompareTypeNEQ(v CompareType) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldCompareType, v))
}

// CompareTypeIn applies the In predicate on the "compare_type" field.
func CompareTypeIn(vs ...CompareType) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldCompareType, vs...))
}

// CompareTypeNotIn applies the NotIn predicate on the "compare_type" field.
func CompareTypeNotIn(vs ...CompareType) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldCompareType, vs...))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldIsDeleted, v))
}

// ContestIDEQ applies the EQ predicate on the "contest_id" field.
func ContestIDEQ(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldContestID, v))
}

// ContestIDNEQ applies the NEQ predicate on the "contest_id" field.
func ContestIDNEQ(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldContestID, v))
}

// ContestIDIn applies the In predicate on the "contest_id" field.
func ContestIDIn(vs ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldContestID, vs...))
}

// ContestIDNotIn applies the NotIn predicate on the "contest_id" field.
func ContestIDNotIn(vs ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldContestID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldUserID, vs...))
}

// VisibilityEQ applies the EQ predicate on the "visibility" field.
func VisibilityEQ(v Visibility) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldVisibility, v))
}

// VisibilityNEQ applies the NEQ predicate on the "visibility" field.
func VisibilityNEQ(v Visibility) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldVisibility, v))
}

// VisibilityIn applies the In predicate on the "visibility" field.
func VisibilityIn(vs ...Visibility) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldVisibility, vs...))
}

// VisibilityNotIn applies the NotIn predicate on the "visibility" field.
func VisibilityNotIn(vs ...Visibility) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldVisibility, vs...))
}

// HasSubmission applies the HasEdge predicate on the "submission" edge.
func HasSubmission() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubmissionTable, SubmissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionWith applies the HasEdge predicate on the "submission" edge with a given conditions (other predicates).
func HasSubmissionWith(preds ...predicate.Submission) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newSubmissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContest applies the HasEdge predicate on the "contest" edge.
func HasContest() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContestTable, ContestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContestWith applies the HasEdge predicate on the "contest" edge with a given conditions (other predicates).
func HasContestWith(preds ...predicate.Contest) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newContestStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProblemType applies the HasEdge predicate on the "problem_type" edge.
func HasProblemType() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProblemTypeTable, ProblemTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemTypeWith applies the HasEdge predicate on the "problem_type" edge with a given conditions (other predicates).
func HasProblemTypeWith(preds ...predicate.ProblemType) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newProblemTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAdjudicators applies the HasEdge predicate on the "adjudicators" edge.
func HasAdjudicators() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AdjudicatorsTable, AdjudicatorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdjudicatorsWith applies the HasEdge predicate on the "adjudicators" edge with a given conditions (other predicates).
func HasAdjudicatorsWith(preds ...predicate.Group) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newAdjudicatorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Problem) predicate.Problem {
	return predicate.Problem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Problem) predicate.Problem {
	return predicate.Problem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Problem) predicate.Problem {
	return predicate.Problem(sql.NotPredicates(p))
}
