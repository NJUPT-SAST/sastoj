// Code generated by ent, DO NOT EDIT.

package group

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupName holds the string denoting the group_name field in the database.
	FieldGroupName = "group_name"
	// FieldRootGroupID holds the string denoting the root_group_id field in the database.
	FieldRootGroupID = "root_group_id"
	// EdgeManage holds the string denoting the manage edge name in mutations.
	EdgeManage = "manage"
	// EdgeContests holds the string denoting the contests edge name in mutations.
	EdgeContests = "contests"
	// EdgeProblems holds the string denoting the problems edge name in mutations.
	EdgeProblems = "problems"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeRootGroup holds the string denoting the root_group edge name in mutations.
	EdgeRootGroup = "root_group"
	// EdgeSubgroups holds the string denoting the subgroups edge name in mutations.
	EdgeSubgroups = "subgroups"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// ManageTable is the table that holds the manage relation/edge. The primary key declared below.
	ManageTable = "contest_managers"
	// ManageInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ManageInverseTable = "contests"
	// ContestsTable is the table that holds the contests relation/edge. The primary key declared below.
	ContestsTable = "contest_contestants"
	// ContestsInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestsInverseTable = "contests"
	// ProblemsTable is the table that holds the problems relation/edge. The primary key declared below.
	ProblemsTable = "problem_judgers"
	// ProblemsInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	ProblemsInverseTable = "problems"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "group_id"
	// RootGroupTable is the table that holds the root_group relation/edge.
	RootGroupTable = "groups"
	// RootGroupColumn is the table column denoting the root_group relation/edge.
	RootGroupColumn = "root_group_id"
	// SubgroupsTable is the table that holds the subgroups relation/edge.
	SubgroupsTable = "groups"
	// SubgroupsColumn is the table column denoting the subgroups relation/edge.
	SubgroupsColumn = "root_group_id"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldGroupName,
	FieldRootGroupID,
}

var (
	// ManagePrimaryKey and ManageColumn2 are the table columns denoting the
	// primary key for the manage relation (M2M).
	ManagePrimaryKey = []string{"contest_id", "group_id"}
	// ContestsPrimaryKey and ContestsColumn2 are the table columns denoting the
	// primary key for the contests relation (M2M).
	ContestsPrimaryKey = []string{"contest_id", "group_id"}
	// ProblemsPrimaryKey and ProblemsColumn2 are the table columns denoting the
	// primary key for the problems relation (M2M).
	ProblemsPrimaryKey = []string{"problem_id", "group_id"}
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
	// DefaultGroupName holds the default value on creation for the "group_name" field.
	DefaultGroupName string
	// DefaultRootGroupID holds the default value on creation for the "root_group_id" field.
	DefaultRootGroupID int64
)

// OrderOption defines the ordering options for the Group queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGroupName orders the results by the group_name field.
func ByGroupName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupName, opts...).ToFunc()
}

// ByRootGroupID orders the results by the root_group_id field.
func ByRootGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRootGroupID, opts...).ToFunc()
}

// ByManageCount orders the results by manage count.
func ByManageCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newManageStep(), opts...)
	}
}

// ByManage orders the results by manage terms.
func ByManage(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newManageStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContestsCount orders the results by contests count.
func ByContestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContestsStep(), opts...)
	}
}

// ByContests orders the results by contests terms.
func ByContests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProblemsCount orders the results by problems count.
func ByProblemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProblemsStep(), opts...)
	}
}

// ByProblems orders the results by problems terms.
func ByProblems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRootGroupField orders the results by root_group field.
func ByRootGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRootGroupStep(), sql.OrderByField(field, opts...))
	}
}

// BySubgroupsCount orders the results by subgroups count.
func BySubgroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubgroupsStep(), opts...)
	}
}

// BySubgroups orders the results by subgroups terms.
func BySubgroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubgroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newManageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ManageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ManageTable, ManagePrimaryKey...),
	)
}
func newContestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ContestsTable, ContestsPrimaryKey...),
	)
}
func newProblemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProblemsTable, ProblemsPrimaryKey...),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
	)
}
func newRootGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RootGroupTable, RootGroupColumn),
	)
}
func newSubgroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubgroupsTable, SubgroupsColumn),
	)
}
