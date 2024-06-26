// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ilcm96/dku-aegis-library/ent/bookrequest"
	"github.com/ilcm96/dku-aegis-library/ent/predicate"
)

// BookRequestDelete is the builder for deleting a BookRequest entity.
type BookRequestDelete struct {
	config
	hooks    []Hook
	mutation *BookRequestMutation
}

// Where appends a list predicates to the BookRequestDelete builder.
func (brd *BookRequestDelete) Where(ps ...predicate.BookRequest) *BookRequestDelete {
	brd.mutation.Where(ps...)
	return brd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (brd *BookRequestDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, brd.sqlExec, brd.mutation, brd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (brd *BookRequestDelete) ExecX(ctx context.Context) int {
	n, err := brd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (brd *BookRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bookrequest.Table, sqlgraph.NewFieldSpec(bookrequest.FieldID, field.TypeInt))
	if ps := brd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, brd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	brd.mutation.done = true
	return affected, err
}

// BookRequestDeleteOne is the builder for deleting a single BookRequest entity.
type BookRequestDeleteOne struct {
	brd *BookRequestDelete
}

// Where appends a list predicates to the BookRequestDelete builder.
func (brdo *BookRequestDeleteOne) Where(ps ...predicate.BookRequest) *BookRequestDeleteOne {
	brdo.brd.mutation.Where(ps...)
	return brdo
}

// Exec executes the deletion query.
func (brdo *BookRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := brdo.brd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bookrequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (brdo *BookRequestDeleteOne) ExecX(ctx context.Context) {
	if err := brdo.Exec(ctx); err != nil {
		panic(err)
	}
}
