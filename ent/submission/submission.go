// Code generated by ent, DO NOT EDIT.

package submission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the submission type in the database.
	Label = "submission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCompileMessage holds the string denoting the compile_message field in the database.
	FieldCompileMessage = "compile_message"
	// FieldPoint holds the string denoting the point field in the database.
	FieldPoint = "point"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldTotalTime holds the string denoting the total_time field in the database.
	FieldTotalTime = "total_time"
	// FieldMaxMemory holds the string denoting the max_memory field in the database.
	FieldMaxMemory = "max_memory"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldCaseVersion holds the string denoting the case_version field in the database.
	FieldCaseVersion = "case_version"
	// FieldProblemID holds the string denoting the problem_id field in the database.
	FieldProblemID = "problem_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeSubmissionCases holds the string denoting the submission_cases edge name in mutations.
	EdgeSubmissionCases = "submission_cases"
	// EdgeProblems holds the string denoting the problems edge name in mutations.
	EdgeProblems = "problems"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeContestResults holds the string denoting the contest_results edge name in mutations.
	EdgeContestResults = "contest_results"
	// Table holds the table name of the submission in the database.
	Table = "submissions"
	// SubmissionCasesTable is the table that holds the submission_cases relation/edge.
	SubmissionCasesTable = "submission_cases"
	// SubmissionCasesInverseTable is the table name for the SubmissionCase entity.
	// It exists in this package in order to avoid circular dependency with the "submissioncase" package.
	SubmissionCasesInverseTable = "submission_cases"
	// SubmissionCasesColumn is the table column denoting the submission_cases relation/edge.
	SubmissionCasesColumn = "submission_id"
	// ProblemsTable is the table that holds the problems relation/edge.
	ProblemsTable = "submissions"
	// ProblemsInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	ProblemsInverseTable = "problems"
	// ProblemsColumn is the table column denoting the problems relation/edge.
	ProblemsColumn = "problem_id"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "submissions"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
	// ContestResultsTable is the table that holds the contest_results relation/edge. The primary key declared below.
	ContestResultsTable = "submission_contest_results"
	// ContestResultsInverseTable is the table name for the ContestResult entity.
	// It exists in this package in order to avoid circular dependency with the "contestresult" package.
	ContestResultsInverseTable = "contest_results"
)

// Columns holds all SQL columns for submission fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldStatus,
	FieldCompileMessage,
	FieldPoint,
	FieldCreateTime,
	FieldTotalTime,
	FieldMaxMemory,
	FieldLanguage,
	FieldCaseVersion,
	FieldProblemID,
	FieldUserID,
}

var (
	// ContestResultsPrimaryKey and ContestResultsColumn2 are the table columns denoting the
	// primary key for the contest_results relation (M2M).
	ContestResultsPrimaryKey = []string{"submission_id", "contest_result_id"}
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime time.Time
)

// OrderOption defines the ordering options for the Submission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCompileMessage orders the results by the compile_message field.
func ByCompileMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompileMessage, opts...).ToFunc()
}

// ByPoint orders the results by the point field.
func ByPoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoint, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByTotalTime orders the results by the total_time field.
func ByTotalTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalTime, opts...).ToFunc()
}

// ByMaxMemory orders the results by the max_memory field.
func ByMaxMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxMemory, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByCaseVersion orders the results by the case_version field.
func ByCaseVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCaseVersion, opts...).ToFunc()
}

// ByProblemID orders the results by the problem_id field.
func ByProblemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProblemID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySubmissionCasesCount orders the results by submission_cases count.
func BySubmissionCasesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmissionCasesStep(), opts...)
	}
}

// BySubmissionCases orders the results by submission_cases terms.
func BySubmissionCases(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmissionCasesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProblemsField orders the results by problems field.
func ByProblemsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemsStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}

// ByContestResultsCount orders the results by contest_results count.
func ByContestResultsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContestResultsStep(), opts...)
	}
}

// ByContestResults orders the results by contest_results terms.
func ByContestResults(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestResultsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubmissionCasesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmissionCasesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmissionCasesTable, SubmissionCasesColumn),
	)
}
func newProblemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProblemsTable, ProblemsColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
func newContestResultsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestResultsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ContestResultsTable, ContestResultsPrimaryKey...),
	)
}
