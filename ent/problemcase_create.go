// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"sastoj/ent/submissioncase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemCaseCreate is the builder for creating a ProblemCase entity.
type ProblemCaseCreate struct {
	config
	mutation *ProblemCaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPoint sets the "point" field.
func (pcc *ProblemCaseCreate) SetPoint(i int16) *ProblemCaseCreate {
	pcc.mutation.SetPoint(i)
	return pcc
}

// SetIndex sets the "index" field.
func (pcc *ProblemCaseCreate) SetIndex(i int16) *ProblemCaseCreate {
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

// SetProblemID sets the "problem_id" field.
func (pcc *ProblemCaseCreate) SetProblemID(i int64) *ProblemCaseCreate {
	pcc.mutation.SetProblemID(i)
	return pcc
}

// SetID sets the "id" field.
func (pcc *ProblemCaseCreate) SetID(i int64) *ProblemCaseCreate {
	pcc.mutation.SetID(i)
	return pcc
}

// AddSubmissionCaseIDs adds the "submission_cases" edge to the SubmissionCase entity by IDs.
func (pcc *ProblemCaseCreate) AddSubmissionCaseIDs(ids ...int64) *ProblemCaseCreate {
	pcc.mutation.AddSubmissionCaseIDs(ids...)
	return pcc
}

// AddSubmissionCases adds the "submission_cases" edges to the SubmissionCase entity.
func (pcc *ProblemCaseCreate) AddSubmissionCases(s ...*SubmissionCase) *ProblemCaseCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pcc.AddSubmissionCaseIDs(ids...)
}

// SetProblem sets the "problem" edge to the Problem entity.
func (pcc *ProblemCaseCreate) SetProblem(p *Problem) *ProblemCaseCreate {
	return pcc.SetProblemID(p.ID)
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
	if _, ok := pcc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "ProblemCase.point"`)}
	}
	if v, ok := pcc.mutation.Point(); ok {
		if err := problemcase.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "ProblemCase.point": %w`, err)}
		}
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
	if _, ok := pcc.mutation.ProblemID(); !ok {
		return &ValidationError{Name: "problem_id", err: errors.New(`ent: missing required field "ProblemCase.problem_id"`)}
	}
	if _, ok := pcc.mutation.ProblemID(); !ok {
		return &ValidationError{Name: "problem", err: errors.New(`ent: missing required edge "ProblemCase.problem"`)}
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
		_node.ID = int64(id)
	}
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *ProblemCaseCreate) createSpec() (*ProblemCase, *sqlgraph.CreateSpec) {
	var (
		_node = &ProblemCase{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(problemcase.Table, sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = pcc.conflict
	if id, ok := pcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pcc.mutation.Point(); ok {
		_spec.SetField(problemcase.FieldPoint, field.TypeInt16, value)
		_node.Point = value
	}
	if value, ok := pcc.mutation.Index(); ok {
		_spec.SetField(problemcase.FieldIndex, field.TypeInt16, value)
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
	if nodes := pcc.mutation.SubmissionCasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   problemcase.SubmissionCasesTable,
			Columns: []string{problemcase.SubmissionCasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submissioncase.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.ProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   problemcase.ProblemTable,
			Columns: []string{problemcase.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProblemID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProblemCase.Create().
//		SetPoint(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProblemCaseUpsert) {
//			SetPoint(v+v).
//		}).
//		Exec(ctx)
func (pcc *ProblemCaseCreate) OnConflict(opts ...sql.ConflictOption) *ProblemCaseUpsertOne {
	pcc.conflict = opts
	return &ProblemCaseUpsertOne{
		create: pcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProblemCase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcc *ProblemCaseCreate) OnConflictColumns(columns ...string) *ProblemCaseUpsertOne {
	pcc.conflict = append(pcc.conflict, sql.ConflictColumns(columns...))
	return &ProblemCaseUpsertOne{
		create: pcc,
	}
}

type (
	// ProblemCaseUpsertOne is the builder for "upsert"-ing
	//  one ProblemCase node.
	ProblemCaseUpsertOne struct {
		create *ProblemCaseCreate
	}

	// ProblemCaseUpsert is the "OnConflict" setter.
	ProblemCaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetPoint sets the "point" field.
func (u *ProblemCaseUpsert) SetPoint(v int16) *ProblemCaseUpsert {
	u.Set(problemcase.FieldPoint, v)
	return u
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *ProblemCaseUpsert) UpdatePoint() *ProblemCaseUpsert {
	u.SetExcluded(problemcase.FieldPoint)
	return u
}

// AddPoint adds v to the "point" field.
func (u *ProblemCaseUpsert) AddPoint(v int16) *ProblemCaseUpsert {
	u.Add(problemcase.FieldPoint, v)
	return u
}

// SetIndex sets the "index" field.
func (u *ProblemCaseUpsert) SetIndex(v int16) *ProblemCaseUpsert {
	u.Set(problemcase.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *ProblemCaseUpsert) UpdateIndex() *ProblemCaseUpsert {
	u.SetExcluded(problemcase.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *ProblemCaseUpsert) AddIndex(v int16) *ProblemCaseUpsert {
	u.Add(problemcase.FieldIndex, v)
	return u
}

// SetIsAuto sets the "is_auto" field.
func (u *ProblemCaseUpsert) SetIsAuto(v bool) *ProblemCaseUpsert {
	u.Set(problemcase.FieldIsAuto, v)
	return u
}

// UpdateIsAuto sets the "is_auto" field to the value that was provided on create.
func (u *ProblemCaseUpsert) UpdateIsAuto() *ProblemCaseUpsert {
	u.SetExcluded(problemcase.FieldIsAuto)
	return u
}

// SetIsDeleted sets the "is_deleted" field.
func (u *ProblemCaseUpsert) SetIsDeleted(v bool) *ProblemCaseUpsert {
	u.Set(problemcase.FieldIsDeleted, v)
	return u
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *ProblemCaseUpsert) UpdateIsDeleted() *ProblemCaseUpsert {
	u.SetExcluded(problemcase.FieldIsDeleted)
	return u
}

// SetProblemID sets the "problem_id" field.
func (u *ProblemCaseUpsert) SetProblemID(v int64) *ProblemCaseUpsert {
	u.Set(problemcase.FieldProblemID, v)
	return u
}

// UpdateProblemID sets the "problem_id" field to the value that was provided on create.
func (u *ProblemCaseUpsert) UpdateProblemID() *ProblemCaseUpsert {
	u.SetExcluded(problemcase.FieldProblemID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ProblemCase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(problemcase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProblemCaseUpsertOne) UpdateNewValues() *ProblemCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(problemcase.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProblemCase.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProblemCaseUpsertOne) Ignore() *ProblemCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProblemCaseUpsertOne) DoNothing() *ProblemCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProblemCaseCreate.OnConflict
// documentation for more info.
func (u *ProblemCaseUpsertOne) Update(set func(*ProblemCaseUpsert)) *ProblemCaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProblemCaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetPoint sets the "point" field.
func (u *ProblemCaseUpsertOne) SetPoint(v int16) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetPoint(v)
	})
}

// AddPoint adds v to the "point" field.
func (u *ProblemCaseUpsertOne) AddPoint(v int16) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.AddPoint(v)
	})
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *ProblemCaseUpsertOne) UpdatePoint() *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdatePoint()
	})
}

// SetIndex sets the "index" field.
func (u *ProblemCaseUpsertOne) SetIndex(v int16) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *ProblemCaseUpsertOne) AddIndex(v int16) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *ProblemCaseUpsertOne) UpdateIndex() *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateIndex()
	})
}

