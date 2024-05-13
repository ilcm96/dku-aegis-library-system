// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ilcm96/dku-aegis-library/ent/booklog"
	"github.com/ilcm96/dku-aegis-library/ent/predicate"
)

// BookLogDelete is the builder for deleting a BookLog entity.
type BookLogDelete struct {
	config
	hooks    []Hook
	mutation *BookLogMutation
}

// Where appends a list predicates to the BookLogDelete builder.
func (bld *BookLogDelete) Where(ps ...predicate.BookLog) *BookLogDelete {
	bld.mutation.Where(ps...)
	return bld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bld *BookLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bld.sqlExec, bld.mutation, bld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bld *BookLogDelete) ExecX(ctx context.Context) int {
	n, err := bld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bld *BookLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(booklog.Table, sqlgraph.NewFieldSpec(booklog.FieldID, field.TypeInt))
	if ps := bld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bld.mutation.done = true
	return affected, err
}

// BookLogDeleteOne is the builder for deleting a single BookLog entity.
type BookLogDeleteOne struct {
	bld *BookLogDelete
}

// Where appends a list predicates to the BookLogDelete builder.
func (bldo *BookLogDeleteOne) Where(ps ...predicate.BookLog) *BookLogDeleteOne {
	bldo.bld.mutation.Where(ps...)
	return bldo
}

// Exec executes the deletion query.
func (bldo *BookLogDeleteOne) Exec(ctx context.Context) error {
	n, err := bldo.bld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{booklog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bldo *BookLogDeleteOne) ExecX(ctx context.Context) {
	if err := bldo.Exec(ctx); err != nil {
		panic(err)
	}
}