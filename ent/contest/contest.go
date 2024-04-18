// Code generated by ent, DO NOT EDIT.

package contest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the contest type in the database.
	Label = "contest"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldExtraTime holds the string denoting the extra_time field in the database.
	FieldExtraTime = "extra_time"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// EdgeProblems holds the string denoting the problems edge name in mutations.
	EdgeProblems = "problems"
	// EdgeContest holds the string denoting the contest edge name in mutations.
	EdgeContest = "contest"
	// EdgeManage holds the string denoting the manage edge name in mutations.
	EdgeManage = "manage"
	// Table holds the table name of the contest in the database.
	Table = "contests"
	// ProblemsTable is the table that holds the problems relation/edge. The primary key declared below.
	ProblemsTable = "contest_problems"
	// ProblemsInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	ProblemsInverseTable = "problems"
	// ContestTable is the table that holds the contest relation/edge. The primary key declared below.
	ContestTable = "group_contestants"
	// ContestInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	ContestInverseTable = "groups"
	// ManageTable is the table that holds the manage relation/edge. The primary key declared below.
	ManageTable = "group_admins"
	// ManageInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	ManageInverseTable = "groups"
)

// Columns holds all SQL columns for contest fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldStatus,
	FieldType,
	FieldStartTime,
	FieldEndTime,
	FieldLanguage,
	FieldExtraTime,
	FieldCreateTime,
}

var (
	// ProblemsPrimaryKey and ProblemsColumn2 are the table columns denoting the
	// primary key for the problems relation (M2M).
	ProblemsPrimaryKey = []string{"contest_id", "problem_id"}
	// ContestPrimaryKey and ContestColumn2 are the table columns denoting the
	// primary key for the contest relation (M2M).
	ContestPrimaryKey = []string{"group_id", "contest_id"}
	// ManagePrimaryKey and ManageColumn2 are the table columns denoting the
	// primary key for the manage relation (M2M).
	ManagePrimaryKey = []string{"group_id", "contest_id"}
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
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(int) error
	// ExtraTimeValidator is a validator for the "extra_time" field. It is called by the builders before save.
	ExtraTimeValidator func(int) error
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime time.Time
)

// OrderOption defines the ordering options for the Contest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByExtraTime orders the results by the extra_time field.
func ByExtraTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraTime, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
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

// ByContestCount orders the results by contest count.
func ByContestCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContestStep(), opts...)
	}
}

// ByContest orders the results by contest terms.
func ByContest(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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
func newProblemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProblemsTable, ProblemsPrimaryKey...),
	)
}
func newContestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ContestTable, ContestPrimaryKey...),
	)
}
func newManageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ManageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ManageTable, ManagePrimaryKey...),
	)
}
