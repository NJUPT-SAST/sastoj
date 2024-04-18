// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sastoj/ent/group"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GroupName holds the value of the "group_name" field.
	GroupName string `json:"group_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges           GroupEdges `json:"edges"`
	group_subgroups *int
	selectValues    sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Admins holds the value of the admins edge.
	Admins []*Contest `json:"admins,omitempty"`
	// Contestants holds the value of the contestants edge.
	Contestants []*Contest `json:"contestants,omitempty"`
	// Problems holds the value of the problems edge.
	Problems []*Problem `json:"problems,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// RootGroup holds the value of the root_group edge.
	RootGroup *Group `json:"root_group,omitempty"`
	// Subgroups holds the value of the subgroups edge.
	Subgroups []*Group `json:"subgroups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) AdminsOrErr() ([]*Contest, error) {
	if e.loadedTypes[0] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// ContestantsOrErr returns the Contestants value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) ContestantsOrErr() ([]*Contest, error) {
	if e.loadedTypes[1] {
		return e.Contestants, nil
	}
	return nil, &NotLoadedError{edge: "contestants"}
}

// ProblemsOrErr returns the Problems value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) ProblemsOrErr() ([]*Problem, error) {
	if e.loadedTypes[2] {
		return e.Problems, nil
	}
	return nil, &NotLoadedError{edge: "problems"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// RootGroupOrErr returns the RootGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) RootGroupOrErr() (*Group, error) {
	if e.loadedTypes[4] {
		if e.RootGroup == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.RootGroup, nil
	}
	return nil, &NotLoadedError{edge: "root_group"}
}

// SubgroupsOrErr returns the Subgroups value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) SubgroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[5] {
		return e.Subgroups, nil
	}
	return nil, &NotLoadedError{edge: "subgroups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			values[i] = new(sql.NullInt64)
		case group.FieldGroupName:
			values[i] = new(sql.NullString)
		case group.ForeignKeys[0]: // group_subgroups
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case group.FieldGroupName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_name", values[i])
			} else if value.Valid {
				gr.GroupName = value.String
			}
		case group.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_subgroups", value)
			} else if value.Valid {
				gr.group_subgroups = new(int)
				*gr.group_subgroups = int(value.Int64)
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QueryAdmins queries the "admins" edge of the Group entity.
func (gr *Group) QueryAdmins() *ContestQuery {
	return NewGroupClient(gr.config).QueryAdmins(gr)
}

// QueryContestants queries the "contestants" edge of the Group entity.
func (gr *Group) QueryContestants() *ContestQuery {
	return NewGroupClient(gr.config).QueryContestants(gr)
}

// QueryProblems queries the "problems" edge of the Group entity.
func (gr *Group) QueryProblems() *ProblemQuery {
	return NewGroupClient(gr.config).QueryProblems(gr)
}

// QueryUsers queries the "users" edge of the Group entity.
func (gr *Group) QueryUsers() *UserQuery {
	return NewGroupClient(gr.config).QueryUsers(gr)
}

// QueryRootGroup queries the "root_group" edge of the Group entity.
func (gr *Group) QueryRootGroup() *GroupQuery {
	return NewGroupClient(gr.config).QueryRootGroup(gr)
}

// QuerySubgroups queries the "subgroups" edge of the Group entity.
func (gr *Group) QuerySubgroups() *GroupQuery {
	return NewGroupClient(gr.config).QuerySubgroups(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("group_name=")
	builder.WriteString(gr.GroupName)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group
