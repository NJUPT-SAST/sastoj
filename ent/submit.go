// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sastoj/ent/problem"
	"sastoj/ent/submit"
	"sastoj/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Submit is the model entity for the Submit schema.
type Submit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// ProblemID holds the value of the "problem_id" field.
	ProblemID int `json:"problem_id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// State holds the value of the "state" field.
	State int `json:"state,omitempty"`
	// Point holds the value of the "point" field.
	Point int `json:"point,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// TotalTime holds the value of the "total_time" field.
	TotalTime time.Time `json:"total_time,omitempty"`
	// MaxMemory holds the value of the "max_memory" field.
	MaxMemory int `json:"max_memory,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// CaseVersion holds the value of the "case_version" field.
	CaseVersion int `json:"case_version,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubmitQuery when eager-loading is set.
	Edges              SubmitEdges `json:"edges"`
	problem_submission *int
	user_submission    *int
	selectValues       sql.SelectValues
}

// SubmitEdges holds the relations/edges for other nodes in the graph.
type SubmitEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// Problems holds the value of the problems edge.
	Problems *Problem `json:"problems,omitempty"`
	// SubmitJudge holds the value of the submit_judge edge.
	SubmitJudge []*SubmitJudge `json:"submit_judge,omitempty"`
	// SubmitCases holds the value of the submit_cases edge.
	SubmitCases []*SubmitCase `json:"submit_cases,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubmitEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ProblemsOrErr returns the Problems value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubmitEdges) ProblemsOrErr() (*Problem, error) {
	if e.loadedTypes[1] {
		if e.Problems == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: problem.Label}
		}
		return e.Problems, nil
	}
	return nil, &NotLoadedError{edge: "problems"}
}

// SubmitJudgeOrErr returns the SubmitJudge value or an error if the edge
// was not loaded in eager-loading.
func (e SubmitEdges) SubmitJudgeOrErr() ([]*SubmitJudge, error) {
	if e.loadedTypes[2] {
		return e.SubmitJudge, nil
	}
	return nil, &NotLoadedError{edge: "submit_judge"}
}

// SubmitCasesOrErr returns the SubmitCases value or an error if the edge
// was not loaded in eager-loading.
func (e SubmitEdges) SubmitCasesOrErr() ([]*SubmitCase, error) {
	if e.loadedTypes[3] {
		return e.SubmitCases, nil
	}
	return nil, &NotLoadedError{edge: "submit_cases"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Submit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case submit.FieldID, submit.FieldUserID, submit.FieldProblemID, submit.FieldState, submit.FieldPoint, submit.FieldMaxMemory, submit.FieldCaseVersion:
			values[i] = new(sql.NullInt64)
		case submit.FieldCode, submit.FieldLanguage:
			values[i] = new(sql.NullString)
		case submit.FieldCreateTime, submit.FieldTotalTime:
			values[i] = new(sql.NullTime)
		case submit.ForeignKeys[0]: // problem_submission
			values[i] = new(sql.NullInt64)
		case submit.ForeignKeys[1]: // user_submission
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Submit fields.
func (s *Submit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case submit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case submit.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = int(value.Int64)
			}
		case submit.FieldProblemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field problem_id", values[i])
			} else if value.Valid {
				s.ProblemID = int(value.Int64)
			}
		case submit.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case submit.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				s.State = int(value.Int64)
			}
		case submit.FieldPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field point", values[i])
			} else if value.Valid {
				s.Point = int(value.Int64)
			}
		case submit.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case submit.FieldTotalTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field total_time", values[i])
			} else if value.Valid {
				s.TotalTime = value.Time
			}
		case submit.FieldMaxMemory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_memory", values[i])
			} else if value.Valid {
				s.MaxMemory = int(value.Int64)
			}
		case submit.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				s.Language = value.String
			}
		case submit.FieldCaseVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field case_version", values[i])
			} else if value.Valid {
				s.CaseVersion = int(value.Int64)
			}
		case submit.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field problem_submission", value)
			} else if value.Valid {
				s.problem_submission = new(int)
				*s.problem_submission = int(value.Int64)
			}
		case submit.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_submission", value)
			} else if value.Valid {
				s.user_submission = new(int)
				*s.user_submission = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Submit.
// This includes values selected through modifiers, order, etc.
func (s *Submit) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Submit entity.
func (s *Submit) QueryUsers() *UserQuery {
	return NewSubmitClient(s.config).QueryUsers(s)
}

// QueryProblems queries the "problems" edge of the Submit entity.
func (s *Submit) QueryProblems() *ProblemQuery {
	return NewSubmitClient(s.config).QueryProblems(s)
}

// QuerySubmitJudge queries the "submit_judge" edge of the Submit entity.
func (s *Submit) QuerySubmitJudge() *SubmitJudgeQuery {
	return NewSubmitClient(s.config).QuerySubmitJudge(s)
}

// QuerySubmitCases queries the "submit_cases" edge of the Submit entity.
func (s *Submit) QuerySubmitCases() *SubmitCaseQuery {
	return NewSubmitClient(s.config).QuerySubmitCases(s)
}

// Update returns a builder for updating this Submit.
// Note that you need to call Submit.Unwrap() before calling this method if this Submit
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Submit) Update() *SubmitUpdateOne {
	return NewSubmitClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Submit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Submit) Unwrap() *Submit {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Submit is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Submit) String() string {
	var builder strings.Builder
	builder.WriteString("Submit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("problem_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ProblemID))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", s.State))
	builder.WriteString(", ")
	builder.WriteString("point=")
	builder.WriteString(fmt.Sprintf("%v", s.Point))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("total_time=")
	builder.WriteString(s.TotalTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("max_memory=")
	builder.WriteString(fmt.Sprintf("%v", s.MaxMemory))
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(s.Language)
	builder.WriteString(", ")
	builder.WriteString("case_version=")
	builder.WriteString(fmt.Sprintf("%v", s.CaseVersion))
	builder.WriteByte(')')
	return builder.String()
}

// Submits is a parsable slice of Submit.
type Submits []*Submit
