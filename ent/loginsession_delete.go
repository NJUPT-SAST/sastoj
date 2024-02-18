// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sastoj/ent/loginsession"
	"sastoj/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginSessionDelete is the builder for deleting a LoginSession entity.
type LoginSessionDelete struct {
	config
	hooks    []Hook
	mutation *LoginSessionMutation
}

// Where appends a list predicates to the LoginSessionDelete builder.
func (lsd *LoginSessionDelete) Where(ps ...predicate.LoginSession) *LoginSessionDelete {
	lsd.mutation.Where(ps...)
	return lsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lsd *LoginSessionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lsd.sqlExec, lsd.mutation, lsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lsd *LoginSessionDelete) ExecX(ctx context.Context) int {
	n, err := lsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lsd *LoginSessionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(loginsession.Table, sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt))
	if ps := lsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lsd.mutation.done = true
	return affected, err
}

// LoginSessionDeleteOne is the builder for deleting a single LoginSession entity.
type LoginSessionDeleteOne struct {
	lsd *LoginSessionDelete
}

// Where appends a list predicates to the LoginSessionDelete builder.
func (lsdo *LoginSessionDeleteOne) Where(ps ...predicate.LoginSession) *LoginSessionDeleteOne {
	lsdo.lsd.mutation.Where(ps...)
	return lsdo
}

// Exec executes the deletion query.
func (lsdo *LoginSessionDeleteOne) Exec(ctx context.Context) error {
	n, err := lsdo.lsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{loginsession.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lsdo *LoginSessionDeleteOne) ExecX(ctx context.Context) {
	if err := lsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
