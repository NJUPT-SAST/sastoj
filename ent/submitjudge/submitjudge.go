// Code generated by ent, DO NOT EDIT.

package submitjudge

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the submitjudge type in the database.
	Label = "submit_judge"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSubmitID holds the string denoting the submit_id field in the database.
	FieldSubmitID = "submit_id"
	// EdgeSubmission holds the string denoting the submission edge name in mutations.
	EdgeSubmission = "submission"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the submitjudge in the database.
	Table = "submit_judge"
	// SubmissionTable is the table that holds the submission relation/edge.
	SubmissionTable = "submit_judge"
	// SubmissionInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmissionInverseTable = "submit"
	// SubmissionColumn is the table column denoting the submission relation/edge.
	SubmissionColumn = "submit_submit_judge"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "submit_judge"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_submit_judge"
)

// Columns holds all SQL columns for submitjudge fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldSubmitID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "submit_judge"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"submit_submit_judge",
	"user_submit_judge",
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
	// SubmitIDValidator is a validator for the "submit_id" field. It is called by the builders before save.
	SubmitIDValidator func(int) error
)

// OrderOption defines the ordering options for the SubmitJudge queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySubmitID orders the results by the submit_id field.
func BySubmitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmitID, opts...).ToFunc()
}

// BySubmissionField orders the results by submission field.
func BySubmissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmissionStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}
func newSubmissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubmissionTable, SubmissionColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
