// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/group"
	"sastoj/ent/loginsession"
	"sastoj/ent/predicate"
	"sastoj/ent/submission"
	"sastoj/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetSalt sets the "salt" field.
func (uu *UserUpdate) SetSalt(s string) *UserUpdate {
	uu.mutation.SetSalt(s)
	return uu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSalt(s *string) *UserUpdate {
	if s != nil {
		uu.SetSalt(*s)
	}
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(i int16) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(i)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(i *int16) *UserUpdate {
	if i != nil {
		uu.SetStatus(*i)
	}
	return uu
}

// AddStatus adds i to the "status" field.
func (uu *UserUpdate) AddStatus(i int16) *UserUpdate {
	uu.mutation.AddStatus(i)
	return uu
}

// SetGroupID sets the "group_id" field.
func (uu *UserUpdate) SetGroupID(i int64) *UserUpdate {
	uu.mutation.SetGroupID(i)
	return uu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGroupID(i *int64) *UserUpdate {
	if i != nil {
		uu.SetGroupID(*i)
	}
	return uu
}

// AddSubmissionIDs adds the "submission" edge to the Submission entity by IDs.
func (uu *UserUpdate) AddSubmissionIDs(ids ...int64) *UserUpdate {
	uu.mutation.AddSubmissionIDs(ids...)
	return uu
}

// AddSubmission adds the "submission" edges to the Submission entity.
func (uu *UserUpdate) AddSubmission(s ...*Submission) *UserUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSubmissionIDs(ids...)
}

// AddLoginSessionIDs adds the "login_sessions" edge to the LoginSession entity by IDs.
func (uu *UserUpdate) AddLoginSessionIDs(ids ...int64) *UserUpdate {
	uu.mutation.AddLoginSessionIDs(ids...)
	return uu
}

// AddLoginSessions adds the "login_sessions" edges to the LoginSession entity.
func (uu *UserUpdate) AddLoginSessions(l ...*LoginSession) *UserUpdate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.AddLoginSessionIDs(ids...)
}

// SetGroupsID sets the "groups" edge to the Group entity by ID.
func (uu *UserUpdate) SetGroupsID(id int64) *UserUpdate {
	uu.mutation.SetGroupsID(id)
	return uu
}

// SetGroups sets the "groups" edge to the Group entity.
func (uu *UserUpdate) SetGroups(g *Group) *UserUpdate {
	return uu.SetGroupsID(g.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSubmission clears all "submission" edges to the Submission entity.
func (uu *UserUpdate) ClearSubmission() *UserUpdate {
	uu.mutation.ClearSubmission()
	return uu
}

// RemoveSubmissionIDs removes the "submission" edge to Submission entities by IDs.
func (uu *UserUpdate) RemoveSubmissionIDs(ids ...int64) *UserUpdate {
	uu.mutation.RemoveSubmissionIDs(ids...)
	return uu
}

// RemoveSubmission removes "submission" edges to Submission entities.
func (uu *UserUpdate) RemoveSubmission(s ...*Submission) *UserUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSubmissionIDs(ids...)
}

// ClearLoginSessions clears all "login_sessions" edges to the LoginSession entity.
func (uu *UserUpdate) ClearLoginSessions() *UserUpdate {
	uu.mutation.ClearLoginSessions()
	return uu
}

// RemoveLoginSessionIDs removes the "login_sessions" edge to LoginSession entities by IDs.
func (uu *UserUpdate) RemoveLoginSessionIDs(ids ...int64) *UserUpdate {
	uu.mutation.RemoveLoginSessionIDs(ids...)
	return uu
}

// RemoveLoginSessions removes "login_sessions" edges to LoginSession entities.
func (uu *UserUpdate) RemoveLoginSessions(l ...*LoginSession) *UserUpdate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.RemoveLoginSessionIDs(ids...)
}

// ClearGroups clears the "groups" edge to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uu.mutation.GroupsID(); uu.mutation.GroupsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.groups"`)
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt16, value)
	}
	if uu.mutation.SubmissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSubmissionIDs(); len(nodes) > 0 && !uu.mutation.SubmissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.LoginSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedLoginSessionsIDs(); len(nodes) > 0 && !uu.mutation.LoginSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.LoginSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetSalt sets the "salt" field.
