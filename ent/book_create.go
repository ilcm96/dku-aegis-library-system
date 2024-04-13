// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ilcm96/dku-aegis-library/ent/book"
	"github.com/ilcm96/dku-aegis-library/ent/category"
	"github.com/ilcm96/dku-aegis-library/ent/user"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (bc *BookCreate) SetTitle(s string) *BookCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetAuthor sets the "author" field.
func (bc *BookCreate) SetAuthor(s string) *BookCreate {
	bc.mutation.SetAuthor(s)
	return bc
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (bc *BookCreate) SetNillableAuthor(s *string) *BookCreate {
	if s != nil {
		bc.SetAuthor(*s)
	}
	return bc
}

// SetPublisher sets the "publisher" field.
func (bc *BookCreate) SetPublisher(s string) *BookCreate {
	bc.mutation.SetPublisher(s)
	return bc
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (bc *BookCreate) SetNillablePublisher(s *string) *BookCreate {
	if s != nil {
		bc.SetPublisher(*s)
	}
	return bc
}

// SetQuantity sets the "quantity" field.
func (bc *BookCreate) SetQuantity(i int) *BookCreate {
	bc.mutation.SetQuantity(i)
	return bc
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (bc *BookCreate) SetNillableQuantity(i *int) *BookCreate {
	if i != nil {
		bc.SetQuantity(*i)
	}
	return bc
}

// SetBorrow sets the "borrow" field.
func (bc *BookCreate) SetBorrow(i int) *BookCreate {
	bc.mutation.SetBorrow(i)
	return bc
}

// SetNillableBorrow sets the "borrow" field if the given value is not nil.
func (bc *BookCreate) SetNillableBorrow(i *int) *BookCreate {
	if i != nil {
		bc.SetBorrow(*i)
	}
	return bc
}

// SetCover sets the "cover" field.
func (bc *BookCreate) SetCover(s string) *BookCreate {
	bc.mutation.SetCover(s)
	return bc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (bc *BookCreate) AddUserIDs(ids ...int) *BookCreate {
	bc.mutation.AddUserIDs(ids...)
	return bc
}

// AddUser adds the "user" edges to the User entity.
func (bc *BookCreate) AddUser(u ...*User) *BookCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddUserIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (bc *BookCreate) AddCategoryIDs(ids ...int) *BookCreate {
	bc.mutation.AddCategoryIDs(ids...)
	return bc
}

// AddCategory adds the "category" edges to the Category entity.
func (bc *BookCreate) AddCategory(c ...*Category) *BookCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCategoryIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookCreate) defaults() {
	if _, ok := bc.mutation.Quantity(); !ok {
		v := book.DefaultQuantity
		bc.mutation.SetQuantity(v)
	}
	if _, ok := bc.mutation.Borrow(); !ok {
		v := book.DefaultBorrow
		bc.mutation.SetBorrow(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Book.title"`)}
	}
	if v, ok := bc.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Book.quantity"`)}
	}
	if _, ok := bc.mutation.Borrow(); !ok {
		return &ValidationError{Name: "borrow", err: errors.New(`ent: missing required field "Book.borrow"`)}
	}
	if _, ok := bc.mutation.Cover(); !ok {
		return &ValidationError{Name: "cover", err: errors.New(`ent: missing required field "Book.cover"`)}
	}
	if len(bc.mutation.CategoryIDs()) == 0 {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required edge "Book.category"`)}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(book.Table, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := bc.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
		_node.Publisher = value
	}
	if value, ok := bc.mutation.Quantity(); ok {
		_spec.SetField(book.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := bc.mutation.Borrow(); ok {
		_spec.SetField(book.FieldBorrow, field.TypeInt, value)
		_node.Borrow = value
	}
	if value, ok := bc.mutation.Cover(); ok {
		_spec.SetField(book.FieldCover, field.TypeString, value)
		_node.Cover = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.UserTable,
			Columns: book.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.CategoryTable,
			Columns: book.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	err      error
	builders []*BookCreate
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
