// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/group"
	"sastoj/ent/loginsession"
	"sastoj/ent/submit"
	"sastoj/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetSalt sets the "salt" field.
func (uc *UserCreate) SetSalt(s string) *UserCreate {
	uc.mutation.SetSalt(s)
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(i int16) *UserCreate {
	uc.mutation.SetStatus(i)
	return uc
}

// SetGroupID sets the "group_id" field.
func (uc *UserCreate) SetGroupID(i int) *UserCreate {
	uc.mutation.SetGroupID(i)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddSubmissionIDs adds the "submission" edge to the Submit entity by IDs.
func (uc *UserCreate) AddSubmissionIDs(ids ...int) *UserCreate {
	uc.mutation.AddSubmissionIDs(ids...)
	return uc
}

// AddSubmission adds the "submission" edges to the Submit entity.
func (uc *UserCreate) AddSubmission(s ...*Submit) *UserCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddSubmissionIDs(ids...)
}

// AddLoginSessionIDs adds the "login_sessions" edge to the LoginSession entity by IDs.
func (uc *UserCreate) AddLoginSessionIDs(ids ...int) *UserCreate {
	uc.mutation.AddLoginSessionIDs(ids...)
	return uc
}

// AddLoginSessions adds the "login_sessions" edges to the LoginSession entity.
func (uc *UserCreate) AddLoginSessions(l ...*LoginSession) *UserCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uc.AddLoginSessionIDs(ids...)
}

// SetGroupsID sets the "groups" edge to the Group entity by ID.
func (uc *UserCreate) SetGroupsID(id int) *UserCreate {
	uc.mutation.SetGroupsID(id)
	return uc
}

// SetGroups sets the "groups" edge to the Group entity.
func (uc *UserCreate) SetGroups(g *Group) *UserCreate {
	return uc.SetGroupsID(g.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Username(); !ok {
		v := user.DefaultUsername
		uc.mutation.SetUsername(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "User.salt"`)}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "User.group_id"`)}
	}
	if _, ok := uc.mutation.GroupsID(); !ok {
		return &ValidationError{Name: "groups", err: errors.New(`ent: missing required edge "User.groups"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
		_node.Salt = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt16, value)
		_node.Status = value
	}
	if nodes := uc.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubmissionTable,
			Columns: []string{user.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.LoginSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginSessionsTable,
			Columns: []string{user.LoginSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
