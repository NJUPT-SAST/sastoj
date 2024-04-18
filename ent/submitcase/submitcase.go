// Code generated by ent, DO NOT EDIT.

package submitcase

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the submitcase type in the database.
	Label = "submit_case"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldPoint holds the string denoting the point field in the database.
	FieldPoint = "point"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldSubmitID holds the string denoting the submit_id field in the database.
	FieldSubmitID = "submit_id"
	// FieldProblemCaseID holds the string denoting the problem_case_id field in the database.
	FieldProblemCaseID = "problem_case_id"
	// EdgeSubmission holds the string denoting the submission edge name in mutations.
	EdgeSubmission = "submission"
	// EdgeProblemCases holds the string denoting the problem_cases edge name in mutations.
	EdgeProblemCases = "problem_cases"
	// Table holds the table name of the submitcase in the database.
	Table = "submit_cases"
	// SubmissionTable is the table that holds the submission relation/edge.
	SubmissionTable = "submit_cases"
	// SubmissionInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmissionInverseTable = "submit"
	// SubmissionColumn is the table column denoting the submission relation/edge.
	SubmissionColumn = "submit_id"
	// ProblemCasesTable is the table that holds the problem_cases relation/edge.
	ProblemCasesTable = "submit_cases"
	// ProblemCasesInverseTable is the table name for the ProblemCase entity.
	// It exists in this package in order to avoid circular dependency with the "problemcase" package.
	ProblemCasesInverseTable = "problem_cases"
	// ProblemCasesColumn is the table column denoting the problem_cases relation/edge.
	ProblemCasesColumn = "problem_case_id"
)

// Columns holds all SQL columns for submitcase fields.
var Columns = []string{
	FieldID,
	FieldState,
	FieldPoint,
	FieldMessage,
	FieldTime,
	FieldMemory,
	FieldSubmitID,
	FieldProblemCaseID,
}

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
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(int16) error
	// PointValidator is a validator for the "point" field. It is called by the builders before save.
	PointValidator func(int16) error
	// TimeValidator is a validator for the "time" field. It is called by the builders before save.
	TimeValidator func(int) error
	// MemoryValidator is a validator for the "memory" field. It is called by the builders before save.
	MemoryValidator func(int) error
)

// OrderOption defines the ordering options for the SubmitCase queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByPoint orders the results by the point field.
func ByPoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoint, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByMemory orders the results by the memory field.
func ByMemory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemory, opts...).ToFunc()
}

// BySubmitID orders the results by the submit_id field.
func BySubmitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmitID, opts...).ToFunc()
}

// ByProblemCaseID orders the results by the problem_case_id field.
func ByProblemCaseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProblemCaseID, opts...).ToFunc()
}

// BySubmissionField orders the results by submission field.
func BySubmissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmissionStep(), sql.OrderByField(field, opts...))
	}
}

// ByProblemCasesField orders the results by problem_cases field.
func ByProblemCasesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemCasesStep(), sql.OrderByField(field, opts...))
	}
}
func newSubmissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubmissionTable, SubmissionColumn),
	)
}
func newProblemCasesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemCasesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProblemCasesTable, ProblemCasesColumn),
	)
}
