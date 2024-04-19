// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/predicate"
	"sastoj/ent/problemcase"
	"sastoj/ent/submit"
	"sastoj/ent/submitcase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmitCaseUpdate is the builder for updating SubmitCase entities.
type SubmitCaseUpdate struct {
	config
	hooks    []Hook
	mutation *SubmitCaseMutation
}

// Where appends a list predicates to the SubmitCaseUpdate builder.
func (scu *SubmitCaseUpdate) Where(ps ...predicate.SubmitCase) *SubmitCaseUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetState sets the "state" field.
func (scu *SubmitCaseUpdate) SetState(i int16) *SubmitCaseUpdate {
	scu.mutation.ResetState()
	scu.mutation.SetState(i)
	return scu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillableState(i *int16) *SubmitCaseUpdate {
	if i != nil {
		scu.SetState(*i)
	}
	return scu
}

// AddState adds i to the "state" field.
func (scu *SubmitCaseUpdate) AddState(i int16) *SubmitCaseUpdate {
	scu.mutation.AddState(i)
	return scu
}

// SetPoint sets the "point" field.
func (scu *SubmitCaseUpdate) SetPoint(i int16) *SubmitCaseUpdate {
	scu.mutation.ResetPoint()
	scu.mutation.SetPoint(i)
	return scu
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillablePoint(i *int16) *SubmitCaseUpdate {
	if i != nil {
		scu.SetPoint(*i)
	}
	return scu
}

// AddPoint adds i to the "point" field.
func (scu *SubmitCaseUpdate) AddPoint(i int16) *SubmitCaseUpdate {
	scu.mutation.AddPoint(i)
	return scu
}

// SetMessage sets the "message" field.
func (scu *SubmitCaseUpdate) SetMessage(s string) *SubmitCaseUpdate {
	scu.mutation.SetMessage(s)
	return scu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillableMessage(s *string) *SubmitCaseUpdate {
	if s != nil {
		scu.SetMessage(*s)
	}
	return scu
}

// SetTime sets the "time" field.
func (scu *SubmitCaseUpdate) SetTime(i int32) *SubmitCaseUpdate {
	scu.mutation.ResetTime()
	scu.mutation.SetTime(i)
	return scu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillableTime(i *int32) *SubmitCaseUpdate {
	if i != nil {
		scu.SetTime(*i)
	}
	return scu
}

// AddTime adds i to the "time" field.
func (scu *SubmitCaseUpdate) AddTime(i int32) *SubmitCaseUpdate {
	scu.mutation.AddTime(i)
	return scu
}

// SetMemory sets the "memory" field.
func (scu *SubmitCaseUpdate) SetMemory(i int32) *SubmitCaseUpdate {
	scu.mutation.ResetMemory()
	scu.mutation.SetMemory(i)
	return scu
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillableMemory(i *int32) *SubmitCaseUpdate {
	if i != nil {
		scu.SetMemory(*i)
	}
	return scu
}

// AddMemory adds i to the "memory" field.
func (scu *SubmitCaseUpdate) AddMemory(i int32) *SubmitCaseUpdate {
	scu.mutation.AddMemory(i)
	return scu
}

// SetSubmitID sets the "submit_id" field.
func (scu *SubmitCaseUpdate) SetSubmitID(i int64) *SubmitCaseUpdate {
	scu.mutation.SetSubmitID(i)
	return scu
}

// SetNillableSubmitID sets the "submit_id" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillableSubmitID(i *int64) *SubmitCaseUpdate {
	if i != nil {
		scu.SetSubmitID(*i)
	}
	return scu
}

// SetProblemCaseID sets the "problem_case_id" field.
func (scu *SubmitCaseUpdate) SetProblemCaseID(i int64) *SubmitCaseUpdate {
	scu.mutation.SetProblemCaseID(i)
	return scu
}

// SetNillableProblemCaseID sets the "problem_case_id" field if the given value is not nil.
func (scu *SubmitCaseUpdate) SetNillableProblemCaseID(i *int64) *SubmitCaseUpdate {
	if i != nil {
		scu.SetProblemCaseID(*i)
	}
	return scu
}

// SetSubmissionID sets the "submission" edge to the Submit entity by ID.
func (scu *SubmitCaseUpdate) SetSubmissionID(id int64) *SubmitCaseUpdate {
	scu.mutation.SetSubmissionID(id)
	return scu
}

// SetSubmission sets the "submission" edge to the Submit entity.
func (scu *SubmitCaseUpdate) SetSubmission(s *Submit) *SubmitCaseUpdate {
	return scu.SetSubmissionID(s.ID)
}

// SetProblemCasesID sets the "problem_cases" edge to the ProblemCase entity by ID.
func (scu *SubmitCaseUpdate) SetProblemCasesID(id int64) *SubmitCaseUpdate {
	scu.mutation.SetProblemCasesID(id)
	return scu
}

// SetProblemCases sets the "problem_cases" edge to the ProblemCase entity.
func (scu *SubmitCaseUpdate) SetProblemCases(p *ProblemCase) *SubmitCaseUpdate {
	return scu.SetProblemCasesID(p.ID)
}

// Mutation returns the SubmitCaseMutation object of the builder.
func (scu *SubmitCaseUpdate) Mutation() *SubmitCaseMutation {
	return scu.mutation
}

// ClearSubmission clears the "submission" edge to the Submit entity.
func (scu *SubmitCaseUpdate) ClearSubmission() *SubmitCaseUpdate {
	scu.mutation.ClearSubmission()
	return scu
}

// ClearProblemCases clears the "problem_cases" edge to the ProblemCase entity.
func (scu *SubmitCaseUpdate) ClearProblemCases() *SubmitCaseUpdate {
	scu.mutation.ClearProblemCases()
	return scu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *SubmitCaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *SubmitCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *SubmitCaseUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *SubmitCaseUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *SubmitCaseUpdate) check() error {
	if v, ok := scu.mutation.State(); ok {
		if err := submitcase.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.state": %w`, err)}
		}
	}
	if v, ok := scu.mutation.Point(); ok {
		if err := submitcase.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.point": %w`, err)}
		}
	}
	if v, ok := scu.mutation.Time(); ok {
		if err := submitcase.TimeValidator(v); err != nil {
			return &ValidationError{Name: "time", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.time": %w`, err)}
		}
	}
	if v, ok := scu.mutation.Memory(); ok {
		if err := submitcase.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.memory": %w`, err)}
		}
	}
	if _, ok := scu.mutation.SubmissionID(); scu.mutation.SubmissionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SubmitCase.submission"`)
	}
	if _, ok := scu.mutation.ProblemCasesID(); scu.mutation.ProblemCasesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SubmitCase.problem_cases"`)
	}
	return nil
}

