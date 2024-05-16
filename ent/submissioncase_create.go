// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/problemcase"
	"sastoj/ent/submission"
	"sastoj/ent/submissioncase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmissionCaseCreate is the builder for creating a SubmissionCase entity.
type SubmissionCaseCreate struct {
	config
	mutation *SubmissionCaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetState sets the "state" field.
func (scc *SubmissionCaseCreate) SetState(i int16) *SubmissionCaseCreate {
	scc.mutation.SetState(i)
	return scc
}

// SetPoint sets the "point" field.
func (scc *SubmissionCaseCreate) SetPoint(i int16) *SubmissionCaseCreate {
	scc.mutation.SetPoint(i)
	return scc
}

// SetMessage sets the "message" field.
func (scc *SubmissionCaseCreate) SetMessage(s string) *SubmissionCaseCreate {
	scc.mutation.SetMessage(s)
	return scc
}

// SetTime sets the "time" field.
func (scc *SubmissionCaseCreate) SetTime(i int32) *SubmissionCaseCreate {
	scc.mutation.SetTime(i)
	return scc
}

// SetMemory sets the "memory" field.
func (scc *SubmissionCaseCreate) SetMemory(i int32) *SubmissionCaseCreate {
	scc.mutation.SetMemory(i)
	return scc
}

// SetSubmissionID sets the "submission_id" field.
func (scc *SubmissionCaseCreate) SetSubmissionID(i int64) *SubmissionCaseCreate {
	scc.mutation.SetSubmissionID(i)
	return scc
}

// SetProblemCaseID sets the "problem_case_id" field.
func (scc *SubmissionCaseCreate) SetProblemCaseID(i int64) *SubmissionCaseCreate {
	scc.mutation.SetProblemCaseID(i)
	return scc
}

// SetID sets the "id" field.
func (scc *SubmissionCaseCreate) SetID(i int64) *SubmissionCaseCreate {
	scc.mutation.SetID(i)
	return scc
}

// SetSubmission sets the "submission" edge to the Submission entity.
func (scc *SubmissionCaseCreate) SetSubmission(s *Submission) *SubmissionCaseCreate {
	return scc.SetSubmissionID(s.ID)
}

// SetProblemCasesID sets the "problem_cases" edge to the ProblemCase entity by ID.
func (scc *SubmissionCaseCreate) SetProblemCasesID(id int64) *SubmissionCaseCreate {
	scc.mutation.SetProblemCasesID(id)
	return scc
}

// SetProblemCases sets the "problem_cases" edge to the ProblemCase entity.
func (scc *SubmissionCaseCreate) SetProblemCases(p *ProblemCase) *SubmissionCaseCreate {
	return scc.SetProblemCasesID(p.ID)
}

// Mutation returns the SubmissionCaseMutation object of the builder.
func (scc *SubmissionCaseCreate) Mutation() *SubmissionCaseMutation {
	return scc.mutation
}

// Save creates the SubmissionCase in the database.
func (scc *SubmissionCaseCreate) Save(ctx context.Context) (*SubmissionCase, error) {
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SubmissionCaseCreate) SaveX(ctx context.Context) *SubmissionCase {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SubmissionCaseCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SubmissionCaseCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SubmissionCaseCreate) check() error {
	if _, ok := scc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "SubmissionCase.state"`)}
	}
	if v, ok := scc.mutation.State(); ok {
		if err := submissioncase.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "SubmissionCase.state": %w`, err)}
		}
	}
	if _, ok := scc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "SubmissionCase.point"`)}
	}
	if v, ok := scc.mutation.Point(); ok {
		if err := submissioncase.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "SubmissionCase.point": %w`, err)}
		}
	}
	if _, ok := scc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "SubmissionCase.message"`)}
	}
	if _, ok := scc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "SubmissionCase.time"`)}
	}
	if v, ok := scc.mutation.Time(); ok {
		if err := submissioncase.TimeValidator(v); err != nil {
			return &ValidationError{Name: "time", err: fmt.Errorf(`ent: validator failed for field "SubmissionCase.time": %w`, err)}
		}
	}
	if _, ok := scc.mutation.Memory(); !ok {
		return &ValidationError{Name: "memory", err: errors.New(`ent: missing required field "SubmissionCase.memory"`)}
	}
	if v, ok := scc.mutation.Memory(); ok {
		if err := submissioncase.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "SubmissionCase.memory": %w`, err)}
		}
	}
	if _, ok := scc.mutation.SubmissionID(); !ok {
		return &ValidationError{Name: "submission_id", err: errors.New(`ent: missing required field "SubmissionCase.submission_id"`)}
	}
	if _, ok := scc.mutation.ProblemCaseID(); !ok {
		return &ValidationError{Name: "problem_case_id", err: errors.New(`ent: missing required field "SubmissionCase.problem_case_id"`)}
	}
	if _, ok := scc.mutation.SubmissionID(); !ok {
		return &ValidationError{Name: "submission", err: errors.New(`ent: missing required edge "SubmissionCase.submission"`)}
	}
	if _, ok := scc.mutation.ProblemCasesID(); !ok {
		return &ValidationError{Name: "problem_cases", err: errors.New(`ent: missing required edge "SubmissionCase.problem_cases"`)}
	}
	return nil
}

func (scc *SubmissionCaseCreate) sqlSave(ctx context.Context) (*SubmissionCase, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *SubmissionCaseCreate) createSpec() (*SubmissionCase, *sqlgraph.CreateSpec) {
	var (
		_node = &SubmissionCase{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(submissioncase.Table, sqlgraph.NewFieldSpec(submissioncase.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = scc.conflict
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := scc.mutation.State(); ok {
		_spec.SetField(submissioncase.FieldState, field.TypeInt16, value)
		_node.State = value
	}
	if value, ok := scc.mutation.Point(); ok {
		_spec.SetField(submissioncase.FieldPoint, field.TypeInt16, value)
		_node.Point = value
	}
	if value, ok := scc.mutation.Message(); ok {
		_spec.SetField(submissioncase.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := scc.mutation.Time(); ok {
		_spec.SetField(submissioncase.FieldTime, field.TypeInt32, value)
		_node.Time = value
	}
	if value, ok := scc.mutation.Memory(); ok {
		_spec.SetField(submissioncase.FieldMemory, field.TypeInt32, value)
		_node.Memory = value
	}
	if nodes := scc.mutation.SubmissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submissioncase.SubmissionTable,
			Columns: []string{submissioncase.SubmissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubmissionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.ProblemCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   submissioncase.ProblemCasesTable,
			Columns: []string{submissioncase.ProblemCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProblemCaseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubmissionCase.Create().
//		SetState(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubmissionCaseUpsert) {
//			SetState(v+v).
//		}).
//		Exec(ctx)
func (scc *SubmissionCaseCreate) OnConflict(opts ...sql.ConflictOption) *SubmissionCaseUpsertOne {
	scc.conflict = opts
	return &SubmissionCaseUpsertOne{
		create: scc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubmissionCase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scc *SubmissionCaseCreate) OnConflictColumns(columns ...string) *SubmissionCaseUpsertOne {
	scc.conflict = append(scc.conflict, sql.ConflictColumns(columns...))
	return &SubmissionCaseUpsertOne{
		create: scc,
	}
}

type (
	// SubmissionCaseUpsertOne is the builder for "upsert"-ing
	//  one SubmissionCase node.
	SubmissionCaseUpsertOne struct {
		create *SubmissionCaseCreate
	}

	// SubmissionCaseUpsert is the "OnConflict" setter.
	SubmissionCaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetState sets the "state" field.
func (u *SubmissionCaseUpsert) SetState(v int16) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdateState() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldState)
	return u
}

// AddState adds v to the "state" field.
func (u *SubmissionCaseUpsert) AddState(v int16) *SubmissionCaseUpsert {
	u.Add(submissioncase.FieldState, v)
	return u
}

// SetPoint sets the "point" field.
func (u *SubmissionCaseUpsert) SetPoint(v int16) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldPoint, v)
	return u
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdatePoint() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldPoint)
	return u
}

// AddPoint adds v to the "point" field.
func (u *SubmissionCaseUpsert) AddPoint(v int16) *SubmissionCaseUpsert {
	u.Add(submissioncase.FieldPoint, v)
	return u
}

// SetMessage sets the "message" field.
func (u *SubmissionCaseUpsert) SetMessage(v string) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdateMessage() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldMessage)
	return u
}

// SetTime sets the "time" field.
func (u *SubmissionCaseUpsert) SetTime(v int32) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdateTime() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldTime)
	return u
}

// AddTime adds v to the "time" field.
func (u *SubmissionCaseUpsert) AddTime(v int32) *SubmissionCaseUpsert {
	u.Add(submissioncase.FieldTime, v)
	return u
}

// SetMemory sets the "memory" field.
func (u *SubmissionCaseUpsert) SetMemory(v int32) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldMemory, v)
	return u
}

// UpdateMemory sets the "memory" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdateMemory() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldMemory)
	return u
}

// AddMemory adds v to the "memory" field.
func (u *SubmissionCaseUpsert) AddMemory(v int32) *SubmissionCaseUpsert {
	u.Add(submissioncase.FieldMemory, v)
	return u
}

// SetSubmissionID sets the "submission_id" field.
func (u *SubmissionCaseUpsert) SetSubmissionID(v int64) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldSubmissionID, v)
	return u
}

// UpdateSubmissionID sets the "submission_id" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdateSubmissionID() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldSubmissionID)
	return u
}

// SetProblemCaseID sets the "problem_case_id" field.
func (u *SubmissionCaseUpsert) SetProblemCaseID(v int64) *SubmissionCaseUpsert {
	u.Set(submissioncase.FieldProblemCaseID, v)
	return u
}

// UpdateProblemCaseID sets the "problem_case_id" field to the value that was provided on create.
func (u *SubmissionCaseUpsert) UpdateProblemCaseID() *SubmissionCaseUpsert {
	u.SetExcluded(submissioncase.FieldProblemCaseID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SubmissionCase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(submissioncase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubmissionCaseUpsertOne) UpdateNewValues() *SubmissionCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(submissioncase.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubmissionCase.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SubmissionCaseUpsertOne) Ignore() *SubmissionCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubmissionCaseUpsertOne) DoNothing() *SubmissionCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubmissionCaseCreate.OnConflict
// documentation for more info.
func (u *SubmissionCaseUpsertOne) Update(set func(*SubmissionCaseUpsert)) *SubmissionCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubmissionCaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetState sets the "state" field.
func (u *SubmissionCaseUpsertOne) SetState(v int16) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *SubmissionCaseUpsertOne) AddState(v int16) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdateState() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateState()
	})
}

// SetPoint sets the "point" field.
func (u *SubmissionCaseUpsertOne) SetPoint(v int16) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetPoint(v)
	})
}

// AddPoint adds v to the "point" field.
func (u *SubmissionCaseUpsertOne) AddPoint(v int16) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddPoint(v)
	})
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdatePoint() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdatePoint()
	})
}

// SetMessage sets the "message" field.
func (u *SubmissionCaseUpsertOne) SetMessage(v string) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdateMessage() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateMessage()
	})
}

// SetTime sets the "time" field.
func (u *SubmissionCaseUpsertOne) SetTime(v int32) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *SubmissionCaseUpsertOne) AddTime(v int32) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdateTime() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateTime()
	})
}

// SetMemory sets the "memory" field.
func (u *SubmissionCaseUpsertOne) SetMemory(v int32) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetMemory(v)
	})
}

// AddMemory adds v to the "memory" field.
func (u *SubmissionCaseUpsertOne) AddMemory(v int32) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddMemory(v)
	})
}

// UpdateMemory sets the "memory" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdateMemory() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateMemory()
	})
}

// SetSubmissionID sets the "submission_id" field.
func (u *SubmissionCaseUpsertOne) SetSubmissionID(v int64) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetSubmissionID(v)
	})
}

// UpdateSubmissionID sets the "submission_id" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdateSubmissionID() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateSubmissionID()
	})
}

// SetProblemCaseID sets the "problem_case_id" field.
func (u *SubmissionCaseUpsertOne) SetProblemCaseID(v int64) *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetProblemCaseID(v)
	})
}

// UpdateProblemCaseID sets the "problem_case_id" field to the value that was provided on create.
func (u *SubmissionCaseUpsertOne) UpdateProblemCaseID() *SubmissionCaseUpsertOne {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateProblemCaseID()
	})
}

// Exec executes the query.
func (u *SubmissionCaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubmissionCaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubmissionCaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubmissionCaseUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubmissionCaseUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubmissionCaseCreateBulk is the builder for creating many SubmissionCase entities in bulk.
type SubmissionCaseCreateBulk struct {
	config
	err      error
	builders []*SubmissionCaseCreate
	conflict []sql.ConflictOption
}

// Save creates the SubmissionCase entities in the database.
func (sccb *SubmissionCaseCreateBulk) Save(ctx context.Context) ([]*SubmissionCase, error) {
	if sccb.err != nil {
		return nil, sccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SubmissionCase, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubmissionCaseMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SubmissionCaseCreateBulk) SaveX(ctx context.Context) []*SubmissionCase {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SubmissionCaseCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SubmissionCaseCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SubmissionCase.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubmissionCaseUpsert) {
//			SetState(v+v).
//		}).
//		Exec(ctx)
func (sccb *SubmissionCaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubmissionCaseUpsertBulk {
	sccb.conflict = opts
	return &SubmissionCaseUpsertBulk{
		create: sccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SubmissionCase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sccb *SubmissionCaseCreateBulk) OnConflictColumns(columns ...string) *SubmissionCaseUpsertBulk {
	sccb.conflict = append(sccb.conflict, sql.ConflictColumns(columns...))
	return &SubmissionCaseUpsertBulk{
		create: sccb,
	}
}

// SubmissionCaseUpsertBulk is the builder for "upsert"-ing
// a bulk of SubmissionCase nodes.
type SubmissionCaseUpsertBulk struct {
	create *SubmissionCaseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SubmissionCase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(submissioncase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubmissionCaseUpsertBulk) UpdateNewValues() *SubmissionCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(submissioncase.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SubmissionCase.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SubmissionCaseUpsertBulk) Ignore() *SubmissionCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubmissionCaseUpsertBulk) DoNothing() *SubmissionCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubmissionCaseCreateBulk.OnConflict
// documentation for more info.
func (u *SubmissionCaseUpsertBulk) Update(set func(*SubmissionCaseUpsert)) *SubmissionCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubmissionCaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetState sets the "state" field.
func (u *SubmissionCaseUpsertBulk) SetState(v int16) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetState(v)
	})
}

// AddState adds v to the "state" field.
func (u *SubmissionCaseUpsertBulk) AddState(v int16) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdateState() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateState()
	})
}

// SetPoint sets the "point" field.
func (u *SubmissionCaseUpsertBulk) SetPoint(v int16) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetPoint(v)
	})
}

// AddPoint adds v to the "point" field.
func (u *SubmissionCaseUpsertBulk) AddPoint(v int16) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddPoint(v)
	})
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdatePoint() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdatePoint()
	})
}

// SetMessage sets the "message" field.
func (u *SubmissionCaseUpsertBulk) SetMessage(v string) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdateMessage() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateMessage()
	})
}

// SetTime sets the "time" field.
func (u *SubmissionCaseUpsertBulk) SetTime(v int32) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *SubmissionCaseUpsertBulk) AddTime(v int32) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdateTime() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateTime()
	})
}

// SetMemory sets the "memory" field.
func (u *SubmissionCaseUpsertBulk) SetMemory(v int32) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetMemory(v)
	})
}

// AddMemory adds v to the "memory" field.
func (u *SubmissionCaseUpsertBulk) AddMemory(v int32) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.AddMemory(v)
	})
}

// UpdateMemory sets the "memory" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdateMemory() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateMemory()
	})
}

// SetSubmissionID sets the "submission_id" field.
func (u *SubmissionCaseUpsertBulk) SetSubmissionID(v int64) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetSubmissionID(v)
	})
}

// UpdateSubmissionID sets the "submission_id" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdateSubmissionID() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateSubmissionID()
	})
}

// SetProblemCaseID sets the "problem_case_id" field.
func (u *SubmissionCaseUpsertBulk) SetProblemCaseID(v int64) *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.SetProblemCaseID(v)
	})
}

// UpdateProblemCaseID sets the "problem_case_id" field to the value that was provided on create.
func (u *SubmissionCaseUpsertBulk) UpdateProblemCaseID() *SubmissionCaseUpsertBulk {
	return u.Update(func(s *SubmissionCaseUpsert) {
		s.UpdateProblemCaseID()
	})
}

// Exec executes the query.
func (u *SubmissionCaseUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SubmissionCaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubmissionCaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubmissionCaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
