// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sastoj/ent/group"
	"sastoj/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"salt,omitempty"`
	// Status holds the value of the "status" field.
	Status int16 `json:"status,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"group_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Submission holds the value of the submission edge.
	Submission []*Submit `json:"submission,omitempty"`
	// LoginSessions holds the value of the login_sessions edge.
	LoginSessions []*LoginSession `json:"login_sessions,omitempty"`
	// Groups holds the value of the groups edge.
	Groups *Group `json:"groups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SubmissionOrErr returns the Submission value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SubmissionOrErr() ([]*Submit, error) {
	if e.loadedTypes[0] {
		return e.Submission, nil
	}
	return nil, &NotLoadedError{edge: "submission"}
}

// LoginSessionsOrErr returns the LoginSessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LoginSessionsOrErr() ([]*LoginSession, error) {
	if e.loadedTypes[1] {
		return e.LoginSessions, nil
	}
	return nil, &NotLoadedError{edge: "login_sessions"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GroupsOrErr() (*Group, error) {
	if e.loadedTypes[2] {
		if e.Groups == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldStatus, user.FieldGroupID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldPassword, user.FieldSalt:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = int16(value.Int64)
			}
		case user.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				u.GroupID = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySubmission queries the "submission" edge of the User entity.
func (u *User) QuerySubmission() *SubmitQuery {
	return NewUserClient(u.config).QuerySubmission(u)
}

// QueryLoginSessions queries the "login_sessions" edge of the User entity.
func (u *User) QueryLoginSessions() *LoginSessionQuery {
	return NewUserClient(u.config).QueryLoginSessions(u)
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return NewUserClient(u.config).QueryGroups(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", u.GroupID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
