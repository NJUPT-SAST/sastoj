// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/group"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"
	"sastoj/ent/problemjudge"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemJudgeUpdate is the builder for updating ProblemJudge entities.
type ProblemJudgeUpdate struct {
	config
	hooks    []Hook
	mutation *ProblemJudgeMutation
}

// Where appends a list predicates to the ProblemJudgeUpdate builder.
func (pju *ProblemJudgeUpdate) Where(ps ...predicate.ProblemJudge) *ProblemJudgeUpdate {
	pju.mutation.Where(ps...)
	return pju
}

// SetGroupID sets the "group_id" field.
func (pju *ProblemJudgeUpdate) SetGroupID(i int) *ProblemJudgeUpdate {
	pju.mutation.ResetGroupID()
	pju.mutation.SetGroupID(i)
	return pju
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (pju *ProblemJudgeUpdate) SetNillableGroupID(i *int) *ProblemJudgeUpdate {
	if i != nil {
		pju.SetGroupID(*i)
	}
	return pju
}

// AddGroupID adds i to the "group_id" field.
func (pju *ProblemJudgeUpdate) AddGroupID(i int) *ProblemJudgeUpdate {
	pju.mutation.AddGroupID(i)
	return pju
}

// SetProblemID sets the "problem_id" field.
func (pju *ProblemJudgeUpdate) SetProblemID(i int) *ProblemJudgeUpdate {
	pju.mutation.ResetProblemID()
	pju.mutation.SetProblemID(i)
	return pju
}

// SetNillableProblemID sets the "problem_id" field if the given value is not nil.
func (pju *ProblemJudgeUpdate) SetNillableProblemID(i *int) *ProblemJudgeUpdate {
	if i != nil {
		pju.SetProblemID(*i)
	}
	return pju
}

// AddProblemID adds i to the "problem_id" field.
func (pju *ProblemJudgeUpdate) AddProblemID(i int) *ProblemJudgeUpdate {
	pju.mutation.AddProblemID(i)
	return pju
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (pju *ProblemJudgeUpdate) SetGroupID(id int) *ProblemJudgeUpdate {
	pju.mutation.SetGroupID(id)
	return pju
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (pju *ProblemJudgeUpdate) SetNillableGroupID(id *int) *ProblemJudgeUpdate {
	if id != nil {
		pju = pju.SetGroupID(*id)
	}
	return pju
}

// SetGroup sets the "group" edge to the Group entity.
func (pju *ProblemJudgeUpdate) SetGroup(g *Group) *ProblemJudgeUpdate {
	return pju.SetGroupID(g.ID)
}

// SetProblemID sets the "problem" edge to the Problem entity by ID.
func (pju *ProblemJudgeUpdate) SetProblemID(id int) *ProblemJudgeUpdate {
	pju.mutation.SetProblemID(id)
	return pju
}

// SetNillableProblemID sets the "problem" edge to the Problem entity by ID if the given value is not nil.
func (pju *ProblemJudgeUpdate) SetNillableProblemID(id *int) *ProblemJudgeUpdate {
	if id != nil {
		pju = pju.SetProblemID(*id)
	}
	return pju
}

// SetProblem sets the "problem" edge to the Problem entity.
func (pju *ProblemJudgeUpdate) SetProblem(p *Problem) *ProblemJudgeUpdate {
	return pju.SetProblemID(p.ID)
}

// Mutation returns the ProblemJudgeMutation object of the builder.
func (pju *ProblemJudgeUpdate) Mutation() *ProblemJudgeMutation {
	return pju.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (pju *ProblemJudgeUpdate) ClearGroup() *ProblemJudgeUpdate {
	pju.mutation.ClearGroup()
	return pju
}

// ClearProblem clears the "problem" edge to the Problem entity.
func (pju *ProblemJudgeUpdate) ClearProblem() *ProblemJudgeUpdate {
	pju.mutation.ClearProblem()
	return pju
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pju *ProblemJudgeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pju.sqlSave, pju.mutation, pju.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pju *ProblemJudgeUpdate) SaveX(ctx context.Context) int {
	affected, err := pju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pju *ProblemJudgeUpdate) Exec(ctx context.Context) error {
	_, err := pju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pju *ProblemJudgeUpdate) ExecX(ctx context.Context) {
	if err := pju.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pju *ProblemJudgeUpdate) check() error {
	if v, ok := pju.mutation.GroupID(); ok {
		if err := problemjudge.GroupIDValidator(v); err != nil {
			return &ValidationError{Name: "group_id", err: fmt.Errorf(`ent: validator failed for field "ProblemJudge.group_id": %w`, err)}
		}
	}
	if v, ok := pju.mutation.ProblemID(); ok {
		if err := problemjudge.ProblemIDValidator(v); err != nil {
			return &ValidationError{Name: "problem_id", err: fmt.Errorf(`ent: validator failed for field "ProblemJudge.problem_id": %w`, err)}
		}
	}
	return nil
}

func (pju *ProblemJudgeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pju.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(problemjudge.Table, problemjudge.Columns, sqlgraph.NewFieldSpec(problemjudge.FieldID, field.TypeInt))
	if ps := pju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pju.mutation.GroupID(); ok {
		_spec.SetField(problemjudge.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := pju.mutation.AddedGroupID(); ok {
		_spec.AddField(problemjudge.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := pju.mutation.ProblemID(); ok {
		_spec.SetField(problemjudge.FieldProblemID, field.TypeInt, value)
	}
	if value, ok := pju.mutation.AddedProblemID(); ok {
		_spec.AddField(problemjudge.FieldProblemID, field.TypeInt, value)
	}
	if pju.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.GroupTable,
			Columns: []string{problemjudge.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pju.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.GroupTable,
			Columns: []string{problemjudge.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pju.mutation.ProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.ProblemTable,
			Columns: []string{problemjudge.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pju.mutation.ProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.ProblemTable,
			Columns: []string{problemjudge.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problemjudge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pju.mutation.done = true
	return n, nil
}

// ProblemJudgeUpdateOne is the builder for updating a single ProblemJudge entity.
type ProblemJudgeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProblemJudgeMutation
}

// SetGroupID sets the "group_id" field.
func (pjuo *ProblemJudgeUpdateOne) SetGroupID(i int) *ProblemJudgeUpdateOne {
	pjuo.mutation.ResetGroupID()
	pjuo.mutation.SetGroupID(i)
	return pjuo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (pjuo *ProblemJudgeUpdateOne) SetNillableGroupID(i *int) *ProblemJudgeUpdateOne {
	if i != nil {
		pjuo.SetGroupID(*i)
	}
	return pjuo
}

// AddGroupID adds i to the "group_id" field.
func (pjuo *ProblemJudgeUpdateOne) AddGroupID(i int) *ProblemJudgeUpdateOne {
	pjuo.mutation.AddGroupID(i)
	return pjuo
}

// SetProblemID sets the "problem_id" field.
func (pjuo *ProblemJudgeUpdateOne) SetProblemID(i int) *ProblemJudgeUpdateOne {
	pjuo.mutation.ResetProblemID()
	pjuo.mutation.SetProblemID(i)
	return pjuo
}

// SetNillableProblemID sets the "problem_id" field if the given value is not nil.
func (pjuo *ProblemJudgeUpdateOne) SetNillableProblemID(i *int) *ProblemJudgeUpdateOne {
	if i != nil {
		pjuo.SetProblemID(*i)
	}
	return pjuo
}

// AddProblemID adds i to the "problem_id" field.
func (pjuo *ProblemJudgeUpdateOne) AddProblemID(i int) *ProblemJudgeUpdateOne {
	pjuo.mutation.AddProblemID(i)
	return pjuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (pjuo *ProblemJudgeUpdateOne) SetGroupID(id int) *ProblemJudgeUpdateOne {
	pjuo.mutation.SetGroupID(id)
	return pjuo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (pjuo *ProblemJudgeUpdateOne) SetNillableGroupID(id *int) *ProblemJudgeUpdateOne {
	if id != nil {
		pjuo = pjuo.SetGroupID(*id)
	}
	return pjuo
}

// SetGroup sets the "group" edge to the Group entity.
func (pjuo *ProblemJudgeUpdateOne) SetGroup(g *Group) *ProblemJudgeUpdateOne {
	return pjuo.SetGroupID(g.ID)
}

// SetProblemID sets the "problem" edge to the Problem entity by ID.
func (pjuo *ProblemJudgeUpdateOne) SetProblemID(id int) *ProblemJudgeUpdateOne {
	pjuo.mutation.SetProblemID(id)
	return pjuo
}

// SetNillableProblemID sets the "problem" edge to the Problem entity by ID if the given value is not nil.
func (pjuo *ProblemJudgeUpdateOne) SetNillableProblemID(id *int) *ProblemJudgeUpdateOne {
	if id != nil {
		pjuo = pjuo.SetProblemID(*id)
	}
	return pjuo
}

// SetProblem sets the "problem" edge to the Problem entity.
func (pjuo *ProblemJudgeUpdateOne) SetProblem(p *Problem) *ProblemJudgeUpdateOne {
	return pjuo.SetProblemID(p.ID)
}

// Mutation returns the ProblemJudgeMutation object of the builder.
func (pjuo *ProblemJudgeUpdateOne) Mutation() *ProblemJudgeMutation {
	return pjuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (pjuo *ProblemJudgeUpdateOne) ClearGroup() *ProblemJudgeUpdateOne {
	pjuo.mutation.ClearGroup()
	return pjuo
}

// ClearProblem clears the "problem" edge to the Problem entity.
func (pjuo *ProblemJudgeUpdateOne) ClearProblem() *ProblemJudgeUpdateOne {
	pjuo.mutation.ClearProblem()
	return pjuo
}

// Where appends a list predicates to the ProblemJudgeUpdate builder.
func (pjuo *ProblemJudgeUpdateOne) Where(ps ...predicate.ProblemJudge) *ProblemJudgeUpdateOne {
	pjuo.mutation.Where(ps...)
	return pjuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pjuo *ProblemJudgeUpdateOne) Select(field string, fields ...string) *ProblemJudgeUpdateOne {
	pjuo.fields = append([]string{field}, fields...)
	return pjuo
}

// Save executes the query and returns the updated ProblemJudge entity.
func (pjuo *ProblemJudgeUpdateOne) Save(ctx context.Context) (*ProblemJudge, error) {
	return withHooks(ctx, pjuo.sqlSave, pjuo.mutation, pjuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pjuo *ProblemJudgeUpdateOne) SaveX(ctx context.Context) *ProblemJudge {
	node, err := pjuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pjuo *ProblemJudgeUpdateOne) Exec(ctx context.Context) error {
	_, err := pjuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pjuo *ProblemJudgeUpdateOne) ExecX(ctx context.Context) {
	if err := pjuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pjuo *ProblemJudgeUpdateOne) check() error {
	if v, ok := pjuo.mutation.GroupID(); ok {
		if err := problemjudge.GroupIDValidator(v); err != nil {
			return &ValidationError{Name: "group_id", err: fmt.Errorf(`ent: validator failed for field "ProblemJudge.group_id": %w`, err)}
		}
	}
	if v, ok := pjuo.mutation.ProblemID(); ok {
		if err := problemjudge.ProblemIDValidator(v); err != nil {
			return &ValidationError{Name: "problem_id", err: fmt.Errorf(`ent: validator failed for field "ProblemJudge.problem_id": %w`, err)}
		}
	}
	return nil
}

func (pjuo *ProblemJudgeUpdateOne) sqlSave(ctx context.Context) (_node *ProblemJudge, err error) {
	if err := pjuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(problemjudge.Table, problemjudge.Columns, sqlgraph.NewFieldSpec(problemjudge.FieldID, field.TypeInt))
	id, ok := pjuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProblemJudge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pjuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problemjudge.FieldID)
		for _, f := range fields {
			if !problemjudge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != problemjudge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pjuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pjuo.mutation.GroupID(); ok {
		_spec.SetField(problemjudge.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := pjuo.mutation.AddedGroupID(); ok {
		_spec.AddField(problemjudge.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := pjuo.mutation.ProblemID(); ok {
		_spec.SetField(problemjudge.FieldProblemID, field.TypeInt, value)
	}
	if value, ok := pjuo.mutation.AddedProblemID(); ok {
		_spec.AddField(problemjudge.FieldProblemID, field.TypeInt, value)
	}
	if pjuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.GroupTable,
			Columns: []string{problemjudge.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pjuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.GroupTable,
			Columns: []string{problemjudge.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pjuo.mutation.ProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.ProblemTable,
			Columns: []string{problemjudge.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pjuo.mutation.ProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problemjudge.ProblemTable,
			Columns: []string{problemjudge.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProblemJudge{config: pjuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pjuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problemjudge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pjuo.mutation.done = true
	return _node, nil
}
