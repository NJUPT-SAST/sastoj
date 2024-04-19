// Code generated by ent, DO NOT EDIT.

package problem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the problem type in the database.
	Label = "problem"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldPoint holds the string denoting the point field in the database.
	FieldPoint = "point"
	// FieldCaseVersion holds the string denoting the case_version field in the database.
	FieldCaseVersion = "case_version"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldContestID holds the string denoting the contest_id field in the database.
	FieldContestID = "contest_id"
	// EdgeProblemCases holds the string denoting the problem_cases edge name in mutations.
	EdgeProblemCases = "problem_cases"
	// EdgeSubmission holds the string denoting the submission edge name in mutations.
	EdgeSubmission = "submission"
	// EdgeContests holds the string denoting the contests edge name in mutations.
	EdgeContests = "contests"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// Table holds the table name of the problem in the database.
	Table = "problems"
	// ProblemCasesTable is the table that holds the problem_cases relation/edge.
	ProblemCasesTable = "problem_cases"
	// ProblemCasesInverseTable is the table name for the ProblemCase entity.
	// It exists in this package in order to avoid circular dependency with the "problemcase" package.
	ProblemCasesInverseTable = "problem_cases"
	// ProblemCasesColumn is the table column denoting the problem_cases relation/edge.
	ProblemCasesColumn = "problem_id"
	// SubmissionTable is the table that holds the submission relation/edge.
	SubmissionTable = "submit"
	// SubmissionInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmissionInverseTable = "submit"
	// SubmissionColumn is the table column denoting the submission relation/edge.
	SubmissionColumn = "problem_id"
	// ContestsTable is the table that holds the contests relation/edge.
	ContestsTable = "problems"
	// ContestsInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestsInverseTable = "contests"
	// ContestsColumn is the table column denoting the contests relation/edge.
	ContestsColumn = "contest_id"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_problems"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
)

// Columns holds all SQL columns for problem fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldPoint,
	FieldCaseVersion,
	FieldIndex,
	FieldIsDeleted,
	FieldConfig,
	FieldContestID,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "problem_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PointValidator is a validator for the "point" field. It is called by the builders before save.
	PointValidator func(int16) error
	// DefaultCaseVersion holds the default value on creation for the "case_version" field.
	DefaultCaseVersion int16
	// IndexValidator is a validator for the "index" field. It is called by the builders before save.
	IndexValidator func(int16) error
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
)

// OrderOption defines the ordering options for the Problem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByPoint orders the results by the point field.
func ByPoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoint, opts...).ToFunc()
}

// ByCaseVersion orders the results by the case_version field.
func ByCaseVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCaseVersion, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}

// ByIsDeleted orders the results by the is_deleted field.
func ByIsDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDeleted, opts...).ToFunc()
}

// ByConfig orders the results by the config field.
func ByConfig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfig, opts...).ToFunc()
}

// ByContestID orders the results by the contest_id field.
func ByContestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContestID, opts...).ToFunc()
}

// ByProblemCasesCount orders the results by problem_cases count.
func ByProblemCasesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProblemCasesStep(), opts...)
	}
}

// ByProblemCases orders the results by problem_cases terms.
func ByProblemCases(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemCasesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubmissionCount orders the results by submission count.
func BySubmissionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmissionStep(), opts...)
	}
}

// BySubmission orders the results by submission terms.
func BySubmission(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmissionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContestsField orders the results by contests field.
func ByContestsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestsStep(), sql.OrderByField(field, opts...))
	}
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProblemCasesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemCasesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProblemCasesTable, ProblemCasesColumn),
	)
}
func newSubmissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmissionTable, SubmissionColumn),
	)
}
func newContestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ContestsTable, ContestsColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
	)
}
