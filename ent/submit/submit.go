// Code generated by ent, DO NOT EDIT.

package submit

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the submit type in the database.
	Label = "submit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldProblemID holds the string denoting the problem_id field in the database.
	FieldProblemID = "problem_id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
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
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeProblems holds the string denoting the problems edge name in mutations.
	EdgeProblems = "problems"
	// EdgeSubmitJudge holds the string denoting the submit_judge edge name in mutations.
	EdgeSubmitJudge = "submit_judge"
	// EdgeSubmitCases holds the string denoting the submit_cases edge name in mutations.
	EdgeSubmitCases = "submit_cases"
	// Table holds the table name of the submit in the database.
	Table = "submit"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "submit"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_submission"
	// ProblemsTable is the table that holds the problems relation/edge.
	ProblemsTable = "submit"
	// ProblemsInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	ProblemsInverseTable = "problems"
	// ProblemsColumn is the table column denoting the problems relation/edge.
	ProblemsColumn = "problem_submission"
	// SubmitJudgeTable is the table that holds the submit_judge relation/edge.
	SubmitJudgeTable = "submit_judge"
	// SubmitJudgeInverseTable is the table name for the SubmitJudge entity.
	// It exists in this package in order to avoid circular dependency with the "submitjudge" package.
	SubmitJudgeInverseTable = "submit_judge"
	// SubmitJudgeColumn is the table column denoting the submit_judge relation/edge.
	SubmitJudgeColumn = "submit_submit_judge"
	// SubmitCasesTable is the table that holds the submit_cases relation/edge.
	SubmitCasesTable = "submit_cases"
	// SubmitCasesInverseTable is the table name for the SubmitCase entity.
	// It exists in this package in order to avoid circular dependency with the "submitcase" package.
	SubmitCasesInverseTable = "submit_cases"
	// SubmitCasesColumn is the table column denoting the submit_cases relation/edge.
	SubmitCasesColumn = "submit_submit_cases"
)

// Columns holds all SQL columns for submit fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldProblemID,
	FieldCode,
	FieldState,
	FieldPoint,
	FieldCreateTime,
	FieldTotalTime,
	FieldMaxMemory,
	FieldLanguage,
	FieldCaseVersion,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "submit"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"problem_submission",
	"user_submission",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(int) error
	// ProblemIDValidator is a validator for the "problem_id" field. It is called by the builders before save.
	ProblemIDValidator func(int) error
)

// OrderOption defines the ordering options for the Submit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByProblemID orders the results by the problem_id field.
func ByProblemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProblemID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
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

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}

// ByProblemsField orders the results by problems field.
func ByProblemsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemsStep(), sql.OrderByField(field, opts...))
	}
}

// BySubmitJudgeCount orders the results by submit_judge count.
func BySubmitJudgeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmitJudgeStep(), opts...)
	}
}

// BySubmitJudge orders the results by submit_judge terms.
func BySubmitJudge(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmitJudgeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubmitCasesCount orders the results by submit_cases count.
func BySubmitCasesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmitCasesStep(), opts...)
	}
}

// BySubmitCases orders the results by submit_cases terms.
func BySubmitCases(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmitCasesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
func newProblemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProblemsTable, ProblemsColumn),
	)
}
func newSubmitJudgeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmitJudgeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmitJudgeTable, SubmitJudgeColumn),
	)
}
func newSubmitCasesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmitCasesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmitCasesTable, SubmitCasesColumn),
	)
}
