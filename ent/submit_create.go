// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/problem"
	"sastoj/ent/submit"
	"sastoj/ent/submitcase"
	"sastoj/ent/submitjudge"
	"sastoj/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmitCreate is the builder for creating a Submit entity.
type SubmitCreate struct {
	config
	mutation *SubmitMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (sc *SubmitCreate) SetUserID(i int) *SubmitCreate {
	sc.mutation.SetUserID(i)
	return sc
}

// SetProblemID sets the "problem_id" field.
func (sc *SubmitCreate) SetProblemID(i int) *SubmitCreate {
	sc.mutation.SetProblemID(i)
	return sc
}

// SetCode sets the "code" field.
func (sc *SubmitCreate) SetCode(s string) *SubmitCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetState sets the "state" field.
func (sc *SubmitCreate) SetState(i int) *SubmitCreate {
	sc.mutation.SetState(i)
	return sc
}

// SetPoint sets the "point" field.
func (sc *SubmitCreate) SetPoint(i int) *SubmitCreate {
	sc.mutation.SetPoint(i)
	return sc
}

// SetCreateTime sets the "create_time" field.
func (sc *SubmitCreate) SetCreateTime(t time.Time) *SubmitCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetTotalTime sets the "total_time" field.
func (sc *SubmitCreate) SetTotalTime(i int) *SubmitCreate {
	sc.mutation.SetTotalTime(i)
	return sc
}

// SetMaxMemory sets the "max_memory" field.
func (sc *SubmitCreate) SetMaxMemory(i int) *SubmitCreate {
	sc.mutation.SetMaxMemory(i)
	return sc
}

// SetLanguage sets the "language" field.
func (sc *SubmitCreate) SetLanguage(s string) *SubmitCreate {
	sc.mutation.SetLanguage(s)
	return sc
}

// SetCaseVersion sets the "case_version" field.
func (sc *SubmitCreate) SetCaseVersion(i int) *SubmitCreate {
	sc.mutation.SetCaseVersion(i)
	return sc
}

// SetID sets the "id" field.
func (sc *SubmitCreate) SetID(i int) *SubmitCreate {
	sc.mutation.SetID(i)
	return sc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (sc *SubmitCreate) SetUsersID(id int) *SubmitCreate {
	sc.mutation.SetUsersID(id)
	return sc
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (sc *SubmitCreate) SetNillableUsersID(id *int) *SubmitCreate {
	if id != nil {
		sc = sc.SetUsersID(*id)
	}
	return sc
}

// SetUsers sets the "users" edge to the User entity.
func (sc *SubmitCreate) SetUsers(u *User) *SubmitCreate {
	return sc.SetUsersID(u.ID)
}

// SetProblemsID sets the "problems" edge to the Problem entity by ID.
func (sc *SubmitCreate) SetProblemsID(id int) *SubmitCreate {
	sc.mutation.SetProblemsID(id)
	return sc
}

// SetNillableProblemsID sets the "problems" edge to the Problem entity by ID if the given value is not nil.
func (sc *SubmitCreate) SetNillableProblemsID(id *int) *SubmitCreate {
	if id != nil {
		sc = sc.SetProblemsID(*id)
	}
	return sc
}

// SetProblems sets the "problems" edge to the Problem entity.
func (sc *SubmitCreate) SetProblems(p *Problem) *SubmitCreate {
	return sc.SetProblemsID(p.ID)
}

// AddSubmitJudgeIDs adds the "submit_judge" edge to the SubmitJudge entity by IDs.
func (sc *SubmitCreate) AddSubmitJudgeIDs(ids ...int) *SubmitCreate {
	sc.mutation.AddSubmitJudgeIDs(ids...)
	return sc
}

// AddSubmitJudge adds the "submit_judge" edges to the SubmitJudge entity.
func (sc *SubmitCreate) AddSubmitJudge(s ...*SubmitJudge) *SubmitCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddSubmitJudgeIDs(ids...)
}

// AddSubmitCaseIDs adds the "submit_cases" edge to the SubmitCase entity by IDs.
func (sc *SubmitCreate) AddSubmitCaseIDs(ids ...int) *SubmitCreate {
	sc.mutation.AddSubmitCaseIDs(ids...)
	return sc
}

// AddSubmitCases adds the "submit_cases" edges to the SubmitCase entity.
func (sc *SubmitCreate) AddSubmitCases(s ...*SubmitCase) *SubmitCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddSubmitCaseIDs(ids...)
}

// Mutation returns the SubmitMutation object of the builder.
func (sc *SubmitCreate) Mutation() *SubmitMutation {
	return sc.mutation
}

// Save creates the Submit in the database.
func (sc *SubmitCreate) Save(ctx context.Context) (*Submit, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubmitCreate) SaveX(ctx context.Context) *Submit {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubmitCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubmitCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubmitCreate) check() error {
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Submit.user_id"`)}
	}
	if v, ok := sc.mutation.UserID(); ok {
		if err := submit.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Submit.user_id": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ProblemID(); !ok {
		return &ValidationError{Name: "problem_id", err: errors.New(`ent: missing required field "Submit.problem_id"`)}
	}
	if v, ok := sc.mutation.ProblemID(); ok {
		if err := submit.ProblemIDValidator(v); err != nil {
			return &ValidationError{Name: "problem_id", err: fmt.Errorf(`ent: validator failed for field "Submit.problem_id": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Submit.code"`)}
	}
	if _, ok := sc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Submit.state"`)}
	}
	if _, ok := sc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "Submit.point"`)}
	}
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Submit.create_time"`)}
	}
	if _, ok := sc.mutation.TotalTime(); !ok {
		return &ValidationError{Name: "total_time", err: errors.New(`ent: missing required field "Submit.total_time"`)}
	}
	if _, ok := sc.mutation.MaxMemory(); !ok {
		return &ValidationError{Name: "max_memory", err: errors.New(`ent: missing required field "Submit.max_memory"`)}
	}
	if _, ok := sc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`ent: missing required field "Submit.language"`)}
	}
	if _, ok := sc.mutation.CaseVersion(); !ok {
		return &ValidationError{Name: "case_version", err: errors.New(`ent: missing required field "Submit.case_version"`)}
	}
	return nil
}