func (uuo *UserUpdateOne) SetSalt(s string) *UserUpdateOne {
	uuo.mutation.SetSalt(s)
	return uuo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSalt(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSalt(*s)
	}
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(i int16) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(i)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(i *int16) *UserUpdateOne {
	if i != nil {
		uuo.SetStatus(*i)
	}
	return uuo
}

// AddStatus adds i to the "status" field.
func (uuo *UserUpdateOne) AddStatus(i int16) *UserUpdateOne {
	uuo.mutation.AddStatus(i)
	return uuo
}

// SetGroupID sets the "group_id" field.
func (uuo *UserUpdateOne) SetGroupID(i int64) *UserUpdateOne {
	uuo.mutation.SetGroupID(i)
	return uuo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGroupID(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetGroupID(*i)
	}
	return uuo
}

// AddSubmissionIDs adds the "submission" edge to the Submission entity by IDs.
func (uuo *UserUpdateOne) AddSubmissionIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.AddSubmissionIDs(ids...)
	return uuo
}

// AddSubmission adds the "submission" edges to the Submission entity.
func (uuo *UserUpdateOne) AddSubmission(s ...*Submission) *UserUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSubmissionIDs(ids...)
}

// AddLoginSessionIDs adds the "login_sessions" edge to the LoginSession entity by IDs.
func (uuo *UserUpdateOne) AddLoginSessionIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.AddLoginSessionIDs(ids...)
	return uuo
}

// AddLoginSessions adds the "login_sessions" edges to the LoginSession entity.
func (uuo *UserUpdateOne) AddLoginSessions(l ...*LoginSession) *UserUpdateOne {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.AddLoginSessionIDs(ids...)
}

// SetGroupsID sets the "groups" edge to the Group entity by ID.
func (uuo *UserUpdateOne) SetGroupsID(id int64) *UserUpdateOne {
	uuo.mutation.SetGroupsID(id)
	return uuo
}

// SetGroups sets the "groups" edge to the Group entity.
func (uuo *UserUpdateOne) SetGroups(g *Group) *UserUpdateOne {
	return uuo.SetGroupsID(g.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSubmission clears all "submission" edges to the Submission entity.
func (uuo *UserUpdateOne) ClearSubmission() *UserUpdateOne {
	uuo.mutation.ClearSubmission()
	return uuo
}

// RemoveSubmissionIDs removes the "submission" edge to Submission entities by IDs.
func (uuo *UserUpdateOne) RemoveSubmissionIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.RemoveSubmissionIDs(ids...)
	return uuo
}

// RemoveSubmission removes "submission" edges to Submission entities.
func (uuo *UserUpdateOne) RemoveSubmission(s ...*Submission) *UserUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSubmissionIDs(ids...)
}

// ClearLoginSessions clears all "login_sessions" edges to the LoginSession entity.
func (uuo *UserUpdateOne) ClearLoginSessions() *UserUpdateOne {
	uuo.mutation.ClearLoginSessions()
	return uuo
}

// RemoveLoginSessionIDs removes the "login_sessions" edge to LoginSession entities by IDs.
func (uuo *UserUpdateOne) RemoveLoginSessionIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.RemoveLoginSessionIDs(ids...)
	return uuo
}

// RemoveLoginSessions removes "login_sessions" edges to LoginSession entities.
func (uuo *UserUpdateOne) RemoveLoginSessions(l ...*LoginSession) *UserUpdateOne {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.RemoveLoginSessionIDs(ids...)
}

// ClearGroups clears the "groups" edge to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.GroupsID(); uuo.mutation.GroupsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.groups"`)
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt16, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt16, value)
	}
	if uuo.mutation.SubmissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSubmissionIDs(); len(nodes) > 0 && !uuo.mutation.SubmissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.LoginSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedLoginSessionsIDs(); len(nodes) > 0 && !uuo.mutation.LoginSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.LoginSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
