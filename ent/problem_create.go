// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"sastoj/ent/submission"
	"sastoj/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemCreate is the builder for creating a Problem entity.
type ProblemCreate struct {
	config
	mutation *ProblemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (pc *ProblemCreate) SetTitle(s string) *ProblemCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *ProblemCreate) SetContent(s string) *ProblemCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetPoint sets the "point" field.
func (pc *ProblemCreate) SetPoint(i int16) *ProblemCreate {
	pc.mutation.SetPoint(i)
	return pc
}

// SetCaseVersion sets the "case_version" field.
func (pc *ProblemCreate) SetCaseVersion(i int16) *ProblemCreate {
	pc.mutation.SetCaseVersion(i)
	return pc
}

// SetNillableCaseVersion sets the "case_version" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableCaseVersion(i *int16) *ProblemCreate {
	if i != nil {
		pc.SetCaseVersion(*i)
	}
	return pc
}

// SetIndex sets the "index" field.
func (pc *ProblemCreate) SetIndex(i int16) *ProblemCreate {
	pc.mutation.SetIndex(i)
	return pc
}

// SetRestrictPresentation sets the "restrict_presentation" field.
func (pc *ProblemCreate) SetRestrictPresentation(b bool) *ProblemCreate {
	pc.mutation.SetRestrictPresentation(b)
	return pc
}

// SetNillableRestrictPresentation sets the "restrict_presentation" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableRestrictPresentation(b *bool) *ProblemCreate {
	if b != nil {
		pc.SetRestrictPresentation(*b)
	}
	return pc
}

