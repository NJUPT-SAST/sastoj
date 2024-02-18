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
	// EdgeContest holds the string denoting the contest edge name in mutations.
	EdgeContest = "contest"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the contestgroup in the database.
	Table = "contest_group"
	// ContestTable is the table that holds the contest relation/edge.
	ContestTable = "contest_group"
	// ContestInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestInverseTable = "contests"
	// ContestColumn is the table column denoting the contest relation/edge.
	ContestColumn = "contest_group_contest"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "contest_group"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "contest_group_group"
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
	"contest_group_contest",
	"contest_group_group",
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

// ByContestField orders the results by contest field.
func ByContestField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestStep(), sql.OrderByField(field, opts...))
	}
}

// ByGroupField orders the results by group field.
func ByGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), sql.OrderByField(field, opts...))
	}
}
func newContestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ContestTable, ContestColumn),
	)
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
	)
}