// SetIsAuto sets the "is_auto" field.
func (u *ProblemCaseUpsertOne) SetIsAuto(v bool) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetIsAuto(v)
	})
}

// UpdateIsAuto sets the "is_auto" field to the value that was provided on create.
func (u *ProblemCaseUpsertOne) UpdateIsAuto() *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateIsAuto()
	})
}

// SetIsDeleted sets the "is_deleted" field.
func (u *ProblemCaseUpsertOne) SetIsDeleted(v bool) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetIsDeleted(v)
	})
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *ProblemCaseUpsertOne) UpdateIsDeleted() *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateIsDeleted()
	})
}

// SetProblemID sets the "problem_id" field.
func (u *ProblemCaseUpsertOne) SetProblemID(v int64) *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetProblemID(v)
	})
}

// UpdateProblemID sets the "problem_id" field to the value that was provided on create.
func (u *ProblemCaseUpsertOne) UpdateProblemID() *ProblemCaseUpsertOne {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateProblemID()
	})
}

// Exec executes the query.
func (u *ProblemCaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProblemCaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProblemCaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProblemCaseUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProblemCaseUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProblemCaseCreateBulk is the builder for creating many ProblemCase entities in bulk.
type ProblemCaseCreateBulk struct {
	config
	err      error
	builders []*ProblemCaseCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = pccb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProblemCase.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProblemCaseUpsert) {
//			SetPoint(v+v).
//		}).
//		Exec(ctx)
func (pccb *ProblemCaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProblemCaseUpsertBulk {
	pccb.conflict = opts
	return &ProblemCaseUpsertBulk{
		create: pccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProblemCase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pccb *ProblemCaseCreateBulk) OnConflictColumns(columns ...string) *ProblemCaseUpsertBulk {
	pccb.conflict = append(pccb.conflict, sql.ConflictColumns(columns...))
	return &ProblemCaseUpsertBulk{
		create: pccb,
	}
}

// ProblemCaseUpsertBulk is the builder for "upsert"-ing
// a bulk of ProblemCase nodes.
type ProblemCaseUpsertBulk struct {
	create *ProblemCaseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProblemCase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(problemcase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProblemCaseUpsertBulk) UpdateNewValues() *ProblemCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(problemcase.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProblemCase.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProblemCaseUpsertBulk) Ignore() *ProblemCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProblemCaseUpsertBulk) DoNothing() *ProblemCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProblemCaseCreateBulk.OnConflict
// documentation for more info.
func (u *ProblemCaseUpsertBulk) Update(set func(*ProblemCaseUpsert)) *ProblemCaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProblemCaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetPoint sets the "point" field.
func (u *ProblemCaseUpsertBulk) SetPoint(v int16) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetPoint(v)
	})
}

// AddPoint adds v to the "point" field.
func (u *ProblemCaseUpsertBulk) AddPoint(v int16) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.AddPoint(v)
	})
}

