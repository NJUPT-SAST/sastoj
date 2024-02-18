// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/contest"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemUpdate is the builder for updating Problem entities.
type ProblemUpdate struct {
	config
	hooks    []Hook
	mutation *ProblemMutation
}

// Where appends a list predicates to the ProblemUpdate builder.
func (pu *ProblemUpdate) Where(ps ...predicate.Problem) *ProblemUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTitle sets the "title" field.
func (pu *ProblemUpdate) SetTitle(s string) *ProblemUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableTitle(s *string) *ProblemUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// SetContent sets the "content" field.
func (pu *ProblemUpdate) SetContent(s string) *ProblemUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableContent(s *string) *ProblemUpdate {
	if s != nil {
		pu.SetContent(*s)
	}
	return pu
}

// SetPoint sets the "point" field.
func (pu *ProblemUpdate) SetPoint(i int) *ProblemUpdate {
	pu.mutation.ResetPoint()
	pu.mutation.SetPoint(i)
	return pu
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillablePoint(i *int) *ProblemUpdate {
	if i != nil {
		pu.SetPoint(*i)
	}
	return pu
}

// AddPoint adds i to the "point" field.
func (pu *ProblemUpdate) AddPoint(i int) *ProblemUpdate {
	pu.mutation.AddPoint(i)
	return pu
}

// SetContestID sets the "contest_id" field.
func (pu *ProblemUpdate) SetContestID(i int) *ProblemUpdate {
	pu.mutation.ResetContestID()
	pu.mutation.SetContestID(i)
	return pu
}

// SetNillableContestID sets the "contest_id" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableContestID(i *int) *ProblemUpdate {
	if i != nil {
		pu.SetContestID(*i)
	}
	return pu
}

// AddContestID adds i to the "contest_id" field.
func (pu *ProblemUpdate) AddContestID(i int) *ProblemUpdate {
	pu.mutation.AddContestID(i)
	return pu
}

// SetCaseVersion sets the "case_version" field.
func (pu *ProblemUpdate) SetCaseVersion(i int) *ProblemUpdate {
	pu.mutation.ResetCaseVersion()
	pu.mutation.SetCaseVersion(i)
	return pu
}

// SetNillableCaseVersion sets the "case_version" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableCaseVersion(i *int) *ProblemUpdate {
	if i != nil {
		pu.SetCaseVersion(*i)
	}
	return pu
}

// AddCaseVersion adds i to the "case_version" field.
func (pu *ProblemUpdate) AddCaseVersion(i int) *ProblemUpdate {
	pu.mutation.AddCaseVersion(i)
	return pu
}

