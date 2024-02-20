// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProblemCase is the model entity for the ProblemCase schema.
type ProblemCase struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProblemID holds the value of the "problem_id" field.
	ProblemID int `json:"problem_id,omitempty"`
	// Point holds the value of the "point" field.
	Point int `json:"point,omitempty"`
	// Index holds the value of the "index" field.
	Index int `json:"index,omitempty"`
	// IsAuto holds the value of the "is_auto" field.
	IsAuto bool `json:"is_auto,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProblemCaseQuery when eager-loading is set.
	Edges                ProblemCaseEdges `json:"edges"`
	problem_case_problem *int
	selectValues         sql.SelectValues
}

// ProblemCaseEdges holds the relations/edges for other nodes in the graph.
type ProblemCaseEdges struct {
	// Problem holds the value of the problem edge.
	Problem *Problem `json:"problem,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProblemOrErr returns the Problem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProblemCaseEdges) ProblemOrErr() (*Problem, error) {
	if e.loadedTypes[0] {
		if e.Problem == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: problem.Label}
		}
		return e.Problem, nil
	}
	return nil, &NotLoadedError{edge: "problem"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProblemCase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case problemcase.FieldIsAuto, problemcase.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case problemcase.FieldID, problemcase.FieldProblemID, problemcase.FieldPoint, problemcase.FieldIndex:
			values[i] = new(sql.NullInt64)
		case problemcase.ForeignKeys[0]: // problem_case_problem
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProblemCase fields.
func (pc *ProblemCase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case problemcase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		case problemcase.FieldProblemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field problem_id", values[i])
			} else if value.Valid {
				pc.ProblemID = int(value.Int64)
			}
		case problemcase.FieldPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field point", values[i])
			} else if value.Valid {
				pc.Point = int(value.Int64)
			}
		case problemcase.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				pc.Index = int(value.Int64)
			}
		case problemcase.FieldIsAuto:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_auto", values[i])
			} else if value.Valid {
				pc.IsAuto = value.Bool
			}
		case problemcase.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				pc.IsDeleted = value.Bool
			}
		case problemcase.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field problem_case_problem", value)
			} else if value.Valid {
				pc.problem_case_problem = new(int)
				*pc.problem_case_problem = int(value.Int64)
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProblemCase.
// This includes values selected through modifiers, order, etc.
func (pc *ProblemCase) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryProblem queries the "problem" edge of the ProblemCase entity.
func (pc *ProblemCase) QueryProblem() *ProblemQuery {
	return NewProblemCaseClient(pc.config).QueryProblem(pc)
}

// Update returns a builder for updating this ProblemCase.
// Note that you need to call ProblemCase.Unwrap() before calling this method if this ProblemCase
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *ProblemCase) Update() *ProblemCaseUpdateOne {
	return NewProblemCaseClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the ProblemCase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *ProblemCase) Unwrap() *ProblemCase {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProblemCase is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *ProblemCase) String() string {
	var builder strings.Builder
	builder.WriteString("ProblemCase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("problem_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.ProblemID))
	builder.WriteString(", ")
	builder.WriteString("point=")
	builder.WriteString(fmt.Sprintf("%v", pc.Point))
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", pc.Index))
	builder.WriteString(", ")
	builder.WriteString("is_auto=")
	builder.WriteString(fmt.Sprintf("%v", pc.IsAuto))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", pc.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// ProblemCases is a parsable slice of ProblemCase.
type ProblemCases []*ProblemCase