// SetIsDeleted sets the "is_deleted" field.
func (pc *ProblemCreate) SetIsDeleted(b bool) *ProblemCreate {
	pc.mutation.SetIsDeleted(b)
	return pc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableIsDeleted(b *bool) *ProblemCreate {
	if b != nil {
		pc.SetIsDeleted(*b)
	}
	return pc
}

// SetConfig sets the "config" field.
func (pc *ProblemCreate) SetConfig(s string) *ProblemCreate {
	pc.mutation.SetConfig(s)
	return pc
}

// SetContestID sets the "contest_id" field.
func (pc *ProblemCreate) SetContestID(i int64) *ProblemCreate {
	pc.mutation.SetContestID(i)
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *ProblemCreate) SetUserID(i int64) *ProblemCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetVisibility sets the "visibility" field.
func (pc *ProblemCreate) SetVisibility(i int8) *ProblemCreate {
	pc.mutation.SetVisibility(i)
	return pc
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableVisibility(i *int8) *ProblemCreate {
	if i != nil {
		pc.SetVisibility(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProblemCreate) SetID(i int64) *ProblemCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddProblemCaseIDs adds the "problem_cases" edge to the ProblemCase entity by IDs.
func (pc *ProblemCreate) AddProblemCaseIDs(ids ...int64) *ProblemCreate {
	pc.mutation.AddProblemCaseIDs(ids...)
	return pc
}

// AddProblemCases adds the "problem_cases" edges to the ProblemCase entity.
func (pc *ProblemCreate) AddProblemCases(p ...*ProblemCase) *ProblemCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProblemCaseIDs(ids...)
}

// AddSubmissionIDs adds the "submission" edge to the Submission entity by IDs.
func (pc *ProblemCreate) AddSubmissionIDs(ids ...int64) *ProblemCreate {
	pc.mutation.AddSubmissionIDs(ids...)
	return pc
}

// AddSubmission adds the "submission" edges to the Submission entity.
func (pc *ProblemCreate) AddSubmission(s ...*Submission) *ProblemCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSubmissionIDs(ids...)
}

// SetContestsID sets the "contests" edge to the Contest entity by ID.
func (pc *ProblemCreate) SetContestsID(id int64) *ProblemCreate {
	pc.mutation.SetContestsID(id)
	return pc
}

// SetContests sets the "contests" edge to the Contest entity.
func (pc *ProblemCreate) SetContests(c *Contest) *ProblemCreate {
	return pc.SetContestsID(c.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *ProblemCreate) SetOwnerID(id int64) *ProblemCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *ProblemCreate) SetOwner(u *User) *ProblemCreate {
	return pc.SetOwnerID(u.ID)
}

// AddJudgerIDs adds the "judgers" edge to the Group entity by IDs.
func (pc *ProblemCreate) AddJudgerIDs(ids ...int64) *ProblemCreate {
	pc.mutation.AddJudgerIDs(ids...)
	return pc
}

// AddJudgers adds the "judgers" edges to the Group entity.
func (pc *ProblemCreate) AddJudgers(g ...*Group) *ProblemCreate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pc.AddJudgerIDs(ids...)
}

// Mutation returns the ProblemMutation object of the builder.
func (pc *ProblemCreate) Mutation() *ProblemMutation {
	return pc.mutation
}

// Save creates the Problem in the database.
func (pc *ProblemCreate) Save(ctx context.Context) (*Problem, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProblemCreate) SaveX(ctx context.Context) *Problem {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProblemCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProblemCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProblemCreate) defaults() {
	if _, ok := pc.mutation.CaseVersion(); !ok {
		v := problem.DefaultCaseVersion
		pc.mutation.SetCaseVersion(v)
	}
	if _, ok := pc.mutation.RestrictPresentation(); !ok {
		v := problem.DefaultRestrictPresentation
		pc.mutation.SetRestrictPresentation(v)
	}
	if _, ok := pc.mutation.IsDeleted(); !ok {
		v := problem.DefaultIsDeleted
		pc.mutation.SetIsDeleted(v)
	}
	if _, ok := pc.mutation.Visibility(); !ok {
		v := problem.DefaultVisibility
		pc.mutation.SetVisibility(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProblemCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Problem.title"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Problem.content"`)}
	}
	if _, ok := pc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "Problem.point"`)}
	}
	if v, ok := pc.mutation.Point(); ok {
		if err := problem.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "Problem.point": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CaseVersion(); !ok {
		return &ValidationError{Name: "case_version", err: errors.New(`ent: missing required field "Problem.case_version"`)}
	}
	if _, ok := pc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "Problem.index"`)}
	}
	if v, ok := pc.mutation.Index(); ok {
		if err := problem.IndexValidator(v); err != nil {
			return &ValidationError{Name: "index", err: fmt.Errorf(`ent: validator failed for field "Problem.index": %w`, err)}
		}
	}
	if _, ok := pc.mutation.RestrictPresentation(); !ok {
		return &ValidationError{Name: "restrict_presentation", err: errors.New(`ent: missing required field "Problem.restrict_presentation"`)}
	}
	if _, ok := pc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Problem.is_deleted"`)}
	}
	if _, ok := pc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "Problem.config"`)}
	}
	if _, ok := pc.mutation.ContestID(); !ok {
		return &ValidationError{Name: "contest_id", err: errors.New(`ent: missing required field "Problem.contest_id"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Problem.user_id"`)}
	}
	if _, ok := pc.mutation.Visibility(); !ok {
		return &ValidationError{Name: "visibility", err: errors.New(`ent: missing required field "Problem.visibility"`)}
	}
	if _, ok := pc.mutation.ContestsID(); !ok {
		return &ValidationError{Name: "contests", err: errors.New(`ent: missing required edge "Problem.contests"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Problem.owner"`)}
	}
	return nil
}

