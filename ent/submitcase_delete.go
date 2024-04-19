// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sastoj/ent/predicate"
	"sastoj/ent/submitcase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmitCaseDelete is the builder for deleting a SubmitCase entity.
type SubmitCaseDelete struct {
	config
	hooks    []Hook
	mutation *SubmitCaseMutation
}

// Where appends a list predicates to the SubmitCaseDelete builder.
func (scd *SubmitCaseDelete) Where(ps ...predicate.SubmitCase) *SubmitCaseDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SubmitCaseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SubmitCaseDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SubmitCaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(submitcase.Table, sqlgraph.NewFieldSpec(submitcase.FieldID, field.TypeInt64))
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	scd.mutation.done = true
	return affected, err
}

// SubmitCaseDeleteOne is the builder for deleting a single SubmitCase entity.
type SubmitCaseDeleteOne struct {
	scd *SubmitCaseDelete
}

// Where appends a list predicates to the SubmitCaseDelete builder.
func (scdo *SubmitCaseDeleteOne) Where(ps ...predicate.SubmitCase) *SubmitCaseDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *SubmitCaseDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{submitcase.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SubmitCaseDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}
