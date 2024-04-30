// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sastoj/ent/predicate"
	"sastoj/ent/submissioncase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubmissionCaseDelete is the builder for deleting a SubmissionCase entity.
type SubmissionCaseDelete struct {
	config
	hooks    []Hook
	mutation *SubmissionCaseMutation
}

// Where appends a list predicates to the SubmissionCaseDelete builder.
func (scd *SubmissionCaseDelete) Where(ps ...predicate.SubmissionCase) *SubmissionCaseDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SubmissionCaseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SubmissionCaseDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SubmissionCaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(submissioncase.Table, sqlgraph.NewFieldSpec(submissioncase.FieldID, field.TypeInt64))
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

// SubmissionCaseDeleteOne is the builder for deleting a single SubmissionCase entity.
type SubmissionCaseDeleteOne struct {
	scd *SubmissionCaseDelete
}

// Where appends a list predicates to the SubmissionCaseDelete builder.
func (scdo *SubmissionCaseDeleteOne) Where(ps ...predicate.SubmissionCase) *SubmissionCaseDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *SubmissionCaseDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{submissioncase.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SubmissionCaseDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}