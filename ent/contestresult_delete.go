// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sastoj/ent/contestresult"
	"sastoj/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContestResultDelete is the builder for deleting a ContestResult entity.
type ContestResultDelete struct {
	config
	hooks    []Hook
	mutation *ContestResultMutation
}

// Where appends a list predicates to the ContestResultDelete builder.
func (crd *ContestResultDelete) Where(ps ...predicate.ContestResult) *ContestResultDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *ContestResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *ContestResultDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *ContestResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(contestresult.Table, sqlgraph.NewFieldSpec(contestresult.FieldID, field.TypeInt))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// ContestResultDeleteOne is the builder for deleting a single ContestResult entity.
type ContestResultDeleteOne struct {
	crd *ContestResultDelete
}

// Where appends a list predicates to the ContestResultDelete builder.
func (crdo *ContestResultDeleteOne) Where(ps ...predicate.ContestResult) *ContestResultDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *ContestResultDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{contestresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *ContestResultDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}