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

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldContent, v))
}

// Point applies equality check predicate on the "point" field. It's identical to PointEQ.
func Point(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldPoint, v))
}

// CaseVersion applies equality check predicate on the "case_version" field. It's identical to CaseVersionEQ.
func CaseVersion(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldCaseVersion, v))
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIndex, v))
}

// RestrictPresentation applies equality check predicate on the "restrict_presentation" field. It's identical to RestrictPresentationEQ.
func RestrictPresentation(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldRestrictPresentation, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIsDeleted, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldConfig, v))
}

// ContestID applies equality check predicate on the "contest_id" field. It's identical to ContestIDEQ.
func ContestID(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldContestID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldUserID, v))
}

// Visibility applies equality check predicate on the "visibility" field. It's identical to VisibilityEQ.
func Visibility(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldVisibility, v))
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

// PointEQ applies the EQ predicate on the "point" field.
func PointEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldPoint, v))
}

// PointNEQ applies the NEQ predicate on the "point" field.
func PointNEQ(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldPoint, v))
}

// PointIn applies the In predicate on the "point" field.
func PointIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldPoint, vs...))
}

// PointNotIn applies the NotIn predicate on the "point" field.
func PointNotIn(vs ...int16) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldPoint, vs...))
}

// PointGT applies the GT predicate on the "point" field.
func PointGT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldPoint, v))
}

// PointGTE applies the GTE predicate on the "point" field.
func PointGTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldPoint, v))
}

// PointLT applies the LT predicate on the "point" field.
func PointLT(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldPoint, v))
}

// PointLTE applies the LTE predicate on the "point" field.
func PointLTE(v int16) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldPoint, v))
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

// RestrictPresentationEQ applies the EQ predicate on the "restrict_presentation" field.
func RestrictPresentationEQ(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldRestrictPresentation, v))
}

// RestrictPresentationNEQ applies the NEQ predicate on the "restrict_presentation" field.
func RestrictPresentationNEQ(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldRestrictPresentation, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldIsDeleted, v))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldConfig, v))
}

// ConfigContains applies the Contains predicate on the "config" field.
func ConfigContains(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContains(FieldConfig, v))
}

// ConfigHasPrefix applies the HasPrefix predicate on the "config" field.
func ConfigHasPrefix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasPrefix(FieldConfig, v))
}

// ConfigHasSuffix applies the HasSuffix predicate on the "config" field.
func ConfigHasSuffix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasSuffix(FieldConfig, v))
}

// ConfigEqualFold applies the EqualFold predicate on the "config" field.
func ConfigEqualFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEqualFold(FieldConfig, v))
}

// ConfigContainsFold applies the ContainsFold predicate on the "config" field.
func ConfigContainsFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContainsFold(FieldConfig, v))
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
func VisibilityEQ(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldVisibility, v))
}

// VisibilityNEQ applies the NEQ predicate on the "visibility" field.
func VisibilityNEQ(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldVisibility, v))
}

// VisibilityIn applies the In predicate on the "visibility" field.
func VisibilityIn(vs ...int8) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldVisibility, vs...))
}

// VisibilityNotIn applies the NotIn predicate on the "visibility" field.
func VisibilityNotIn(vs ...int8) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldVisibility, vs...))
}

// VisibilityGT applies the GT predicate on the "visibility" field.
func VisibilityGT(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldVisibility, v))
}

// VisibilityGTE applies the GTE predicate on the "visibility" field.
func VisibilityGTE(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldVisibility, v))
}

// VisibilityLT applies the LT predicate on the "visibility" field.
func VisibilityLT(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldVisibility, v))
}

// VisibilityLTE applies the LTE predicate on the "visibility" field.
func VisibilityLTE(v int8) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldVisibility, v))
}

// HasProblemCases applies the HasEdge predicate on the "problem_cases" edge.
func HasProblemCases() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProblemCasesTable, ProblemCasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemCasesWith applies the HasEdge predicate on the "problem_cases" edge with a given conditions (other predicates).
func HasProblemCasesWith(preds ...predicate.ProblemCase) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newProblemCasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
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

// HasContests applies the HasEdge predicate on the "contests" edge.
func HasContests() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContestsTable, ContestsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContestsWith applies the HasEdge predicate on the "contests" edge with a given conditions (other predicates).
func HasContestsWith(preds ...predicate.Contest) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newContestsStep()
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

// HasJudgers applies the HasEdge predicate on the "judgers" edge.
func HasJudgers() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, JudgersTable, JudgersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJudgersWith applies the HasEdge predicate on the "judgers" edge with a given conditions (other predicates).
func HasJudgersWith(preds ...predicate.Group) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newJudgersStep()
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
