// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/predicate"
	"go.infratographer.com/resource-provider-api/internal/ent/generated/resourceprovider"
	"go.infratographer.com/x/gidx"
)

// ResourceProviderQuery is the builder for querying ResourceProvider entities.
type ResourceProviderQuery struct {
	config
	ctx        *QueryContext
	order      []resourceprovider.OrderOption
	inters     []Interceptor
	predicates []predicate.ResourceProvider
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*ResourceProvider) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResourceProviderQuery builder.
func (rpq *ResourceProviderQuery) Where(ps ...predicate.ResourceProvider) *ResourceProviderQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit the number of records to be returned by this query.
func (rpq *ResourceProviderQuery) Limit(limit int) *ResourceProviderQuery {
	rpq.ctx.Limit = &limit
	return rpq
}

// Offset to start from.
func (rpq *ResourceProviderQuery) Offset(offset int) *ResourceProviderQuery {
	rpq.ctx.Offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *ResourceProviderQuery) Unique(unique bool) *ResourceProviderQuery {
	rpq.ctx.Unique = &unique
	return rpq
}

// Order specifies how the records should be ordered.
func (rpq *ResourceProviderQuery) Order(o ...resourceprovider.OrderOption) *ResourceProviderQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// First returns the first ResourceProvider entity from the query.
// Returns a *NotFoundError when no ResourceProvider was found.
func (rpq *ResourceProviderQuery) First(ctx context.Context) (*ResourceProvider, error) {
	nodes, err := rpq.Limit(1).All(setContextOp(ctx, rpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resourceprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *ResourceProviderQuery) FirstX(ctx context.Context) *ResourceProvider {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ResourceProvider ID from the query.
// Returns a *NotFoundError when no ResourceProvider ID was found.
func (rpq *ResourceProviderQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = rpq.Limit(1).IDs(setContextOp(ctx, rpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resourceprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *ResourceProviderQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ResourceProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ResourceProvider entity is found.
// Returns a *NotFoundError when no ResourceProvider entities are found.
func (rpq *ResourceProviderQuery) Only(ctx context.Context) (*ResourceProvider, error) {
	nodes, err := rpq.Limit(2).All(setContextOp(ctx, rpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resourceprovider.Label}
	default:
		return nil, &NotSingularError{resourceprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *ResourceProviderQuery) OnlyX(ctx context.Context) *ResourceProvider {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ResourceProvider ID in the query.
// Returns a *NotSingularError when more than one ResourceProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (rpq *ResourceProviderQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = rpq.Limit(2).IDs(setContextOp(ctx, rpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resourceprovider.Label}
	default:
		err = &NotSingularError{resourceprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *ResourceProviderQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ResourceProviders.
func (rpq *ResourceProviderQuery) All(ctx context.Context) ([]*ResourceProvider, error) {
	ctx = setContextOp(ctx, rpq.ctx, ent.OpQueryAll)
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ResourceProvider, *ResourceProviderQuery]()
	return withInterceptors[[]*ResourceProvider](ctx, rpq, qr, rpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rpq *ResourceProviderQuery) AllX(ctx context.Context) []*ResourceProvider {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ResourceProvider IDs.
func (rpq *ResourceProviderQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if rpq.ctx.Unique == nil && rpq.path != nil {
		rpq.Unique(true)
	}
	ctx = setContextOp(ctx, rpq.ctx, ent.OpQueryIDs)
	if err = rpq.Select(resourceprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *ResourceProviderQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *ResourceProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rpq.ctx, ent.OpQueryCount)
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rpq, querierCount[*ResourceProviderQuery](), rpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *ResourceProviderQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *ResourceProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rpq.ctx, ent.OpQueryExist)
	switch _, err := rpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *ResourceProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResourceProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *ResourceProviderQuery) Clone() *ResourceProviderQuery {
	if rpq == nil {
		return nil
	}
	return &ResourceProviderQuery{
		config:     rpq.config,
		ctx:        rpq.ctx.Clone(),
		order:      append([]resourceprovider.OrderOption{}, rpq.order...),
		inters:     append([]Interceptor{}, rpq.inters...),
		predicates: append([]predicate.ResourceProvider{}, rpq.predicates...),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ResourceProvider.Query().
//		GroupBy(resourceprovider.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (rpq *ResourceProviderQuery) GroupBy(field string, fields ...string) *ResourceProviderGroupBy {
	rpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ResourceProviderGroupBy{build: rpq}
	grbuild.flds = &rpq.ctx.Fields
	grbuild.label = resourceprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ResourceProvider.Query().
//		Select(resourceprovider.FieldCreatedAt).
//		Scan(ctx, &v)
func (rpq *ResourceProviderQuery) Select(fields ...string) *ResourceProviderSelect {
	rpq.ctx.Fields = append(rpq.ctx.Fields, fields...)
	sbuild := &ResourceProviderSelect{ResourceProviderQuery: rpq}
	sbuild.label = resourceprovider.Label
	sbuild.flds, sbuild.scan = &rpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ResourceProviderSelect configured with the given aggregations.
func (rpq *ResourceProviderQuery) Aggregate(fns ...AggregateFunc) *ResourceProviderSelect {
	return rpq.Select().Aggregate(fns...)
}

func (rpq *ResourceProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rpq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rpq); err != nil {
				return err
			}
		}
	}
	for _, f := range rpq.ctx.Fields {
		if !resourceprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *ResourceProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ResourceProvider, error) {
	var (
		nodes = []*ResourceProvider{}
		_spec = rpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ResourceProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ResourceProvider{config: rpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(rpq.modifiers) > 0 {
		_spec.Modifiers = rpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range rpq.loadTotal {
		if err := rpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rpq *ResourceProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	if len(rpq.modifiers) > 0 {
		_spec.Modifiers = rpq.modifiers
	}
	_spec.Node.Columns = rpq.ctx.Fields
	if len(rpq.ctx.Fields) > 0 {
		_spec.Unique = rpq.ctx.Unique != nil && *rpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *ResourceProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(resourceprovider.Table, resourceprovider.Columns, sqlgraph.NewFieldSpec(resourceprovider.FieldID, field.TypeString))
	_spec.From = rpq.sql
	if unique := rpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rpq.path != nil {
		_spec.Unique = true
	}
	if fields := rpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourceprovider.FieldID)
		for i := range fields {
			if fields[i] != resourceprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *ResourceProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(resourceprovider.Table)
	columns := rpq.ctx.Fields
	if len(columns) == 0 {
		columns = resourceprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rpq.ctx.Unique != nil && *rpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResourceProviderGroupBy is the group-by builder for ResourceProvider entities.
type ResourceProviderGroupBy struct {
	selector
	build *ResourceProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *ResourceProviderGroupBy) Aggregate(fns ...AggregateFunc) *ResourceProviderGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the selector query and scans the result into the given value.
func (rpgb *ResourceProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rpgb.build.ctx, ent.OpQueryGroupBy)
	if err := rpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResourceProviderQuery, *ResourceProviderGroupBy](ctx, rpgb.build, rpgb, rpgb.build.inters, v)
}

func (rpgb *ResourceProviderGroupBy) sqlScan(ctx context.Context, root *ResourceProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rpgb.fns))
	for _, fn := range rpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rpgb.flds)+len(rpgb.fns))
		for _, f := range *rpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ResourceProviderSelect is the builder for selecting fields of ResourceProvider entities.
type ResourceProviderSelect struct {
	*ResourceProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rps *ResourceProviderSelect) Aggregate(fns ...AggregateFunc) *ResourceProviderSelect {
	rps.fns = append(rps.fns, fns...)
	return rps
}

// Scan applies the selector query and scans the result into the given value.
func (rps *ResourceProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rps.ctx, ent.OpQuerySelect)
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResourceProviderQuery, *ResourceProviderSelect](ctx, rps.ResourceProviderQuery, rps, rps.inters, v)
}

func (rps *ResourceProviderSelect) sqlScan(ctx context.Context, root *ResourceProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rps.fns))
	for _, fn := range rps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
