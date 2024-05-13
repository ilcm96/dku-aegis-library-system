// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ilcm96/dku-aegis-library/ent/booklog"
	"github.com/ilcm96/dku-aegis-library/ent/predicate"
)

// BookLogUpdate is the builder for updating BookLog entities.
type BookLogUpdate struct {
	config
	hooks    []Hook
	mutation *BookLogMutation
}

// Where appends a list predicates to the BookLogUpdate builder.
func (blu *BookLogUpdate) Where(ps ...predicate.BookLog) *BookLogUpdate {
	blu.mutation.Where(ps...)
	return blu
}

// SetAction sets the "action" field.
func (blu *BookLogUpdate) SetAction(b booklog.Action) *BookLogUpdate {
	blu.mutation.SetAction(b)
	return blu
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (blu *BookLogUpdate) SetNillableAction(b *booklog.Action) *BookLogUpdate {
	if b != nil {
		blu.SetAction(*b)
	}
	return blu
}

// SetUserID sets the "user_id" field.
func (blu *BookLogUpdate) SetUserID(i int) *BookLogUpdate {
	blu.mutation.ResetUserID()
	blu.mutation.SetUserID(i)
	return blu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (blu *BookLogUpdate) SetNillableUserID(i *int) *BookLogUpdate {
	if i != nil {
		blu.SetUserID(*i)
	}
	return blu
}

// AddUserID adds i to the "user_id" field.
func (blu *BookLogUpdate) AddUserID(i int) *BookLogUpdate {
	blu.mutation.AddUserID(i)
	return blu
}

// SetBookID sets the "book_id" field.
func (blu *BookLogUpdate) SetBookID(i int) *BookLogUpdate {
	blu.mutation.ResetBookID()
	blu.mutation.SetBookID(i)
	return blu
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (blu *BookLogUpdate) SetNillableBookID(i *int) *BookLogUpdate {
	if i != nil {
		blu.SetBookID(*i)
	}
	return blu
}

// AddBookID adds i to the "book_id" field.
func (blu *BookLogUpdate) AddBookID(i int) *BookLogUpdate {
	blu.mutation.AddBookID(i)
	return blu
}

// SetBookTitle sets the "book_title" field.
func (blu *BookLogUpdate) SetBookTitle(s string) *BookLogUpdate {
	blu.mutation.SetBookTitle(s)
	return blu
}

// SetNillableBookTitle sets the "book_title" field if the given value is not nil.
func (blu *BookLogUpdate) SetNillableBookTitle(s *string) *BookLogUpdate {
	if s != nil {
		blu.SetBookTitle(*s)
	}
	return blu
}

// SetRequestID sets the "request_id" field.
func (blu *BookLogUpdate) SetRequestID(s string) *BookLogUpdate {
	blu.mutation.SetRequestID(s)
	return blu
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (blu *BookLogUpdate) SetNillableRequestID(s *string) *BookLogUpdate {
	if s != nil {
		blu.SetRequestID(*s)
	}
	return blu
}

// SetCreatedAt sets the "created_at" field.
func (blu *BookLogUpdate) SetCreatedAt(t time.Time) *BookLogUpdate {
	blu.mutation.SetCreatedAt(t)
	return blu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (blu *BookLogUpdate) SetNillableCreatedAt(t *time.Time) *BookLogUpdate {
	if t != nil {
		blu.SetCreatedAt(*t)
	}
	return blu
}

// Mutation returns the BookLogMutation object of the builder.
func (blu *BookLogUpdate) Mutation() *BookLogMutation {
	return blu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (blu *BookLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, blu.sqlSave, blu.mutation, blu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (blu *BookLogUpdate) SaveX(ctx context.Context) int {
	affected, err := blu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (blu *BookLogUpdate) Exec(ctx context.Context) error {
	_, err := blu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blu *BookLogUpdate) ExecX(ctx context.Context) {
	if err := blu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (blu *BookLogUpdate) check() error {
	if v, ok := blu.mutation.Action(); ok {
		if err := booklog.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "BookLog.action": %w`, err)}
		}
	}
	return nil
}

func (blu *BookLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := blu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(booklog.Table, booklog.Columns, sqlgraph.NewFieldSpec(booklog.FieldID, field.TypeInt))
	if ps := blu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := blu.mutation.Action(); ok {
		_spec.SetField(booklog.FieldAction, field.TypeEnum, value)
	}
	if value, ok := blu.mutation.UserID(); ok {
		_spec.SetField(booklog.FieldUserID, field.TypeInt, value)
	}
	if value, ok := blu.mutation.AddedUserID(); ok {
		_spec.AddField(booklog.FieldUserID, field.TypeInt, value)
	}
	if value, ok := blu.mutation.BookID(); ok {
		_spec.SetField(booklog.FieldBookID, field.TypeInt, value)
	}
	if value, ok := blu.mutation.AddedBookID(); ok {
		_spec.AddField(booklog.FieldBookID, field.TypeInt, value)
	}
	if value, ok := blu.mutation.BookTitle(); ok {
		_spec.SetField(booklog.FieldBookTitle, field.TypeString, value)
	}
	if value, ok := blu.mutation.RequestID(); ok {
		_spec.SetField(booklog.FieldRequestID, field.TypeString, value)
	}
	if value, ok := blu.mutation.CreatedAt(); ok {
		_spec.SetField(booklog.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, blu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	blu.mutation.done = true
	return n, nil
}

// BookLogUpdateOne is the builder for updating a single BookLog entity.
type BookLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookLogMutation
}

// SetAction sets the "action" field.
func (bluo *BookLogUpdateOne) SetAction(b booklog.Action) *BookLogUpdateOne {
	bluo.mutation.SetAction(b)
	return bluo
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (bluo *BookLogUpdateOne) SetNillableAction(b *booklog.Action) *BookLogUpdateOne {
	if b != nil {
		bluo.SetAction(*b)
	}
	return bluo
}

// SetUserID sets the "user_id" field.
func (bluo *BookLogUpdateOne) SetUserID(i int) *BookLogUpdateOne {
	bluo.mutation.ResetUserID()
	bluo.mutation.SetUserID(i)
	return bluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bluo *BookLogUpdateOne) SetNillableUserID(i *int) *BookLogUpdateOne {
	if i != nil {
		bluo.SetUserID(*i)
	}
	return bluo
}

// AddUserID adds i to the "user_id" field.
func (bluo *BookLogUpdateOne) AddUserID(i int) *BookLogUpdateOne {
	bluo.mutation.AddUserID(i)
	return bluo
}

// SetBookID sets the "book_id" field.
func (bluo *BookLogUpdateOne) SetBookID(i int) *BookLogUpdateOne {
	bluo.mutation.ResetBookID()
	bluo.mutation.SetBookID(i)
	return bluo
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (bluo *BookLogUpdateOne) SetNillableBookID(i *int) *BookLogUpdateOne {
	if i != nil {
		bluo.SetBookID(*i)
	}
	return bluo
}

// AddBookID adds i to the "book_id" field.
func (bluo *BookLogUpdateOne) AddBookID(i int) *BookLogUpdateOne {
	bluo.mutation.AddBookID(i)
	return bluo
}

// SetBookTitle sets the "book_title" field.
func (bluo *BookLogUpdateOne) SetBookTitle(s string) *BookLogUpdateOne {
	bluo.mutation.SetBookTitle(s)
	return bluo
}

// SetNillableBookTitle sets the "book_title" field if the given value is not nil.
func (bluo *BookLogUpdateOne) SetNillableBookTitle(s *string) *BookLogUpdateOne {
	if s != nil {
		bluo.SetBookTitle(*s)
	}
	return bluo
}

// SetRequestID sets the "request_id" field.
func (bluo *BookLogUpdateOne) SetRequestID(s string) *BookLogUpdateOne {
	bluo.mutation.SetRequestID(s)
	return bluo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (bluo *BookLogUpdateOne) SetNillableRequestID(s *string) *BookLogUpdateOne {
	if s != nil {
		bluo.SetRequestID(*s)
	}
	return bluo
}

// SetCreatedAt sets the "created_at" field.
func (bluo *BookLogUpdateOne) SetCreatedAt(t time.Time) *BookLogUpdateOne {
	bluo.mutation.SetCreatedAt(t)
	return bluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bluo *BookLogUpdateOne) SetNillableCreatedAt(t *time.Time) *BookLogUpdateOne {
	if t != nil {
		bluo.SetCreatedAt(*t)
	}
	return bluo
}

// Mutation returns the BookLogMutation object of the builder.
func (bluo *BookLogUpdateOne) Mutation() *BookLogMutation {
	return bluo.mutation
}

// Where appends a list predicates to the BookLogUpdate builder.
func (bluo *BookLogUpdateOne) Where(ps ...predicate.BookLog) *BookLogUpdateOne {
	bluo.mutation.Where(ps...)
	return bluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bluo *BookLogUpdateOne) Select(field string, fields ...string) *BookLogUpdateOne {
	bluo.fields = append([]string{field}, fields...)
	return bluo
}

// Save executes the query and returns the updated BookLog entity.
func (bluo *BookLogUpdateOne) Save(ctx context.Context) (*BookLog, error) {
	return withHooks(ctx, bluo.sqlSave, bluo.mutation, bluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bluo *BookLogUpdateOne) SaveX(ctx context.Context) *BookLog {
	node, err := bluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bluo *BookLogUpdateOne) Exec(ctx context.Context) error {
	_, err := bluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bluo *BookLogUpdateOne) ExecX(ctx context.Context) {
	if err := bluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bluo *BookLogUpdateOne) check() error {
	if v, ok := bluo.mutation.Action(); ok {
		if err := booklog.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "BookLog.action": %w`, err)}
		}
	}
	return nil
}

func (bluo *BookLogUpdateOne) sqlSave(ctx context.Context) (_node *BookLog, err error) {
	if err := bluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(booklog.Table, booklog.Columns, sqlgraph.NewFieldSpec(booklog.FieldID, field.TypeInt))
	id, ok := bluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booklog.FieldID)
		for _, f := range fields {
			if !booklog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booklog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bluo.mutation.Action(); ok {
		_spec.SetField(booklog.FieldAction, field.TypeEnum, value)
	}
	if value, ok := bluo.mutation.UserID(); ok {
		_spec.SetField(booklog.FieldUserID, field.TypeInt, value)
	}
	if value, ok := bluo.mutation.AddedUserID(); ok {
		_spec.AddField(booklog.FieldUserID, field.TypeInt, value)
	}
	if value, ok := bluo.mutation.BookID(); ok {
		_spec.SetField(booklog.FieldBookID, field.TypeInt, value)
	}
	if value, ok := bluo.mutation.AddedBookID(); ok {
		_spec.AddField(booklog.FieldBookID, field.TypeInt, value)
	}
	if value, ok := bluo.mutation.BookTitle(); ok {
		_spec.SetField(booklog.FieldBookTitle, field.TypeString, value)
	}
	if value, ok := bluo.mutation.RequestID(); ok {
		_spec.SetField(booklog.FieldRequestID, field.TypeString, value)
	}
	if value, ok := bluo.mutation.CreatedAt(); ok {
		_spec.SetField(booklog.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &BookLog{config: bluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bluo.mutation.done = true
	return _node, nil
}