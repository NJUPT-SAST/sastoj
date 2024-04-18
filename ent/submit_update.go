// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"
	"sastoj/ent/submit"
	"sastoj/ent/submitcase"
	"sastoj/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmitUpdate is the builder for updating Submit entities.
type SubmitUpdate struct {
	config
	hooks    []Hook
	mutation *SubmitMutation
}

// Where appends a list predicates to the SubmitUpdate builder.
func (su *SubmitUpdate) Where(ps ...predicate.Submit) *SubmitUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCode sets the "code" field.
func (su *SubmitUpdate) SetCode(s string) *SubmitUpdate {
	su.mutation.SetCode(s)
	return su
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableCode(s *string) *SubmitUpdate {
	if s != nil {
		su.SetCode(*s)
	}
	return su
}

// SetStatus sets the "status" field.
func (su *SubmitUpdate) SetStatus(i int8) *SubmitUpdate {
	su.mutation.ResetStatus()
	su.mutation.SetStatus(i)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableStatus(i *int8) *SubmitUpdate {
	if i != nil {
		su.SetStatus(*i)
	}
	return su
}

// AddStatus adds i to the "status" field.
func (su *SubmitUpdate) AddStatus(i int8) *SubmitUpdate {
	su.mutation.AddStatus(i)
	return su
}

// SetPoint sets the "point" field.
func (su *SubmitUpdate) SetPoint(i int16) *SubmitUpdate {
	su.mutation.ResetPoint()
	su.mutation.SetPoint(i)
	return su
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (su *SubmitUpdate) SetNillablePoint(i *int16) *SubmitUpdate {
	if i != nil {
		su.SetPoint(*i)
	}
	return su
}

// AddPoint adds i to the "point" field.
func (su *SubmitUpdate) AddPoint(i int16) *SubmitUpdate {
	su.mutation.AddPoint(i)
	return su
}

// SetCreateTime sets the "create_time" field.
func (su *SubmitUpdate) SetCreateTime(t time.Time) *SubmitUpdate {
	su.mutation.SetCreateTime(t)
	return su
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableCreateTime(t *time.Time) *SubmitUpdate {
	if t != nil {
		su.SetCreateTime(*t)
	}
	return su
}

// SetTotalTime sets the "total_time" field.
func (su *SubmitUpdate) SetTotalTime(i int) *SubmitUpdate {
	su.mutation.ResetTotalTime()
	su.mutation.SetTotalTime(i)
	return su
}

// SetNillableTotalTime sets the "total_time" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableTotalTime(i *int) *SubmitUpdate {
	if i != nil {
		su.SetTotalTime(*i)
	}
	return su
}

// AddTotalTime adds i to the "total_time" field.
func (su *SubmitUpdate) AddTotalTime(i int) *SubmitUpdate {
	su.mutation.AddTotalTime(i)
	return su
}

// SetMaxMemory sets the "max_memory" field.
func (su *SubmitUpdate) SetMaxMemory(i int) *SubmitUpdate {
	su.mutation.ResetMaxMemory()
	su.mutation.SetMaxMemory(i)
	return su
}

// SetNillableMaxMemory sets the "max_memory" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableMaxMemory(i *int) *SubmitUpdate {
	if i != nil {
		su.SetMaxMemory(*i)
	}
	return su
}

// AddMaxMemory adds i to the "max_memory" field.
func (su *SubmitUpdate) AddMaxMemory(i int) *SubmitUpdate {
	su.mutation.AddMaxMemory(i)
	return su
}

// SetLanguage sets the "language" field.
func (su *SubmitUpdate) SetLanguage(s string) *SubmitUpdate {
	su.mutation.SetLanguage(s)
	return su
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableLanguage(s *string) *SubmitUpdate {
	if s != nil {
		su.SetLanguage(*s)
	}
	return su
}

// SetCaseVersion sets the "case_version" field.
func (su *SubmitUpdate) SetCaseVersion(i int8) *SubmitUpdate {
	su.mutation.ResetCaseVersion()
	su.mutation.SetCaseVersion(i)
	return su
}

// SetNillableCaseVersion sets the "case_version" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableCaseVersion(i *int8) *SubmitUpdate {
	if i != nil {
		su.SetCaseVersion(*i)
	}
	return su
}

// AddCaseVersion adds i to the "case_version" field.
func (su *SubmitUpdate) AddCaseVersion(i int8) *SubmitUpdate {
	su.mutation.AddCaseVersion(i)
	return su
}

// SetProblemID sets the "problem_id" field.
func (su *SubmitUpdate) SetProblemID(i int) *SubmitUpdate {
	su.mutation.SetProblemID(i)
	return su
}

// SetNillableProblemID sets the "problem_id" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableProblemID(i *int) *SubmitUpdate {
	if i != nil {
		su.SetProblemID(*i)
	}
	return su
}

// SetUserID sets the "user_id" field.
func (su *SubmitUpdate) SetUserID(i int) *SubmitUpdate {
	su.mutation.SetUserID(i)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SubmitUpdate) SetNillableUserID(i *int) *SubmitUpdate {
	if i != nil {
		su.SetUserID(*i)
	}
	return su
}

// AddSubmitCaseIDs adds the "submit_cases" edge to the SubmitCase entity by IDs.
func (su *SubmitUpdate) AddSubmitCaseIDs(ids ...int) *SubmitUpdate {
	su.mutation.AddSubmitCaseIDs(ids...)
	return su
}

// AddSubmitCases adds the "submit_cases" edges to the SubmitCase entity.
func (su *SubmitUpdate) AddSubmitCases(s ...*SubmitCase) *SubmitUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSubmitCaseIDs(ids...)
}

