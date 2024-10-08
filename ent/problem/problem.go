// Code generated by ent, DO NOT EDIT.

package problem

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the problem type in the database.
	Label = "problem"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProblemTypeID holds the string denoting the problem_type_id field in the database.
	FieldProblemTypeID = "problem_type_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldCaseVersion holds the string denoting the case_version field in the database.
	FieldCaseVersion = "case_version"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldCompareType holds the string denoting the compare_type field in the database.
	FieldCompareType = "compare_type"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldContestID holds the string denoting the contest_id field in the database.
	FieldContestID = "contest_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeSubmission holds the string denoting the submission edge name in mutations.
	EdgeSubmission = "submission"
	// EdgeContest holds the string denoting the contest edge name in mutations.
	EdgeContest = "contest"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeProblemType holds the string denoting the problem_type edge name in mutations.
	EdgeProblemType = "problem_type"
	// EdgeAdjudicators holds the string denoting the adjudicators edge name in mutations.
	EdgeAdjudicators = "adjudicators"
	// Table holds the table name of the problem in the database.
	Table = "problems"
	// SubmissionTable is the table that holds the submission relation/edge.
	SubmissionTable = "submissions"
	// SubmissionInverseTable is the table name for the Submission entity.
	// It exists in this package in order to avoid circular dependency with the "submission" package.
	SubmissionInverseTable = "submissions"
	// SubmissionColumn is the table column denoting the submission relation/edge.
	SubmissionColumn = "problem_id"
	// ContestTable is the table that holds the contest relation/edge.
	ContestTable = "problems"
	// ContestInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestInverseTable = "contests"
	// ContestColumn is the table column denoting the contest relation/edge.
	ContestColumn = "contest_id"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "problems"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
	// ProblemTypeTable is the table that holds the problem_type relation/edge.
	ProblemTypeTable = "problems"
	// ProblemTypeInverseTable is the table name for the ProblemType entity.
	// It exists in this package in order to avoid circular dependency with the "problemtype" package.
	ProblemTypeInverseTable = "problem_types"
	// ProblemTypeColumn is the table column denoting the problem_type relation/edge.
	ProblemTypeColumn = "problem_type_id"
	// AdjudicatorsTable is the table that holds the adjudicators relation/edge. The primary key declared below.
	AdjudicatorsTable = "problem_adjudicators"
	// AdjudicatorsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	AdjudicatorsInverseTable = "groups"
)

// Columns holds all SQL columns for problem fields.
var Columns = []string{
	FieldID,
	FieldProblemTypeID,
	FieldTitle,
	FieldContent,
	FieldScore,
	FieldCaseVersion,
	FieldIndex,
	FieldCompareType,
	FieldIsDeleted,
	FieldContestID,
	FieldUserID,
	FieldVisibility,
	FieldMetadata,
}

var (
	// AdjudicatorsPrimaryKey and AdjudicatorsColumn2 are the table columns denoting the
	// primary key for the adjudicators relation (M2M).
	AdjudicatorsPrimaryKey = []string{"problem_id", "group_id"}
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
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int16) error
	// DefaultCaseVersion holds the default value on creation for the "case_version" field.
	DefaultCaseVersion int16
	// IndexValidator is a validator for the "index" field. It is called by the builders before save.
	IndexValidator func(int16) error
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
	// DefaultMetadata holds the default value on creation for the "metadata" field.
	DefaultMetadata map[string]string
)

// CompareType defines the type for the "compare_type" enum field.
type CompareType string

// CompareTypeSTRICT is the default value of the CompareType enum.
const DefaultCompareType = CompareTypeSTRICT

// CompareType values.
const (
	CompareTypeSTRICT                                   CompareType = "STRICT"
	CompareTypeIGNORE_LINE_END_SPACE_AND_TEXT_END_ENTER CompareType = "IGNORE_LINE_END_SPACE_AND_TEXT_END_ENTER"
)

func (ct CompareType) String() string {
	return string(ct)
}

// CompareTypeValidator is a validator for the "compare_type" field enum values. It is called by the builders before save.
func CompareTypeValidator(ct CompareType) error {
	switch ct {
	case CompareTypeSTRICT, CompareTypeIGNORE_LINE_END_SPACE_AND_TEXT_END_ENTER:
		return nil
	default:
		return fmt.Errorf("problem: invalid enum value for compare_type field: %q", ct)
	}
}

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityPRIVATE is the default value of the Visibility enum.
const DefaultVisibility = VisibilityPRIVATE

// Visibility values.
const (
	VisibilityPRIVATE Visibility = "PRIVATE"
	VisibilityPUBLIC  Visibility = "PUBLIC"
	VisibilityCONTEST Visibility = "CONTEST"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityPRIVATE, VisibilityPUBLIC, VisibilityCONTEST:
		return nil
	default:
		return fmt.Errorf("problem: invalid enum value for visibility field: %q", v)
	}
}

// OrderOption defines the ordering options for the Problem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProblemTypeID orders the results by the problem_type_id field.
func ByProblemTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProblemTypeID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByCaseVersion orders the results by the case_version field.
func ByCaseVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCaseVersion, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}

// ByCompareType orders the results by the compare_type field.
func ByCompareType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompareType, opts...).ToFunc()
}

// ByIsDeleted orders the results by the is_deleted field.
func ByIsDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDeleted, opts...).ToFunc()
}

// ByContestID orders the results by the contest_id field.
func ByContestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContestID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
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

// ByContestField orders the results by contest field.
func ByContestField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestStep(), sql.OrderByField(field, opts...))
	}
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByProblemTypeField orders the results by problem_type field.
func ByProblemTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByAdjudicatorsCount orders the results by adjudicators count.
func ByAdjudicatorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdjudicatorsStep(), opts...)
	}
}

// ByAdjudicators orders the results by adjudicators terms.
func ByAdjudicators(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdjudicatorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubmissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmissionTable, SubmissionColumn),
	)
}
func newContestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ContestTable, ContestColumn),
	)
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newProblemTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProblemTypeTable, ProblemTypeColumn),
	)
}
func newAdjudicatorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdjudicatorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AdjudicatorsTable, AdjudicatorsPrimaryKey...),
	)
}
