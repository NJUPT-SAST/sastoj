// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"sastoj/ent/submitcase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemCaseUpdate is the builder for updating ProblemCase entities.
type ProblemCaseUpdate struct {
	config
	hooks    []Hook
	mutation *ProblemCaseMutation
}

// Where appends a list predicates to the ProblemCaseUpdate builder.
func (pcu *ProblemCaseUpdate) Where(ps ...predicate.ProblemCase) *ProblemCaseUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetPoint sets the "point" field.
func (pcu *ProblemCaseUpdate) SetPoint(i int16) *ProblemCaseUpdate {
	pcu.mutation.ResetPoint()
	pcu.mutation.SetPoint(i)
	return pcu
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (pcu *ProblemCaseUpdate) SetNillablePoint(i *int16) *ProblemCaseUpdate {
	if i != nil {
		pcu.SetPoint(*i)
	}
	return pcu
}

// AddPoint adds i to the "point" field.
func (pcu *ProblemCaseUpdate) AddPoint(i int16) *ProblemCaseUpdate {
	pcu.mutation.AddPoint(i)
	return pcu
}

// SetIndex sets the "index" field.
func (pcu *ProblemCaseUpdate) SetIndex(i int16) *ProblemCaseUpdate {
	pcu.mutation.ResetIndex()
	pcu.mutation.SetIndex(i)
	return pcu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (pcu *ProblemCaseUpdate) SetNillableIndex(i *int16) *ProblemCaseUpdate {
	if i != nil {
		pcu.SetIndex(*i)
	}
	return pcu
}

// AddIndex adds i to the "index" field.
func (pcu *ProblemCaseUpdate) AddIndex(i int16) *ProblemCaseUpdate {
	pcu.mutation.AddIndex(i)
	return pcu
}

// SetIsAuto sets the "is_auto" field.
func (pcu *ProblemCaseUpdate) SetIsAuto(b bool) *ProblemCaseUpdate {
	pcu.mutation.SetIsAuto(b)
	return pcu
}

// SetNillableIsAuto sets the "is_auto" field if the given value is not nil.
func (pcu *ProblemCaseUpdate) SetNillableIsAuto(b *bool) *ProblemCaseUpdate {
	if b != nil {
		pcu.SetIsAuto(*b)
	}
	return pcu
}

// SetIsDeleted sets the "is_deleted" field.
func (pcu *ProblemCaseUpdate) SetIsDeleted(b bool) *ProblemCaseUpdate {
	pcu.mutation.SetIsDeleted(b)
	return pcu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pcu *ProblemCaseUpdate) SetNillableIsDeleted(b *bool) *ProblemCaseUpdate {
	if b != nil {
		pcu.SetIsDeleted(*b)
	}
	return pcu
}

// SetProblemID sets the "problem_id" field.
func (pcu *ProblemCaseUpdate) SetProblemID(i int) *ProblemCaseUpdate {
	pcu.mutation.SetProblemID(i)
	return pcu
}

// SetNillableProblemID sets the "problem_id" field if the given value is not nil.
func (pcu *ProblemCaseUpdate) SetNillableProblemID(i *int) *ProblemCaseUpdate {
	if i != nil {
		pcu.SetProblemID(*i)
	}
	return pcu
}

// AddSubmitCaseIDs adds the "submit_cases" edge to the SubmitCase entity by IDs.
func (pcu *ProblemCaseUpdate) AddSubmitCaseIDs(ids ...int) *ProblemCaseUpdate {
	pcu.mutation.AddSubmitCaseIDs(ids...)
	return pcu
}

// AddSubmitCases adds the "submit_cases" edges to the SubmitCase entity.
func (pcu *ProblemCaseUpdate) AddSubmitCases(s ...*SubmitCase) *ProblemCaseUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pcu.AddSubmitCaseIDs(ids...)
}

// SetProblemsID sets the "problems" edge to the Problem entity by ID.
func (pcu *ProblemCaseUpdate) SetProblemsID(id int) *ProblemCaseUpdate {
	pcu.mutation.SetProblemsID(id)
	return pcu
}

// SetProblems sets the "problems" edge to the Problem entity.
func (pcu *ProblemCaseUpdate) SetProblems(p *Problem) *ProblemCaseUpdate {
	return pcu.SetProblemsID(p.ID)
}

// Mutation returns the ProblemCaseMutation object of the builder.
func (pcu *ProblemCaseUpdate) Mutation() *ProblemCaseMutation {
	return pcu.mutation
}

// ClearSubmitCases clears all "submit_cases" edges to the SubmitCase entity.
func (pcu *ProblemCaseUpdate) ClearSubmitCases() *ProblemCaseUpdate {
	pcu.mutation.ClearSubmitCases()
	return pcu
}

// RemoveSubmitCaseIDs removes the "submit_cases" edge to SubmitCase entities by IDs.
func (pcu *ProblemCaseUpdate) RemoveSubmitCaseIDs(ids ...int) *ProblemCaseUpdate {
	pcu.mutation.RemoveSubmitCaseIDs(ids...)
	return pcu
}

// RemoveSubmitCases removes "submit_cases" edges to SubmitCase entities.
func (pcu *ProblemCaseUpdate) RemoveSubmitCases(s ...*SubmitCase) *ProblemCaseUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pcu.RemoveSubmitCaseIDs(ids...)
}