// SetProblemsID sets the "problems" edge to the Problem entity by ID.
func (su *SubmitUpdate) SetProblemsID(id int) *SubmitUpdate {
	su.mutation.SetProblemsID(id)
	return su
}

// SetProblems sets the "problems" edge to the Problem entity.
func (su *SubmitUpdate) SetProblems(p *Problem) *SubmitUpdate {
	return su.SetProblemsID(p.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (su *SubmitUpdate) SetUsersID(id int) *SubmitUpdate {
	su.mutation.SetUsersID(id)
	return su
}

// SetUsers sets the "users" edge to the User entity.
func (su *SubmitUpdate) SetUsers(u *User) *SubmitUpdate {
	return su.SetUsersID(u.ID)
}

// Mutation returns the SubmitMutation object of the builder.
func (su *SubmitUpdate) Mutation() *SubmitMutation {
	return su.mutation
}

// ClearSubmitCases clears all "submit_cases" edges to the SubmitCase entity.
func (su *SubmitUpdate) ClearSubmitCases() *SubmitUpdate {
	su.mutation.ClearSubmitCases()
	return su
}

// RemoveSubmitCaseIDs removes the "submit_cases" edge to SubmitCase entities by IDs.
func (su *SubmitUpdate) RemoveSubmitCaseIDs(ids ...int) *SubmitUpdate {
	su.mutation.RemoveSubmitCaseIDs(ids...)
	return su
}

// RemoveSubmitCases removes "submit_cases" edges to SubmitCase entities.
func (su *SubmitUpdate) RemoveSubmitCases(s ...*SubmitCase) *SubmitUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSubmitCaseIDs(ids...)
}

// ClearProblems clears the "problems" edge to the Problem entity.
func (su *SubmitUpdate) ClearProblems() *SubmitUpdate {
	su.mutation.ClearProblems()
	return su
}

// ClearUsers clears the "users" edge to the User entity.
func (su *SubmitUpdate) ClearUsers() *SubmitUpdate {
	su.mutation.ClearUsers()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubmitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubmitUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubmitUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubmitUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubmitUpdate) check() error {
	if v, ok := su.mutation.Status(); ok {
		if err := submit.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Submit.status": %w`, err)}
		}
	}
	if v, ok := su.mutation.Point(); ok {
		if err := submit.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "Submit.point": %w`, err)}
		}
	}
	if v, ok := su.mutation.TotalTime(); ok {
		if err := submit.TotalTimeValidator(v); err != nil {
			return &ValidationError{Name: "total_time", err: fmt.Errorf(`ent: validator failed for field "Submit.total_time": %w`, err)}
		}
	}
	if v, ok := su.mutation.MaxMemory(); ok {
		if err := submit.MaxMemoryValidator(v); err != nil {
			return &ValidationError{Name: "max_memory", err: fmt.Errorf(`ent: validator failed for field "Submit.max_memory": %w`, err)}
		}
	}
	if v, ok := su.mutation.CaseVersion(); ok {
		if err := submit.CaseVersionValidator(v); err != nil {
			return &ValidationError{Name: "case_version", err: fmt.Errorf(`ent: validator failed for field "Submit.case_version": %w`, err)}
		}
	}
	if _, ok := su.mutation.ProblemsID(); su.mutation.ProblemsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Submit.problems"`)
	}
	if _, ok := su.mutation.UsersID(); su.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Submit.users"`)
	}
	return nil
}