func (pc *ProblemCreate) sqlSave(ctx context.Context) (*Problem, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProblemCreate) createSpec() (*Problem, *sqlgraph.CreateSpec) {
	var (
		_node = &Problem{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(problem.Table, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(problem.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(problem.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.Point(); ok {
		_spec.SetField(problem.FieldPoint, field.TypeInt16, value)
		_node.Point = value
	}
	if value, ok := pc.mutation.CaseVersion(); ok {
		_spec.SetField(problem.FieldCaseVersion, field.TypeInt16, value)
		_node.CaseVersion = value
	}
	if value, ok := pc.mutation.Index(); ok {
		_spec.SetField(problem.FieldIndex, field.TypeInt16, value)
		_node.Index = value
	}
	if value, ok := pc.mutation.RestrictPresentation(); ok {
		_spec.SetField(problem.FieldRestrictPresentation, field.TypeBool, value)
		_node.RestrictPresentation = value
	}
	if value, ok := pc.mutation.IsDeleted(); ok {
		_spec.SetField(problem.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = value
	}
	if value, ok := pc.mutation.Config(); ok {
		_spec.SetField(problem.FieldConfig, field.TypeString, value)
		_node.Config = value
	}
	if value, ok := pc.mutation.Visibility(); ok {
		_spec.SetField(problem.FieldVisibility, field.TypeInt8, value)
		_node.Visibility = value
	}
	if nodes := pc.mutation.ProblemCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   problem.ProblemCasesTable,
			Columns: []string{problem.ProblemCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   problem.SubmissionTable,
			Columns: []string{problem.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ContestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   problem.ContestsTable,
			Columns: []string{problem.ContestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ContestID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   problem.OwnerTable,
			Columns: []string{problem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.JudgersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   problem.JudgersTable,
			Columns: problem.JudgersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Problem.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProblemUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (pc *ProblemCreate) OnConflict(opts ...sql.ConflictOption) *ProblemUpsertOne {
	pc.conflict = opts
	return &ProblemUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Problem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProblemCreate) OnConflictColumns(columns ...string) *ProblemUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProblemUpsertOne{
		create: pc,
	}
}

type (
	// ProblemUpsertOne is the builder for "upsert"-ing
	//  one Problem node.
	ProblemUpsertOne struct {
		create *ProblemCreate
	}

	// ProblemUpsert is the "OnConflict" setter.
	ProblemUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *ProblemUpsert) SetTitle(v string) *ProblemUpsert {
	u.Set(problem.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateTitle() *ProblemUpsert {
	u.SetExcluded(problem.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *ProblemUpsert) SetContent(v string) *ProblemUpsert {
	u.Set(problem.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateContent() *ProblemUpsert {
	u.SetExcluded(problem.FieldContent)
	return u
}

// SetPoint sets the "point" field.
func (u *ProblemUpsert) SetPoint(v int16) *ProblemUpsert {
	u.Set(problem.FieldPoint, v)
	return u
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *ProblemUpsert) UpdatePoint() *ProblemUpsert {
	u.SetExcluded(problem.FieldPoint)
	return u
}

// AddPoint adds v to the "point" field.
func (u *ProblemUpsert) AddPoint(v int16) *ProblemUpsert {
	u.Add(problem.FieldPoint, v)
	return u
}

// SetCaseVersion sets the "case_version" field.
func (u *ProblemUpsert) SetCaseVersion(v int16) *ProblemUpsert {
	u.Set(problem.FieldCaseVersion, v)
	return u
}

// UpdateCaseVersion sets the "case_version" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateCaseVersion() *ProblemUpsert {
	u.SetExcluded(problem.FieldCaseVersion)
	return u
}

// AddCaseVersion adds v to the "case_version" field.
func (u *ProblemUpsert) AddCaseVersion(v int16) *ProblemUpsert {
	u.Add(problem.FieldCaseVersion, v)
	return u
}

// SetIndex sets the "index" field.
func (u *ProblemUpsert) SetIndex(v int16) *ProblemUpsert {
	u.Set(problem.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateIndex() *ProblemUpsert {
	u.SetExcluded(problem.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *ProblemUpsert) AddIndex(v int16) *ProblemUpsert {
	u.Add(problem.FieldIndex, v)
	return u
}

// SetRestrictPresentation sets the "restrict_presentation" field.
func (u *ProblemUpsert) SetRestrictPresentation(v bool) *ProblemUpsert {
	u.Set(problem.FieldRestrictPresentation, v)
	return u
}

// UpdateRestrictPresentation sets the "restrict_presentation" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateRestrictPresentation() *ProblemUpsert {
	u.SetExcluded(problem.FieldRestrictPresentation)
	return u
}

// SetIsDeleted sets the "is_deleted" field.
func (u *ProblemUpsert) SetIsDeleted(v bool) *ProblemUpsert {
	u.Set(problem.FieldIsDeleted, v)
	return u
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateIsDeleted() *ProblemUpsert {
	u.SetExcluded(problem.FieldIsDeleted)
	return u
}

// SetConfig sets the "config" field.
func (u *ProblemUpsert) SetConfig(v string) *ProblemUpsert {
	u.Set(problem.FieldConfig, v)
	return u
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateConfig() *ProblemUpsert {
	u.SetExcluded(problem.FieldConfig)
	return u
}

// SetContestID sets the "contest_id" field.
func (u *ProblemUpsert) SetContestID(v int64) *ProblemUpsert {
	u.Set(problem.FieldContestID, v)
	return u
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateContestID() *ProblemUpsert {
	u.SetExcluded(problem.FieldContestID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ProblemUpsert) SetUserID(v int64) *ProblemUpsert {
	u.Set(problem.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateUserID() *ProblemUpsert {
	u.SetExcluded(problem.FieldUserID)
	return u
}

// SetVisibility sets the "visibility" field.
func (u *ProblemUpsert) SetVisibility(v int8) *ProblemUpsert {
	u.Set(problem.FieldVisibility, v)
	return u
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *ProblemUpsert) UpdateVisibility() *ProblemUpsert {
	u.SetExcluded(problem.FieldVisibility)
	return u
}

// AddVisibility adds v to the "visibility" field.
func (u *ProblemUpsert) AddVisibility(v int8) *ProblemUpsert {
	u.Add(problem.FieldVisibility, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Problem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(problem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProblemUpsertOne) UpdateNewValues() *ProblemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(problem.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Problem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProblemUpsertOne) Ignore() *ProblemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProblemUpsertOne) DoNothing() *ProblemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProblemCreate.OnConflict
// documentation for more info.
func (u *ProblemUpsertOne) Update(set func(*ProblemUpsert)) *ProblemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProblemUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *ProblemUpsertOne) SetTitle(v string) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateTitle() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *ProblemUpsertOne) SetContent(v string) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateContent() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateContent()
	})
}

// SetPoint sets the "point" field.
func (u *ProblemUpsertOne) SetPoint(v int16) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetPoint(v)
	})
}

// AddPoint adds v to the "point" field.
func (u *ProblemUpsertOne) AddPoint(v int16) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.AddPoint(v)
	})
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdatePoint() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdatePoint()
	})
}

// SetCaseVersion sets the "case_version" field.
func (u *ProblemUpsertOne) SetCaseVersion(v int16) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetCaseVersion(v)
	})
}

// AddCaseVersion adds v to the "case_version" field.
func (u *ProblemUpsertOne) AddCaseVersion(v int16) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.AddCaseVersion(v)
	})
}

// UpdateCaseVersion sets the "case_version" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateCaseVersion() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateCaseVersion()
	})
}

// SetIndex sets the "index" field.
func (u *ProblemUpsertOne) SetIndex(v int16) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *ProblemUpsertOne) AddIndex(v int16) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateIndex() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateIndex()
	})
}

