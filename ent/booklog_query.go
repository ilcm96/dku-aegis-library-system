// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ilcm96/dku-aegis-library/ent/booklog"
	"github.com/ilcm96/dku-aegis-library/ent/predicate"
)

// BookLogQuery is the builder for querying BookLog entities.
type BookLogQuery struct {
	config
	ctx        *QueryContext
	order      []booklog.OrderOption
	inters     []Interceptor
	predicates []predicate.BookLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookLogQuery builder.
func (blq *BookLogQuery) Where(ps ...predicate.BookLog) *BookLogQuery {
	blq.predicates = append(blq.predicates, ps...)
	return blq
}

// Limit the number of records to be returned by this query.
func (blq *BookLogQuery) Limit(limit int) *BookLogQuery {
	blq.ctx.Limit = &limit
	return blq
}

// Offset to start from.
func (blq *BookLogQuery) Offset(offset int) *BookLogQuery {
	blq.ctx.Offset = &offset
	return blq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (blq *BookLogQuery) Unique(unique bool) *BookLogQuery {
	blq.ctx.Unique = &unique
	return blq
}

// Order specifies how the records should be ordered.
func (blq *BookLogQuery) Order(o ...booklog.OrderOption) *BookLogQuery {
	blq.order = append(blq.order, o...)
	return blq
}

// First returns the first BookLog entity from the query.
// Returns a *NotFoundError when no BookLog was found.
func (blq *BookLogQuery) First(ctx context.Context) (*BookLog, error) {
	nodes, err := blq.Limit(1).All(setContextOp(ctx, blq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{booklog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (blq *BookLogQuery) FirstX(ctx context.Context) *BookLog {
	node, err := blq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BookLog ID from the query.
// Returns a *NotFoundError when no BookLog ID was found.
func (blq *BookLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = blq.Limit(1).IDs(setContextOp(ctx, blq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{booklog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (blq *BookLogQuery) FirstIDX(ctx context.Context) int {
	id, err := blq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BookLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BookLog entity is found.
// Returns a *NotFoundError when no BookLog entities are found.
func (blq *BookLogQuery) Only(ctx context.Context) (*BookLog, error) {
	nodes, err := blq.Limit(2).All(setContextOp(ctx, blq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{booklog.Label}
	default:
		return nil, &NotSingularError{booklog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (blq *BookLogQuery) OnlyX(ctx context.Context) *BookLog {
	node, err := blq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BookLog ID in the query.
// Returns a *NotSingularError when more than one BookLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (blq *BookLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = blq.Limit(2).IDs(setContextOp(ctx, blq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{booklog.Label}
	default:
		err = &NotSingularError{booklog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (blq *BookLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := blq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BookLogs.
func (blq *BookLogQuery) All(ctx context.Context) ([]*BookLog, error) {
	ctx = setContextOp(ctx, blq.ctx, "All")
	if err := blq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BookLog, *BookLogQuery]()
	return withInterceptors[[]*BookLog](ctx, blq, qr, blq.inters)
}

// AllX is like All, but panics if an error occurs.
func (blq *BookLogQuery) AllX(ctx context.Context) []*BookLog {
	nodes, err := blq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BookLog IDs.
func (blq *BookLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if blq.ctx.Unique == nil && blq.path != nil {
		blq.Unique(true)
	}
	ctx = setContextOp(ctx, blq.ctx, "IDs")
	if err = blq.Select(booklog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (blq *BookLogQuery) IDsX(ctx context.Context) []int {
	ids, err := blq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (blq *BookLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, blq.ctx, "Count")
	if err := blq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, blq, querierCount[*BookLogQuery](), blq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (blq *BookLogQuery) CountX(ctx context.Context) int {
	count, err := blq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (blq *BookLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, blq.ctx, "Exist")
	switch _, err := blq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (blq *BookLogQuery) ExistX(ctx context.Context) bool {
	exist, err := blq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (blq *BookLogQuery) Clone() *BookLogQuery {
	if blq == nil {
		return nil
	}
	return &BookLogQuery{
		config:     blq.config,
		ctx:        blq.ctx.Clone(),
		order:      append([]booklog.OrderOption{}, blq.order...),
		inters:     append([]Interceptor{}, blq.inters...),
		predicates: append([]predicate.BookLog{}, blq.predicates...),
		// clone intermediate query.
		sql:  blq.sql.Clone(),
		path: blq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Action booklog.Action `json:"action,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BookLog.Query().
//		GroupBy(booklog.FieldAction).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (blq *BookLogQuery) GroupBy(field string, fields ...string) *BookLogGroupBy {
	blq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BookLogGroupBy{build: blq}
	grbuild.flds = &blq.ctx.Fields
	grbuild.label = booklog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Action booklog.Action `json:"action,omitempty"`
//	}
//
//	client.BookLog.Query().
//		Select(booklog.FieldAction).
//		Scan(ctx, &v)
func (blq *BookLogQuery) Select(fields ...string) *BookLogSelect {
	blq.ctx.Fields = append(blq.ctx.Fields, fields...)
	sbuild := &BookLogSelect{BookLogQuery: blq}
	sbuild.label = booklog.Label
	sbuild.flds, sbuild.scan = &blq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookLogSelect configured with the given aggregations.
func (blq *BookLogQuery) Aggregate(fns ...AggregateFunc) *BookLogSelect {
	return blq.Select().Aggregate(fns...)
}

func (blq *BookLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range blq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, blq); err != nil {
				return err
			}
		}
	}
	for _, f := range blq.ctx.Fields {
		if !booklog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if blq.path != nil {
		prev, err := blq.path(ctx)
		if err != nil {
			return err
		}
		blq.sql = prev
	}
	return nil
}

func (blq *BookLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BookLog, error) {
	var (
		nodes = []*BookLog{}
		_spec = blq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BookLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BookLog{config: blq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, blq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (blq *BookLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := blq.querySpec()
	_spec.Node.Columns = blq.ctx.Fields
	if len(blq.ctx.Fields) > 0 {
		_spec.Unique = blq.ctx.Unique != nil && *blq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, blq.driver, _spec)
}

func (blq *BookLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(booklog.Table, booklog.Columns, sqlgraph.NewFieldSpec(booklog.FieldID, field.TypeInt))
	_spec.From = blq.sql
	if unique := blq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if blq.path != nil {
		_spec.Unique = true
	}
	if fields := blq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booklog.FieldID)
		for i := range fields {
			if fields[i] != booklog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := blq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := blq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := blq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := blq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (blq *BookLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(blq.driver.Dialect())
	t1 := builder.Table(booklog.Table)
	columns := blq.ctx.Fields
	if len(columns) == 0 {
		columns = booklog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if blq.sql != nil {
		selector = blq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if blq.ctx.Unique != nil && *blq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range blq.predicates {
		p(selector)
	}
	for _, p := range blq.order {
		p(selector)
	}
	if offset := blq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := blq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookLogGroupBy is the group-by builder for BookLog entities.
type BookLogGroupBy struct {
	selector
	build *BookLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (blgb *BookLogGroupBy) Aggregate(fns ...AggregateFunc) *BookLogGroupBy {
	blgb.fns = append(blgb.fns, fns...)
	return blgb
}

// Scan applies the selector query and scans the result into the given value.
func (blgb *BookLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, blgb.build.ctx, "GroupBy")
	if err := blgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookLogQuery, *BookLogGroupBy](ctx, blgb.build, blgb, blgb.build.inters, v)
}

func (blgb *BookLogGroupBy) sqlScan(ctx context.Context, root *BookLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(blgb.fns))
	for _, fn := range blgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*blgb.flds)+len(blgb.fns))
		for _, f := range *blgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*blgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := blgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookLogSelect is the builder for selecting fields of BookLog entities.
type BookLogSelect struct {
	*BookLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bls *BookLogSelect) Aggregate(fns ...AggregateFunc) *BookLogSelect {
	bls.fns = append(bls.fns, fns...)
	return bls
}

// Scan applies the selector query and scans the result into the given value.
func (bls *BookLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bls.ctx, "Select")
	if err := bls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookLogQuery, *BookLogSelect](ctx, bls.BookLogQuery, bls, bls.inters, v)
}

func (bls *BookLogSelect) sqlScan(ctx context.Context, root *BookLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bls.fns))
	for _, fn := range bls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