func (su *SubmitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Code(); ok {
		_spec.SetField(submit.FieldCode, field.TypeString, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(submit.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AddedStatus(); ok {
		_spec.AddField(submit.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := su.mutation.Point(); ok {
		_spec.SetField(submit.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := su.mutation.AddedPoint(); ok {
		_spec.AddField(submit.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := su.mutation.CreateTime(); ok {
		_spec.SetField(submit.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.TotalTime(); ok {
		_spec.SetField(submit.FieldTotalTime, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedTotalTime(); ok {
		_spec.AddField(submit.FieldTotalTime, field.TypeInt, value)
	}
	if value, ok := su.mutation.MaxMemory(); ok {
		_spec.SetField(submit.FieldMaxMemory, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedMaxMemory(); ok {
		_spec.AddField(submit.FieldMaxMemory, field.TypeInt, value)
	}
	if value, ok := su.mutation.Language(); ok {
		_spec.SetField(submit.FieldLanguage, field.TypeString, value)
	}
	if value, ok := su.mutation.CaseVersion(); ok {
		_spec.SetField(submit.FieldCaseVersion, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AddedCaseVersion(); ok {
		_spec.AddField(submit.FieldCaseVersion, field.TypeInt8, value)
	}
	if su.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSubmitCasesIDs(); len(nodes) > 0 && !su.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SubmitCasesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ProblemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProblemsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubmitUpdateOne is the builder for updating a single Submit entity.
type SubmitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubmitMutation
}

// SetCode sets the "code" field.
func (suo *SubmitUpdateOne) SetCode(s string) *SubmitUpdateOne {
	suo.mutation.SetCode(s)
	return suo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableCode(s *string) *SubmitUpdateOne {
	if s != nil {
		suo.SetCode(*s)
	}
	return suo
}

// SetStatus sets the "status" field.
func (suo *SubmitUpdateOne) SetStatus(i int8) *SubmitUpdateOne {
	suo.mutation.ResetStatus()
	suo.mutation.SetStatus(i)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableStatus(i *int8) *SubmitUpdateOne {
	if i != nil {
		suo.SetStatus(*i)
	}
	return suo
}

// AddStatus adds i to the "status" field.
func (suo *SubmitUpdateOne) AddStatus(i int8) *SubmitUpdateOne {
	suo.mutation.AddStatus(i)
	return suo
}

// SetPoint sets the "point" field.
func (suo *SubmitUpdateOne) SetPoint(i int16) *SubmitUpdateOne {
	suo.mutation.ResetPoint()
	suo.mutation.SetPoint(i)
	return suo
}

// SetNillablePoint sets the "point" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillablePoint(i *int16) *SubmitUpdateOne {
	if i != nil {
		suo.SetPoint(*i)
	}
	return suo
}

// AddPoint adds i to the "point" field.
func (suo *SubmitUpdateOne) AddPoint(i int16) *SubmitUpdateOne {
	suo.mutation.AddPoint(i)
	return suo
}

// SetCreateTime sets the "create_time" field.
func (suo *SubmitUpdateOne) SetCreateTime(t time.Time) *SubmitUpdateOne {
	suo.mutation.SetCreateTime(t)
	return suo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableCreateTime(t *time.Time) *SubmitUpdateOne {
	if t != nil {
		suo.SetCreateTime(*t)
	}
	return suo
}

// SetTotalTime sets the "total_time" field.
func (suo *SubmitUpdateOne) SetTotalTime(i int) *SubmitUpdateOne {
	suo.mutation.ResetTotalTime()
	suo.mutation.SetTotalTime(i)
	return suo
}

// SetNillableTotalTime sets the "total_time" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableTotalTime(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetTotalTime(*i)
	}
	return suo
}

// AddTotalTime adds i to the "total_time" field.
func (suo *SubmitUpdateOne) AddTotalTime(i int) *SubmitUpdateOne {
	suo.mutation.AddTotalTime(i)
	return suo
}

// SetMaxMemory sets the "max_memory" field.
func (suo *SubmitUpdateOne) SetMaxMemory(i int) *SubmitUpdateOne {
	suo.mutation.ResetMaxMemory()
	suo.mutation.SetMaxMemory(i)
	return suo
}

// SetNillableMaxMemory sets the "max_memory" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableMaxMemory(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetMaxMemory(*i)
	}
	return suo
}

// AddMaxMemory adds i to the "max_memory" field.
func (suo *SubmitUpdateOne) AddMaxMemory(i int) *SubmitUpdateOne {
	suo.mutation.AddMaxMemory(i)
	return suo
}

// SetLanguage sets the "language" field.
func (suo *SubmitUpdateOne) SetLanguage(s string) *SubmitUpdateOne {
	suo.mutation.SetLanguage(s)
	return suo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableLanguage(s *string) *SubmitUpdateOne {
	if s != nil {
		suo.SetLanguage(*s)
	}
	return suo
}

// SetCaseVersion sets the "case_version" field.
func (suo *SubmitUpdateOne) SetCaseVersion(i int8) *SubmitUpdateOne {
	suo.mutation.ResetCaseVersion()
	suo.mutation.SetCaseVersion(i)
	return suo
}

// SetNillableCaseVersion sets the "case_version" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableCaseVersion(i *int8) *SubmitUpdateOne {
	if i != nil {
		suo.SetCaseVersion(*i)
	}
	return suo
}

// AddCaseVersion adds i to the "case_version" field.
func (suo *SubmitUpdateOne) AddCaseVersion(i int8) *SubmitUpdateOne {
	suo.mutation.AddCaseVersion(i)
	return suo
}

// SetProblemID sets the "problem_id" field.
func (suo *SubmitUpdateOne) SetProblemID(i int) *SubmitUpdateOne {
	suo.mutation.SetProblemID(i)
	return suo
}

// SetNillableProblemID sets the "problem_id" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableProblemID(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetProblemID(*i)
	}
	return suo
}

// SetUserID sets the "user_id" field.
func (suo *SubmitUpdateOne) SetUserID(i int) *SubmitUpdateOne {
	suo.mutation.SetUserID(i)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SubmitUpdateOne) SetNillableUserID(i *int) *SubmitUpdateOne {
	if i != nil {
		suo.SetUserID(*i)
	}
	return suo
}

// AddSubmitCaseIDs adds the "submit_cases" edge to the SubmitCase entity by IDs.
func (suo *SubmitUpdateOne) AddSubmitCaseIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.AddSubmitCaseIDs(ids...)
	return suo
}

// AddSubmitCases adds the "submit_cases" edges to the SubmitCase entity.
func (suo *SubmitUpdateOne) AddSubmitCases(s ...*SubmitCase) *SubmitUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSubmitCaseIDs(ids...)
}

// SetProblemsID sets the "problems" edge to the Problem entity by ID.
func (suo *SubmitUpdateOne) SetProblemsID(id int) *SubmitUpdateOne {
	suo.mutation.SetProblemsID(id)
	return suo
}

// SetProblems sets the "problems" edge to the Problem entity.
func (suo *SubmitUpdateOne) SetProblems(p *Problem) *SubmitUpdateOne {
	return suo.SetProblemsID(p.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (suo *SubmitUpdateOne) SetUsersID(id int) *SubmitUpdateOne {
	suo.mutation.SetUsersID(id)
	return suo
}

// SetUsers sets the "users" edge to the User entity.
func (suo *SubmitUpdateOne) SetUsers(u *User) *SubmitUpdateOne {
	return suo.SetUsersID(u.ID)
}

// Mutation returns the SubmitMutation object of the builder.
func (suo *SubmitUpdateOne) Mutation() *SubmitMutation {
	return suo.mutation
}

// ClearSubmitCases clears all "submit_cases" edges to the SubmitCase entity.
func (suo *SubmitUpdateOne) ClearSubmitCases() *SubmitUpdateOne {
	suo.mutation.ClearSubmitCases()
	return suo
}

// RemoveSubmitCaseIDs removes the "submit_cases" edge to SubmitCase entities by IDs.
func (suo *SubmitUpdateOne) RemoveSubmitCaseIDs(ids ...int) *SubmitUpdateOne {
	suo.mutation.RemoveSubmitCaseIDs(ids...)
	return suo
}

// RemoveSubmitCases removes "submit_cases" edges to SubmitCase entities.
func (suo *SubmitUpdateOne) RemoveSubmitCases(s ...*SubmitCase) *SubmitUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSubmitCaseIDs(ids...)
}

// ClearProblems clears the "problems" edge to the Problem entity.
func (suo *SubmitUpdateOne) ClearProblems() *SubmitUpdateOne {
	suo.mutation.ClearProblems()
	return suo
}

// ClearUsers clears the "users" edge to the User entity.
func (suo *SubmitUpdateOne) ClearUsers() *SubmitUpdateOne {
	suo.mutation.ClearUsers()
	return suo
}

// Where appends a list predicates to the SubmitUpdate builder.
func (suo *SubmitUpdateOne) Where(ps ...predicate.Submit) *SubmitUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubmitUpdateOne) Select(field string, fields ...string) *SubmitUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Submit entity.
func (suo *SubmitUpdateOne) Save(ctx context.Context) (*Submit, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubmitUpdateOne) SaveX(ctx context.Context) *Submit {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubmitUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubmitUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubmitUpdateOne) check() error {
	if v, ok := suo.mutation.Status(); ok {
		if err := submit.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Submit.status": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Point(); ok {
		if err := submit.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "Submit.point": %w`, err)}
		}
	}
	if v, ok := suo.mutation.TotalTime(); ok {
		if err := submit.TotalTimeValidator(v); err != nil {
			return &ValidationError{Name: "total_time", err: fmt.Errorf(`ent: validator failed for field "Submit.total_time": %w`, err)}
		}
	}
	if v, ok := suo.mutation.MaxMemory(); ok {
		if err := submit.MaxMemoryValidator(v); err != nil {
			return &ValidationError{Name: "max_memory", err: fmt.Errorf(`ent: validator failed for field "Submit.max_memory": %w`, err)}
		}
	}
	if v, ok := suo.mutation.CaseVersion(); ok {
		if err := submit.CaseVersionValidator(v); err != nil {
			return &ValidationError{Name: "case_version", err: fmt.Errorf(`ent: validator failed for field "Submit.case_version": %w`, err)}
		}
	}
	if _, ok := suo.mutation.ProblemsID(); suo.mutation.ProblemsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Submit.problems"`)
	}
	if _, ok := suo.mutation.UsersID(); suo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Submit.users"`)
	}
	return nil
}

func (suo *SubmitUpdateOne) sqlSave(ctx context.Context) (_node *Submit, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(submit.Table, submit.Columns, sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Submit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, submit.FieldID)
		for _, f := range fields {
			if !submit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != submit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Code(); ok {
		_spec.SetField(submit.FieldCode, field.TypeString, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(submit.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AddedStatus(); ok {
		_spec.AddField(submit.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.Point(); ok {
		_spec.SetField(submit.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := suo.mutation.AddedPoint(); ok {
		_spec.AddField(submit.FieldPoint, field.TypeInt16, value)
	}
	if value, ok := suo.mutation.CreateTime(); ok {
		_spec.SetField(submit.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.TotalTime(); ok {
		_spec.SetField(submit.FieldTotalTime, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedTotalTime(); ok {
		_spec.AddField(submit.FieldTotalTime, field.TypeInt, value)
	}
	if value, ok := suo.mutation.MaxMemory(); ok {
		_spec.SetField(submit.FieldMaxMemory, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedMaxMemory(); ok {
		_spec.AddField(submit.FieldMaxMemory, field.TypeInt, value)
	}
	if value, ok := suo.mutation.Language(); ok {
		_spec.SetField(submit.FieldLanguage, field.TypeString, value)
	}
	if value, ok := suo.mutation.CaseVersion(); ok {
		_spec.SetField(submit.FieldCaseVersion, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AddedCaseVersion(); ok {
		_spec.AddField(submit.FieldCaseVersion, field.TypeInt8, value)
	}
	if suo.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSubmitCasesIDs(); len(nodes) > 0 && !suo.mutation.SubmitCasesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SubmitCasesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ProblemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProblemsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Submit{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{submit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