// SetRestrictPresentation sets the "restrict_presentation" field.
func (u *ProblemUpsertOne) SetRestrictPresentation(v bool) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetRestrictPresentation(v)
	})
}

// UpdateRestrictPresentation sets the "restrict_presentation" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateRestrictPresentation() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateRestrictPresentation()
	})
}

// SetIsDeleted sets the "is_deleted" field.
func (u *ProblemUpsertOne) SetIsDeleted(v bool) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetIsDeleted(v)
	})
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateIsDeleted() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateIsDeleted()
	})
}

// SetConfig sets the "config" field.
func (u *ProblemUpsertOne) SetConfig(v string) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateConfig() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateConfig()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ProblemUpsertOne) SetContestID(v int64) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateContestID() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateContestID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ProblemUpsertOne) SetUserID(v int64) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateUserID() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateUserID()
	})
}

// SetVisibility sets the "visibility" field.
func (u *ProblemUpsertOne) SetVisibility(v int8) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.SetVisibility(v)
	})
}

// AddVisibility adds v to the "visibility" field.
func (u *ProblemUpsertOne) AddVisibility(v int8) *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.AddVisibility(v)
	})
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *ProblemUpsertOne) UpdateVisibility() *ProblemUpsertOne {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateVisibility()
	})
}