// SetIndex sets the "index" field.
func (pu *ProblemUpdate) SetIndex(i int) *ProblemUpdate {
	pu.mutation.ResetIndex()
	pu.mutation.SetIndex(i)
	return pu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableIndex(i *int) *ProblemUpdate {
	if i != nil {
		pu.SetIndex(*i)
	}
	return pu
}

// AddIndex adds i to the "index" field.
func (pu *ProblemUpdate) AddIndex(i int) *ProblemUpdate {
	pu.mutation.AddIndex(i)
	return pu
}

// SetIsDeleted sets the "is_deleted" field.
func (pu *ProblemUpdate) SetIsDeleted(b bool) *ProblemUpdate {
	pu.mutation.SetIsDeleted(b)
	return pu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableIsDeleted(b *bool) *ProblemUpdate {
	if b != nil {
		pu.SetIsDeleted(*b)
	}
	return pu
}

// SetConfig sets the "config" field.
func (pu *ProblemUpdate) SetConfig(s string) *ProblemUpdate {
	pu.mutation.SetConfig(s)
	return pu
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableConfig(s *string) *ProblemUpdate {
	if s != nil {
		pu.SetConfig(*s)
	}
	return pu
}

// SetContestID sets the "contest" edge to the Contest entity by ID.
func (pu *ProblemUpdate) SetContestID(id int) *ProblemUpdate {
	pu.mutation.SetContestID(id)
	return pu
}

// SetNillableContestID sets the "contest" edge to the Contest entity by ID if the given value is not nil.
func (pu *ProblemUpdate) SetNillableContestID(id *int) *ProblemUpdate {
	if id != nil {
		pu = pu.SetContestID(*id)
	}
	return pu
}

// SetContest sets the "contest" edge to the Contest entity.
func (pu *ProblemUpdate) SetContest(c *Contest) *ProblemUpdate {
	return pu.SetContestID(c.ID)
}

// Mutation returns the ProblemMutation object of the builder.
func (pu *ProblemUpdate) Mutation() *ProblemMutation {
	return pu.mutation
}

// ClearContest clears the "contest" edge to the Contest entity.
func (pu *ProblemUpdate) ClearContest() *ProblemUpdate {
	pu.mutation.ClearContest()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProblemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProblemUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProblemUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProblemUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProblemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(problem.Table, problem.Columns, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(problem.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(problem.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.Point(); ok {
		_spec.SetField(problem.FieldPoint, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedPoint(); ok {
		_spec.AddField(problem.FieldPoint, field.TypeInt, value)
	}
	if value, ok := pu.mutation.ContestID(); ok {
		_spec.SetField(problem.FieldContestID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedContestID(); ok {
		_spec.AddField(problem.FieldContestID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.CaseVersion(); ok {
		_spec.SetField(problem.FieldCaseVersion, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedCaseVersion(); ok {
		_spec.AddField(problem.FieldCaseVersion, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Index(); ok {
		_spec.SetField(problem.FieldIndex, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedIndex(); ok {
		_spec.AddField(problem.FieldIndex, field.TypeInt, value)
	}
	if value, ok := pu.mutation.IsDeleted(); ok {
		_spec.SetField(problem.FieldIsDeleted, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Config(); ok {
		_spec.SetField(problem.FieldConfig, field.TypeString, value)
	}
	if pu.mutation.ContestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problem.ContestTable,
			Columns: []string{problem.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problem.ContestTable,
			Columns: []string{problem.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProblemUpdateOne is the builder for updating a single Problem entity.
type ProblemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProblemMutation
}

// SetTitle sets the "title" field.
func (puo *ProblemUpdateOne) SetTitle(s string) *ProblemUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableTitle(s *string) *ProblemUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// SetContent sets the "content" field.
func (puo *ProblemUpdateOne) SetContent(s string) *ProblemUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableContent(s *string) *ProblemUpdateOne {
	if s != nil {
		puo.SetContent(*s)
	}
	return puo
}

// SetPoint sets the "point" field.
func (puo *ProblemUpdateOne) SetPoint(i int) *ProblemUpdateOne {
	puo.mutation.ResetPoint()
	puo.mutation.SetPoint(i)
	return puo
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillablePoint(i *int) *ProblemUpdateOne {
	if i != nil {
		puo.SetPoint(*i)
	}
	return puo
}

// AddPoint adds i to the "point" field.
func (puo *ProblemUpdateOne) AddPoint(i int) *ProblemUpdateOne {
	puo.mutation.AddPoint(i)
	return puo
}

// SetContestID sets the "contest_id" field.
func (puo *ProblemUpdateOne) SetContestID(i int) *ProblemUpdateOne {
	puo.mutation.ResetContestID()
	puo.mutation.SetContestID(i)
	return puo
}

// SetNillableContestID sets the "contest_id" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableContestID(i *int) *ProblemUpdateOne {
	if i != nil {
		puo.SetContestID(*i)
	}
	return puo
}

// AddContestID adds i to the "contest_id" field.
func (puo *ProblemUpdateOne) AddContestID(i int) *ProblemUpdateOne {
	puo.mutation.AddContestID(i)
	return puo
}

// SetCaseVersion sets the "case_version" field.
func (puo *ProblemUpdateOne) SetCaseVersion(i int) *ProblemUpdateOne {
	puo.mutation.ResetCaseVersion()
	puo.mutation.SetCaseVersion(i)
	return puo
}

// SetNillableCaseVersion sets the "case_version" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableCaseVersion(i *int) *ProblemUpdateOne {
	if i != nil {
		puo.SetCaseVersion(*i)
	}
	return puo
}

// AddCaseVersion adds i to the "case_version" field.
func (puo *ProblemUpdateOne) AddCaseVersion(i int) *ProblemUpdateOne {
	puo.mutation.AddCaseVersion(i)
	return puo
}

// SetIndex sets the "index" field.
func (puo *ProblemUpdateOne) SetIndex(i int) *ProblemUpdateOne {
	puo.mutation.ResetIndex()
	puo.mutation.SetIndex(i)
	return puo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableIndex(i *int) *ProblemUpdateOne {
	if i != nil {
		puo.SetIndex(*i)
	}
	return puo
}

// AddIndex adds i to the "index" field.
func (puo *ProblemUpdateOne) AddIndex(i int) *ProblemUpdateOne {
	puo.mutation.AddIndex(i)
	return puo
}

// SetIsDeleted sets the "is_deleted" field.
func (puo *ProblemUpdateOne) SetIsDeleted(b bool) *ProblemUpdateOne {
	puo.mutation.SetIsDeleted(b)
	return puo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableIsDeleted(b *bool) *ProblemUpdateOne {
	if b != nil {
		puo.SetIsDeleted(*b)
	}
	return puo
}

// SetConfig sets the "config" field.
func (puo *ProblemUpdateOne) SetConfig(s string) *ProblemUpdateOne {
	puo.mutation.SetConfig(s)
	return puo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableConfig(s *string) *ProblemUpdateOne {
	if s != nil {
		puo.SetConfig(*s)
	}
	return puo
}

// SetContestID sets the "contest" edge to the Contest entity by ID.
func (puo *ProblemUpdateOne) SetContestID(id int) *ProblemUpdateOne {
	puo.mutation.SetContestID(id)
	return puo
}

// SetNillableContestID sets the "contest" edge to the Contest entity by ID if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableContestID(id *int) *ProblemUpdateOne {
	if id != nil {
		puo = puo.SetContestID(*id)
	}
	return puo
}

// SetContest sets the "contest" edge to the Contest entity.
func (puo *ProblemUpdateOne) SetContest(c *Contest) *ProblemUpdateOne {
	return puo.SetContestID(c.ID)
}

// Mutation returns the ProblemMutation object of the builder.
func (puo *ProblemUpdateOne) Mutation() *ProblemMutation {
	return puo.mutation
}

// ClearContest clears the "contest" edge to the Contest entity.
func (puo *ProblemUpdateOne) ClearContest() *ProblemUpdateOne {
	puo.mutation.ClearContest()
	return puo
}

// Where appends a list predicates to the ProblemUpdate builder.
func (puo *ProblemUpdateOne) Where(ps ...predicate.Problem) *ProblemUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProblemUpdateOne) Select(field string, fields ...string) *ProblemUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Problem entity.
func (puo *ProblemUpdateOne) Save(ctx context.Context) (*Problem, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProblemUpdateOne) SaveX(ctx context.Context) *Problem {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProblemUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProblemUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProblemUpdateOne) sqlSave(ctx context.Context) (_node *Problem, err error) {
	_spec := sqlgraph.NewUpdateSpec(problem.Table, problem.Columns, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Problem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problem.FieldID)
		for _, f := range fields {
			if !problem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != problem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(problem.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(problem.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.Point(); ok {
		_spec.SetField(problem.FieldPoint, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedPoint(); ok {
		_spec.AddField(problem.FieldPoint, field.TypeInt, value)
	}
	if value, ok := puo.mutation.ContestID(); ok {
		_spec.SetField(problem.FieldContestID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedContestID(); ok {
		_spec.AddField(problem.FieldContestID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.CaseVersion(); ok {
		_spec.SetField(problem.FieldCaseVersion, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedCaseVersion(); ok {
		_spec.AddField(problem.FieldCaseVersion, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Index(); ok {
		_spec.SetField(problem.FieldIndex, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedIndex(); ok {
		_spec.AddField(problem.FieldIndex, field.TypeInt, value)
	}
	if value, ok := puo.mutation.IsDeleted(); ok {
		_spec.SetField(problem.FieldIsDeleted, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Config(); ok {
		_spec.SetField(problem.FieldConfig, field.TypeString, value)
	}
	if puo.mutation.ContestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problem.ContestTable,
			Columns: []string{problem.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   problem.ContestTable,
			Columns: []string{problem.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Problem{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