// UpdatePoint sets the "point" field to the value that was provided on create.
func (u *ProblemCaseUpsertBulk) UpdatePoint() *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdatePoint()
	})
}

// SetIndex sets the "index" field.
func (u *ProblemCaseUpsertBulk) SetIndex(v int16) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *ProblemCaseUpsertBulk) AddIndex(v int16) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *ProblemCaseUpsertBulk) UpdateIndex() *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateIndex()
	})
}

// SetIsAuto sets the "is_auto" field.
func (u *ProblemCaseUpsertBulk) SetIsAuto(v bool) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetIsAuto(v)
	})
}

// UpdateIsAuto sets the "is_auto" field to the value that was provided on create.
func (u *ProblemCaseUpsertBulk) UpdateIsAuto() *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateIsAuto()
	})
}

// SetIsDeleted sets the "is_deleted" field.
func (u *ProblemCaseUpsertBulk) SetIsDeleted(v bool) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetIsDeleted(v)
	})
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *ProblemCaseUpsertBulk) UpdateIsDeleted() *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateIsDeleted()
	})
}

// SetProblemID sets the "problem_id" field.
func (u *ProblemCaseUpsertBulk) SetProblemID(v int64) *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.SetProblemID(v)
	})
}

// UpdateProblemID sets the "problem_id" field to the value that was provided on create.
func (u *ProblemCaseUpsertBulk) UpdateProblemID() *ProblemCaseUpsertBulk {
	return u.Update(func(s *ProblemCaseUpsert) {
		s.UpdateProblemID()
	})
}

// Exec executes the query.
func (u *ProblemCaseUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProblemCaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProblemCaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProblemCaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