// Exec executes the query.
func (u *ProblemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProblemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProblemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProblemUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProblemUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProblemCreateBulk is the builder for creating many Problem entities in bulk.
type ProblemCreateBulk struct {
	config
	err      error
	builders []*ProblemCreate
	conflict []sql.ConflictOption
}

// Save creates the Problem entities in the database.
func (pcb *ProblemCreateBulk) Save(ctx context.Context) ([]*Problem, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Problem, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProblemMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProblemCreateBulk) SaveX(ctx context.Context) []*Problem {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProblemCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProblemCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Problem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProblemUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProblemCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProblemUpsertBulk {
	pcb.conflict = opts
	return &ProblemUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Problem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProblemCreateBulk) OnConflictColumns(columns ...string) *ProblemUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProblemUpsertBulk{
		create: pcb,
	}
}

// ProblemUpsertBulk is the builder for "upsert"-ing
// a bulk of Problem nodes.
type ProblemUpsertBulk struct {
	create *ProblemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Problem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(problem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProblemUpsertBulk) UpdateNewValues() *ProblemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(problem.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Problem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProblemUpsertBulk) Ignore() *ProblemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProblemUpsertBulk) DoNothing() *ProblemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProblemCreateBulk.OnConflict
// documentation for more info.
func (u *ProblemUpsertBulk) Update(set func(*ProblemUpsert)) *ProblemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProblemUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *ProblemUpsertBulk) SetTitle(v string) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateTitle() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *ProblemUpsertBulk) SetContent(v string) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateContent() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateContent()
	})
}

// SetPoint sets the "point" field.
func (u *ProblemUpsertBulk) SetPoint(v int16) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetPoint(v)
	})
}

// AddPoint adds v to the "point" field.
func (u *ProblemUpsertBulk) AddPoint(v int16) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.AddPoint(v)
	})
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdatePoint() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdatePoint()
	})
}

// SetCaseVersion sets the "case_version" field.
func (u *ProblemUpsertBulk) SetCaseVersion(v int16) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetCaseVersion(v)
	})
}

// AddCaseVersion adds v to the "case_version" field.
func (u *ProblemUpsertBulk) AddCaseVersion(v int16) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.AddCaseVersion(v)
	})
}

// UpdateCaseVersion sets the "case_version" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateCaseVersion() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateCaseVersion()
	})
}

// SetIndex sets the "index" field.
func (u *ProblemUpsertBulk) SetIndex(v int16) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *ProblemUpsertBulk) AddIndex(v int16) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateIndex() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateIndex()
	})
}

// SetRestrictPresentation sets the "restrict_presentation" field.
func (u *ProblemUpsertBulk) SetRestrictPresentation(v bool) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetRestrictPresentation(v)
	})
}

// UpdateRestrictPresentation sets the "restrict_presentation" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateRestrictPresentation() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateRestrictPresentation()
	})
}

// SetIsDeleted sets the "is_deleted" field.
func (u *ProblemUpsertBulk) SetIsDeleted(v bool) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetIsDeleted(v)
	})
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateIsDeleted() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateIsDeleted()
	})
}

// SetConfig sets the "config" field.
func (u *ProblemUpsertBulk) SetConfig(v string) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetConfig(v)
	})
}

// UpdateConfig sets the "config" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateConfig() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateConfig()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ProblemUpsertBulk) SetContestID(v int64) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateContestID() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateContestID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ProblemUpsertBulk) SetUserID(v int64) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateUserID() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateUserID()
	})
}

// SetVisibility sets the "visibility" field.
func (u *ProblemUpsertBulk) SetVisibility(v int8) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.SetVisibility(v)
	})
}

// AddVisibility adds v to the "visibility" field.
func (u *ProblemUpsertBulk) AddVisibility(v int8) *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.AddVisibility(v)
	})
}

// UpdateVisibility sets the "visibility" field to the value that was provided on create.
func (u *ProblemUpsertBulk) UpdateVisibility() *ProblemUpsertBulk {
	return u.Update(func(s *ProblemUpsert) {
		s.UpdateVisibility()
	})
}

// Exec executes the query.
func (u *ProblemUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProblemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProblemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProblemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