// ClearProblems clears the "problems" edge to the Problem entity.
func (pcu *ProblemCaseUpdate) ClearProblems() *ProblemCaseUpdate {
	pcu.mutation.ClearProblems()
	return pcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProblemCaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProblemCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProblemCaseUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProblemCaseUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcu *ProblemCaseUpdate) check() error {
	if v, ok := pcu.mutation.Point(); ok {
		if err := problemcase.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.point": %w`, err)}
		}
	}
	if v, ok := pcu.mutation.Index(); ok {
		if err := problemcase.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.index": %w`, err)}
		}
	}
	if _, ok := pcu.mutation.ProblemsID(); pcu.mutation.ProblemsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProblemCase.problems"`)
	}
	return nil
}

func (pcu *ProblemCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(problemcase.Table, problemcase.Columns, sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Point(); ok {
		_spec.SetField(problemcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := pcu.mutation.AddedPoint(); ok {
		_spec.AddField(problemcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := pcu.mutation.Index(); ok {
		_spec.SetField(problemcase.FieldIndex, field.TypeInt16, value)
	}
	if value, ok := pcu.mutation.AddedIndex(); ok {
		_spec.AddField(problemcase.FieldIndex, field.TypeInt16, value)
	}
	if value, ok := pcu.mutation.IsAuto(); ok {
		_spec.SetField(problemcase.FieldIsAuto, field.TypeBool, value)
	}
	if value, ok := pcu.mutation.IsDeleted(); ok {
		_spec.SetField(problemcase.FieldIsDeleted, field.TypeBool, value)
	}
	if pcu.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedSubmitCasesIDs(); len(nodes) > 0 && !pcu.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.SubmitCasesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.ProblemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ProblemsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problemcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// ProblemCaseUpdateOne is the builder for updating a single ProblemCase entity.
type ProblemCaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProblemCaseMutation
}

// SetPoint sets the "point" field.
func (pcuo *ProblemCaseUpdateOne) SetPoint(i int16) *ProblemCaseUpdateOne {
	pcuo.mutation.ResetPoint()
	pcuo.mutation.SetPoint(i)
	return pcuo
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (pcuo *ProblemCaseUpdateOne) SetNillablePoint(i *int16) *ProblemCaseUpdateOne {
	if i != nil {
		pcuo.SetPoint(*i)
	}
	return pcuo
}

// AddPoint adds i to the "point" field.
func (pcuo *ProblemCaseUpdateOne) AddPoint(i int16) *ProblemCaseUpdateOne {
	pcuo.mutation.AddPoint(i)
	return pcuo
}

// SetIndex sets the "index" field.
func (pcuo *ProblemCaseUpdateOne) SetIndex(i int16) *ProblemCaseUpdateOne {
	pcuo.mutation.ResetIndex()
	pcuo.mutation.SetIndex(i)
	return pcuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (pcuo *ProblemCaseUpdateOne) SetNillableIndex(i *int16) *ProblemCaseUpdateOne {
	if i != nil {
		pcuo.SetIndex(*i)
	}
	return pcuo
}

// AddIndex adds i to the "index" field.
func (pcuo *ProblemCaseUpdateOne) AddIndex(i int16) *ProblemCaseUpdateOne {
	pcuo.mutation.AddIndex(i)
	return pcuo
}

// SetIsAuto sets the "is_auto" field.
func (pcuo *ProblemCaseUpdateOne) SetIsAuto(b bool) *ProblemCaseUpdateOne {
	pcuo.mutation.SetIsAuto(b)
	return pcuo
}

// SetNillableIsAuto sets the "is_auto" field if the given value is not nil.
func (pcuo *ProblemCaseUpdateOne) SetNillableIsAuto(b *bool) *ProblemCaseUpdateOne {
	if b != nil {
		pcuo.SetIsAuto(*b)
	}
	return pcuo
}

// SetIsDeleted sets the "is_deleted" field.
func (pcuo *ProblemCaseUpdateOne) SetIsDeleted(b bool) *ProblemCaseUpdateOne {
	pcuo.mutation.SetIsDeleted(b)
	return pcuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pcuo *ProblemCaseUpdateOne) SetNillableIsDeleted(b *bool) *ProblemCaseUpdateOne {
	if b != nil {
		pcuo.SetIsDeleted(*b)
	}
	return pcuo
}

// SetProblemID sets the "problem_id" field.
func (pcuo *ProblemCaseUpdateOne) SetProblemID(i int) *ProblemCaseUpdateOne {
	pcuo.mutation.SetProblemID(i)
	return pcuo
}

// SetNillableProblemID sets the "problem_id" field if the given value is not nil.
func (pcuo *ProblemCaseUpdateOne) SetNillableProblemID(i *int) *ProblemCaseUpdateOne {
	if i != nil {
		pcuo.SetProblemID(*i)
	}
	return pcuo
}

// AddSubmitCaseIDs adds the "submit_cases" edge to the SubmitCase entity by IDs.
func (pcuo *ProblemCaseUpdateOne) AddSubmitCaseIDs(ids ...int) *ProblemCaseUpdateOne {
	pcuo.mutation.AddSubmitCaseIDs(ids...)
	return pcuo
}

// AddSubmitCases adds the "submit_cases" edges to the SubmitCase entity.
func (pcuo *ProblemCaseUpdateOne) AddSubmitCases(s ...*SubmitCase) *ProblemCaseUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pcuo.AddSubmitCaseIDs(ids...)
}

// SetProblemsID sets the "problems" edge to the Problem entity by ID.
func (pcuo *ProblemCaseUpdateOne) SetProblemsID(id int) *ProblemCaseUpdateOne {
	pcuo.mutation.SetProblemsID(id)
	return pcuo
}

// SetProblems sets the "problems" edge to the Problem entity.
func (pcuo *ProblemCaseUpdateOne) SetProblems(p *Problem) *ProblemCaseUpdateOne {
	return pcuo.SetProblemsID(p.ID)
}

// Mutation returns the ProblemCaseMutation object of the builder.
func (pcuo *ProblemCaseUpdateOne) Mutation() *ProblemCaseMutation {
	return pcuo.mutation
}

// ClearSubmitCases clears all "submit_cases" edges to the SubmitCase entity.
func (pcuo *ProblemCaseUpdateOne) ClearSubmitCases() *ProblemCaseUpdateOne {
	pcuo.mutation.ClearSubmitCases()
	return pcuo
}

// RemoveSubmitCaseIDs removes the "submit_cases" edge to SubmitCase entities by IDs.
func (pcuo *ProblemCaseUpdateOne) RemoveSubmitCaseIDs(ids ...int) *ProblemCaseUpdateOne {
	pcuo.mutation.RemoveSubmitCaseIDs(ids...)
	return pcuo
}

// RemoveSubmitCases removes "submit_cases" edges to SubmitCase entities.
func (pcuo *ProblemCaseUpdateOne) RemoveSubmitCases(s ...*SubmitCase) *ProblemCaseUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pcuo.RemoveSubmitCaseIDs(ids...)
}

// ClearProblems clears the "problems" edge to the Problem entity.
func (pcuo *ProblemCaseUpdateOne) ClearProblems() *ProblemCaseUpdateOne {
	pcuo.mutation.ClearProblems()
	return pcuo
}

// Where appends a list predicates to the ProblemCaseUpdate builder.
func (pcuo *ProblemCaseUpdateOne) Where(ps ...predicate.ProblemCase) *ProblemCaseUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProblemCaseUpdateOne) Select(field string, fields ...string) *ProblemCaseUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProblemCase entity.
func (pcuo *ProblemCaseUpdateOne) Save(ctx context.Context) (*ProblemCase, error) {
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProblemCaseUpdateOne) SaveX(ctx context.Context) *ProblemCase {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProblemCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProblemCaseUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcuo *ProblemCaseUpdateOne) check() error {
	if v, ok := pcuo.mutation.Point(); ok {
		if err := problemcase.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.point": %w`, err)}
		}
	}
	if v, ok := pcuo.mutation.Index(); ok {
		if err := problemcase.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.index": %w`, err)}
		}
	}
	if _, ok := pcuo.mutation.ProblemsID(); pcuo.mutation.ProblemsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProblemCase.problems"`)
	}
	return nil
}

func (pcuo *ProblemCaseUpdateOne) sqlSave(ctx context.Context) (_node *ProblemCase, err error) {
	if err := pcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(problemcase.Table, problemcase.Columns, sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProblemCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problemcase.FieldID)
		for _, f := range fields {
			if !problemcase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != problemcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Point(); ok {
		_spec.SetField(problemcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := pcuo.mutation.AddedPoint(); ok {
		_spec.AddField(problemcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := pcuo.mutation.Index(); ok {
		_spec.SetField(problemcase.FieldIndex, field.TypeInt16, value)
	}
	if value, ok := pcuo.mutation.AddedIndex(); ok {
		_spec.AddField(problemcase.FieldIndex, field.TypeInt16, value)
	}
	if value, ok := pcuo.mutation.IsAuto(); ok {
		_spec.SetField(problemcase.FieldIsAuto, field.TypeBool, value)
	}
	if value, ok := pcuo.mutation.IsDeleted(); ok {
		_spec.SetField(problemcase.FieldIsDeleted, field.TypeBool, value)
	}
	if pcuo.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedSubmitCasesIDs(); len(nodes) > 0 && !pcuo.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.SubmitCasesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.ProblemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ProblemsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProblemCase{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problemcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