func (sc *SubmitCreate) sqlSave(ctx context.Context) (*Submit, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubmitCreate) createSpec() (*Submit, *sqlgraph.CreateSpec) {
	var (
		_node = &Submit{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(submit.Table, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.UserID(); ok {
		_spec.SetField(submit.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := sc.mutation.ProblemID(); ok {
		_spec.SetField(submit.FieldProblemID, field.TypeInt, value)
		_node.ProblemID = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.SetField(submit.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := sc.mutation.State(); ok {
		_spec.SetField(submit.FieldState, field.TypeInt, value)
		_node.State = value
	}
	if value, ok := sc.mutation.Point(); ok {
		_spec.SetField(submit.FieldPoint, field.TypeInt, value)
		_node.Point = value
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(submit.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.TotalTime(); ok {
		_spec.SetField(submit.FieldTotalTime, field.TypeInt, value)
		_node.TotalTime = value
	}
	if value, ok := sc.mutation.MaxMemory(); ok {
		_spec.SetField(submit.FieldMaxMemory, field.TypeInt, value)
		_node.MaxMemory = value
	}
	if value, ok := sc.mutation.Language(); ok {
		_spec.SetField(submit.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := sc.mutation.CaseVersion(); ok {
		_spec.SetField(submit.FieldCaseVersion, field.TypeInt, value)
		_node.CaseVersion = value
	}
	if nodes := sc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.UsersTable,
			Columns: []string{submit.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_submission = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submit.ProblemsTable,
			Columns: []string{submit.ProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.problem_submission = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SubmitJudgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.SubmitJudgeTable,
			Columns: []string{submit.SubmitJudgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submitjudge.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SubmitCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   submit.SubmitCasesTable,
			Columns: []string{submit.SubmitCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submitcase.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubmitCreateBulk is the builder for creating many Submit entities in bulk.
type SubmitCreateBulk struct {
	config
	err      error
	builders []*SubmitCreate
}

// Save creates the Submit entities in the database.
func (scb *SubmitCreateBulk) Save(ctx context.Context) ([]*Submit, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Submit, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubmitMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubmitCreateBulk) SaveX(ctx context.Context) []*Submit {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubmitCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubmitCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