func (scu *SubmitCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(submitcase.Table, submitcase.Columns, sqlgraph.NewFieldSpec(submitcase.FieldID, field.TypeInt64))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.State(); ok {
		_spec.SetField(submitcase.FieldState, field.TypeInt16, value)
	}
	if value, ok := scu.mutation.AddedState(); ok {
		_spec.AddField(submitcase.FieldState, field.TypeInt16, value)
	}
	if value, ok := scu.mutation.Point(); ok {
		_spec.SetField(submitcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := scu.mutation.AddedPoint(); ok {
		_spec.AddField(submitcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := scu.mutation.Message(); ok {
		_spec.SetField(submitcase.FieldMessage, field.TypeString, value)
	}
	if value, ok := scu.mutation.Time(); ok {
		_spec.SetField(submitcase.FieldTime, field.TypeInt32, value)
	}
	if value, ok := scu.mutation.AddedTime(); ok {
		_spec.AddField(submitcase.FieldTime, field.TypeInt32, value)
	}
	if value, ok := scu.mutation.Memory(); ok {
		_spec.SetField(submitcase.FieldMemory, field.TypeInt32, value)
	}
	if value, ok := scu.mutation.AddedMemory(); ok {
		_spec.AddField(submitcase.FieldMemory, field.TypeInt32, value)
	}
	if scu.mutation.SubmissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.SubmissionTable,
			Columns: []string{submitcase.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.SubmissionTable,
			Columns: []string{submitcase.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scu.mutation.ProblemCasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.ProblemCasesTable,
			Columns: []string{submitcase.ProblemCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.ProblemCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.ProblemCasesTable,
			Columns: []string{submitcase.ProblemCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submitcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// SubmitCaseUpdateOne is the builder for updating a single SubmitCase entity.
type SubmitCaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubmitCaseMutation
}

// SetState sets the "state" field.
func (scuo *SubmitCaseUpdateOne) SetState(i int16) *SubmitCaseUpdateOne {
	scuo.mutation.ResetState()
	scuo.mutation.SetState(i)
	return scuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillableState(i *int16) *SubmitCaseUpdateOne {
	if i != nil {
		scuo.SetState(*i)
	}
	return scuo
}

// AddState adds i to the "state" field.
func (scuo *SubmitCaseUpdateOne) AddState(i int16) *SubmitCaseUpdateOne {
	scuo.mutation.AddState(i)
	return scuo
}

// SetPoint sets the "point" field.
func (scuo *SubmitCaseUpdateOne) SetPoint(i int16) *SubmitCaseUpdateOne {
	scuo.mutation.ResetPoint()
	scuo.mutation.SetPoint(i)
	return scuo
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillablePoint(i *int16) *SubmitCaseUpdateOne {
	if i != nil {
		scuo.SetPoint(*i)
	}
	return scuo
}

// AddPoint adds i to the "point" field.
func (scuo *SubmitCaseUpdateOne) AddPoint(i int16) *SubmitCaseUpdateOne {
	scuo.mutation.AddPoint(i)
	return scuo
}

// SetMessage sets the "message" field.
func (scuo *SubmitCaseUpdateOne) SetMessage(s string) *SubmitCaseUpdateOne {
	scuo.mutation.SetMessage(s)
	return scuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillableMessage(s *string) *SubmitCaseUpdateOne {
	if s != nil {
		scuo.SetMessage(*s)
	}
	return scuo
}

// SetTime sets the "time" field.
func (scuo *SubmitCaseUpdateOne) SetTime(i int32) *SubmitCaseUpdateOne {
	scuo.mutation.ResetTime()
	scuo.mutation.SetTime(i)
	return scuo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillableTime(i *int32) *SubmitCaseUpdateOne {
	if i != nil {
		scuo.SetTime(*i)
	}
	return scuo
}

// AddTime adds i to the "time" field.
func (scuo *SubmitCaseUpdateOne) AddTime(i int32) *SubmitCaseUpdateOne {
	scuo.mutation.AddTime(i)
	return scuo
}

// SetMemory sets the "memory" field.
func (scuo *SubmitCaseUpdateOne) SetMemory(i int32) *SubmitCaseUpdateOne {
	scuo.mutation.ResetMemory()
	scuo.mutation.SetMemory(i)
	return scuo
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillableMemory(i *int32) *SubmitCaseUpdateOne {
	if i != nil {
		scuo.SetMemory(*i)
	}
	return scuo
}

// AddMemory adds i to the "memory" field.
func (scuo *SubmitCaseUpdateOne) AddMemory(i int32) *SubmitCaseUpdateOne {
	scuo.mutation.AddMemory(i)
	return scuo
}

// SetSubmitID sets the "submit_id" field.
func (scuo *SubmitCaseUpdateOne) SetSubmitID(i int64) *SubmitCaseUpdateOne {
	scuo.mutation.SetSubmitID(i)
	return scuo
}

// SetNillableSubmitID sets the "submit_id" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillableSubmitID(i *int64) *SubmitCaseUpdateOne {
	if i != nil {
		scuo.SetSubmitID(*i)
	}
	return scuo
}

// SetProblemCaseID sets the "problem_case_id" field.
func (scuo *SubmitCaseUpdateOne) SetProblemCaseID(i int64) *SubmitCaseUpdateOne {
	scuo.mutation.SetProblemCaseID(i)
	return scuo
}

// SetNillableProblemCaseID sets the "problem_case_id" field if the given value is not nil.
func (scuo *SubmitCaseUpdateOne) SetNillableProblemCaseID(i *int64) *SubmitCaseUpdateOne {
	if i != nil {
		scuo.SetProblemCaseID(*i)
	}
	return scuo
}

// SetSubmissionID sets the "submission" edge to the Submit entity by ID.
func (scuo *SubmitCaseUpdateOne) SetSubmissionID(id int64) *SubmitCaseUpdateOne {
	scuo.mutation.SetSubmissionID(id)
	return scuo
}

// SetSubmission sets the "submission" edge to the Submit entity.
func (scuo *SubmitCaseUpdateOne) SetSubmission(s *Submit) *SubmitCaseUpdateOne {
	return scuo.SetSubmissionID(s.ID)
}

// SetProblemCasesID sets the "problem_cases" edge to the ProblemCase entity by ID.
func (scuo *SubmitCaseUpdateOne) SetProblemCasesID(id int64) *SubmitCaseUpdateOne {
	scuo.mutation.SetProblemCasesID(id)
	return scuo
}

// SetProblemCases sets the "problem_cases" edge to the ProblemCase entity.
func (scuo *SubmitCaseUpdateOne) SetProblemCases(p *ProblemCase) *SubmitCaseUpdateOne {
	return scuo.SetProblemCasesID(p.ID)
}

// Mutation returns the SubmitCaseMutation object of the builder.
func (scuo *SubmitCaseUpdateOne) Mutation() *SubmitCaseMutation {
	return scuo.mutation
}

// ClearSubmission clears the "submission" edge to the Submit entity.
func (scuo *SubmitCaseUpdateOne) ClearSubmission() *SubmitCaseUpdateOne {
	scuo.mutation.ClearSubmission()
	return scuo
}

// ClearProblemCases clears the "problem_cases" edge to the ProblemCase entity.
func (scuo *SubmitCaseUpdateOne) ClearProblemCases() *SubmitCaseUpdateOne {
	scuo.mutation.ClearProblemCases()
	return scuo
}

// Where appends a list predicates to the SubmitCaseUpdate builder.
func (scuo *SubmitCaseUpdateOne) Where(ps ...predicate.SubmitCase) *SubmitCaseUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *SubmitCaseUpdateOne) Select(field string, fields ...string) *SubmitCaseUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated SubmitCase entity.
func (scuo *SubmitCaseUpdateOne) Save(ctx context.Context) (*SubmitCase, error) {
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *SubmitCaseUpdateOne) SaveX(ctx context.Context) *SubmitCase {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *SubmitCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *SubmitCaseUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *SubmitCaseUpdateOne) check() error {
	if v, ok := scuo.mutation.State(); ok {
		if err := submitcase.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.state": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.Point(); ok {
		if err := submitcase.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.point": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.Time(); ok {
		if err := submitcase.TimeValidator(v); err != nil {
			return &ValidationError{Name: "time", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.time": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.Memory(); ok {
		if err := submitcase.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "SubmitCase.memory": %w`, err)}
		}
	}
	if _, ok := scuo.mutation.SubmissionID(); scuo.mutation.SubmissionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SubmitCase.submission"`)
	}
	if _, ok := scuo.mutation.ProblemCasesID(); scuo.mutation.ProblemCasesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SubmitCase.problem_cases"`)
	}
	return nil
}

func (scuo *SubmitCaseUpdateOne) sqlSave(ctx context.Context) (_node *SubmitCase, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(submitcase.Table, submitcase.Columns, sqlgraph.NewFieldSpec(submitcase.FieldID, field.TypeInt64))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubmitCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submitcase.FieldID)
		for _, f := range fields {
			if !submitcase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != submitcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.State(); ok {
		_spec.SetField(submitcase.FieldState, field.TypeInt16, value)
	}
	if value, ok := scuo.mutation.AddedState(); ok {
		_spec.AddField(submitcase.FieldState, field.TypeInt16, value)
	}
	if value, ok := scuo.mutation.Point(); ok {
		_spec.SetField(submitcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := scuo.mutation.AddedPoint(); ok {
		_spec.AddField(submitcase.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := scuo.mutation.Message(); ok {
		_spec.SetField(submitcase.FieldMessage, field.TypeString, value)
	}
	if value, ok := scuo.mutation.Time(); ok {
		_spec.SetField(submitcase.FieldTime, field.TypeInt32, value)
	}
	if value, ok := scuo.mutation.AddedTime(); ok {
		_spec.AddField(submitcase.FieldTime, field.TypeInt32, value)
	}
	if value, ok := scuo.mutation.Memory(); ok {
		_spec.SetField(submitcase.FieldMemory, field.TypeInt32, value)
	}
	if value, ok := scuo.mutation.AddedMemory(); ok {
		_spec.AddField(submitcase.FieldMemory, field.TypeInt32, value)
	}
	if scuo.mutation.SubmissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.SubmissionTable,
			Columns: []string{submitcase.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.SubmissionTable,
			Columns: []string{submitcase.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scuo.mutation.ProblemCasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.ProblemCasesTable,
			Columns: []string{submitcase.ProblemCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.ProblemCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submitcase.ProblemCasesTable,
			Columns: []string{submitcase.ProblemCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SubmitCase{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submitcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
