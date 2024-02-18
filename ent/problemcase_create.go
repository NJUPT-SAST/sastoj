// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"sastoj/ent/submitcase"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemCaseCreate is the builder for creating a ProblemCase entity.
type ProblemCaseCreate struct {
	config
	mutation *ProblemCaseMutation
	hooks    []Hook
}

// SetProblemID sets the "problem_id" field.
func (pcc *ProblemCaseCreate) SetProblemID(i int) *ProblemCaseCreate {
	pcc.mutation.SetProblemID(i)
	return pcc
}

// SetPoint sets the "point" field.
func (pcc *ProblemCaseCreate) SetPoint(i int) *ProblemCaseCreate {
	pcc.mutation.SetPoint(i)
	return pcc
}

// SetIndex sets the "index" field.
func (pcc *ProblemCaseCreate) SetIndex(i int) *ProblemCaseCreate {
	pcc.mutation.SetIndex(i)
	return pcc
}

// SetIsAuto sets the "is_auto" field.
func (pcc *ProblemCaseCreate) SetIsAuto(b bool) *ProblemCaseCreate {
	pcc.mutation.SetIsAuto(b)
	return pcc
}

// SetNillableIsAuto sets the "is_auto" field if the given value is not nil.
func (pcc *ProblemCaseCreate) SetNillableIsAuto(b *bool) *ProblemCaseCreate {
	if b != nil {
		pcc.SetIsAuto(*b)
	}
	return pcc
}

// SetIsDeleted sets the "is_deleted" field.
func (pcc *ProblemCaseCreate) SetIsDeleted(b bool) *ProblemCaseCreate {
	pcc.mutation.SetIsDeleted(b)
	return pcc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pcc *ProblemCaseCreate) SetNillableIsDeleted(b *bool) *ProblemCaseCreate {
	if b != nil {
		pcc.SetIsDeleted(*b)
	}
	return pcc
}

// SetID sets the "id" field.
func (pcc *ProblemCaseCreate) SetID(i int) *ProblemCaseCreate {
	pcc.mutation.SetID(i)
	return pcc
}

// SetProblemsID sets the "problems" edge to the Problem entity by ID.
func (pcc *ProblemCaseCreate) SetProblemsID(id int) *ProblemCaseCreate {
	pcc.mutation.SetProblemsID(id)
	return pcc
}

// SetNillableProblemsID sets the "problems" edge to the Problem entity by ID if the given value is not nil.
func (pcc *ProblemCaseCreate) SetNillableProblemsID(id *int) *ProblemCaseCreate {
	if id != nil {
		pcc = pcc.SetProblemsID(*id)
	}
	return pcc
}

// SetProblems sets the "problems" edge to the Problem entity.
func (pcc *ProblemCaseCreate) SetProblems(p *Problem) *ProblemCaseCreate {
	return pcc.SetProblemsID(p.ID)
}

// AddSubmitCaseIDs adds the "submit_cases" edge to the SubmitCase entity by IDs.
func (pcc *ProblemCaseCreate) AddSubmitCaseIDs(ids ...int) *ProblemCaseCreate {
	pcc.mutation.AddSubmitCaseIDs(ids...)
	return pcc
}

// AddSubmitCases adds the "submit_cases" edges to the SubmitCase entity.
func (pcc *ProblemCaseCreate) AddSubmitCases(s ...*SubmitCase) *ProblemCaseCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pcc.AddSubmitCaseIDs(ids...)
}

// Mutation returns the ProblemCaseMutation object of the builder.
func (pcc *ProblemCaseCreate) Mutation() *ProblemCaseMutation {
	return pcc.mutation
}

// Save creates the ProblemCase in the database.
func (pcc *ProblemCaseCreate) Save(ctx context.Context) (*ProblemCase, error) {
	pcc.defaults()
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *ProblemCaseCreate) SaveX(ctx context.Context) *ProblemCase {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *ProblemCaseCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *ProblemCaseCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *ProblemCaseCreate) defaults() {
	if _, ok := pcc.mutation.IsAuto(); !ok {
		v := problemcase.DefaultIsAuto
		pcc.mutation.SetIsAuto(v)
	}
	if _, ok := pcc.mutation.IsDeleted(); !ok {
		v := problemcase.DefaultIsDeleted
		pcc.mutation.SetIsDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *ProblemCaseCreate) check() error {
	if _, ok := pcc.mutation.ProblemID(); !ok {
		return &ValidationError{Name: "problem_id", err: errors.New(`ent: missing required field "ProblemCase.problem_id"`)}
	}
	if v, ok := pcc.mutation.ProblemID(); ok {
		if err := problemcase.ProblemIDValidator(v); err != nil {
			return &ValidationError{Name: "problem_id", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.problem_id": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "ProblemCase.point"`)}
	}
	if _, ok := pcc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "ProblemCase.index"`)}
	}
	if v, ok := pcc.mutation.Index(); ok {
		if err := problemcase.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.index": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.IsAuto(); !ok {
		return &ValidationError{Name: "is_auto", err: errors.New(`ent: missing required field "ProblemCase.is_auto"`)}
	}
	if _, ok := pcc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "ProblemCase.is_deleted"`)}
	}
	return nil
}

func (pcc *ProblemCaseCreate) sqlSave(ctx context.Context) (*ProblemCase, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *ProblemCaseCreate) createSpec() (*ProblemCase, *sqlgraph.CreateSpec) {
	var (
		_node = &ProblemCase{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(problemcase.Table, sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt))
	)
	if id, ok := pcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pcc.mutation.ProblemID(); ok {
		_spec.SetField(problemcase.FieldProblemID, field.TypeInt, value)
		_node.ProblemID = value
	}
	if value, ok := pcc.mutation.Point(); ok {
		_spec.SetField(problemcase.FieldPoint, field.TypeInt, value)
		_node.Point = value
	}
	if value, ok := pcc.mutation.Index(); ok {
		_spec.SetField(problemcase.FieldIndex, field.TypeInt, value)
		_node.Index = value
	}
	if value, ok := pcc.mutation.IsAuto(); ok {
		_spec.SetField(problemcase.FieldIsAuto, field.TypeBool, value)
		_node.IsAuto = value
	}
	if value, ok := pcc.mutation.IsDeleted(); ok {
		_spec.SetField(problemcase.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = value
	}
	if nodes := pcc.mutation.ProblemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   problemcase.ProblemsTable,
			Columns: []string{problemcase.ProblemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.problem_problem_cases = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.SubmitCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   problemcase.SubmitCasesTable,
			Columns: []string{problemcase.SubmitCasesColumn},
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

// ProblemCaseCreateBulk is the builder for creating many ProblemCase entities in bulk.
type ProblemCaseCreateBulk struct {
	config
	err      error
	builders []*ProblemCaseCreate
}

// Save creates the ProblemCase entities in the database.
func (pccb *ProblemCaseCreateBulk) Save(ctx context.Context) ([]*ProblemCase, error) {
	if pccb.err != nil {
		return nil, pccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*ProblemCase, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProblemCaseMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *ProblemCaseCreateBulk) SaveX(ctx context.Context) []*ProblemCase {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *ProblemCaseCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *ProblemCaseCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
