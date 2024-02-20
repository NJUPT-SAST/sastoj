// Code generated by ent, DO NOT EDIT.

package contestgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the contestgroup type in the database.
	Label = "contest_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContestID holds the string denoting the contest_id field in the database.
	FieldContestID = "contest_id"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// EdgeContests holds the string denoting the contests edge name in mutations.
	EdgeContests = "contests"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// Table holds the table name of the contestgroup in the database.
	Table = "contest_group"
	// ContestsTable is the table that holds the contests relation/edge.
	ContestsTable = "contest_group"
	// ContestsInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestsInverseTable = "contests"
	// ContestsColumn is the table column denoting the contests relation/edge.
	ContestsColumn = "contest_contest_group"
	// GroupsTable is the table that holds the groups relation/edge.
	GroupsTable = "contest_group"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "group_contest_group"
)

// Columns holds all SQL columns for contestgroup fields.
var Columns = []string{
	FieldID,
	FieldContestID,
	FieldGroupID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "contest_group"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"contest_contest_group",
	"group_contest_group",
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

// OrderOption defines the ordering options for the ContestGroup queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContestID orders the results by the contest_id field.
func ByContestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContestID, opts...).ToFunc()
}

// ByGroupID orders the results by the group_id field.
func ByGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupID, opts...).ToFunc()
}

// ByContestsField orders the results by contests field.
func ByContestsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestsStep(), sql.OrderByField(field, opts...))
	}
}

// ByGroupsField orders the results by groups field.
func ByGroupsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), sql.OrderByField(field, opts...))
	}
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
		sqlgraph.Edge(sqlgraph.M2O, true, GroupsTable, GroupsColumn),
	)
}
