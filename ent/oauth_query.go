// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"enttest/ent/oauth"
	"enttest/ent/predicate"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OAuthQuery is the builder for querying OAuth entities.
type OAuthQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OAuth
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OAuthQuery builder.
func (oq *OAuthQuery) Where(ps ...predicate.OAuth) *OAuthQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OAuthQuery) Limit(limit int) *OAuthQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OAuthQuery) Offset(offset int) *OAuthQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OAuthQuery) Unique(unique bool) *OAuthQuery {
	oq.unique = &unique
	return oq
}

// Order adds an order step to the query.
func (oq *OAuthQuery) Order(o ...OrderFunc) *OAuthQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// First returns the first OAuth entity from the query.
// Returns a *NotFoundError when no OAuth was found.
func (oq *OAuthQuery) First(ctx context.Context) (*OAuth, error) {
	nodes, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauth.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OAuthQuery) FirstX(ctx context.Context) *OAuth {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OAuth ID from the query.
// Returns a *NotFoundError when no OAuth ID was found.
func (oq *OAuthQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauth.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OAuthQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OAuth entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OAuth entity is not found.
// Returns a *NotFoundError when no OAuth entities are found.
func (oq *OAuthQuery) Only(ctx context.Context) (*OAuth, error) {
	nodes, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauth.Label}
	default:
		return nil, &NotSingularError{oauth.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OAuthQuery) OnlyX(ctx context.Context) *OAuth {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OAuth ID in the query.
// Returns a *NotSingularError when exactly one OAuth ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oq *OAuthQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = &NotSingularError{oauth.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OAuthQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OAuths.
func (oq *OAuthQuery) All(ctx context.Context) ([]*OAuth, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *OAuthQuery) AllX(ctx context.Context) []*OAuth {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OAuth IDs.
func (oq *OAuthQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(oauth.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OAuthQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OAuthQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OAuthQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OAuthQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OAuthQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OAuthQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OAuthQuery) Clone() *OAuthQuery {
	if oq == nil {
		return nil
	}
	return &OAuthQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		predicates: append([]predicate.OAuth{}, oq.predicates...),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OAuth.Query().
//		GroupBy(oauth.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oq *OAuthQuery) GroupBy(field string, fields ...string) *OAuthGroupBy {
	group := &OAuthGroupBy{config: oq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.OAuth.Query().
//		Select(oauth.FieldName).
//		Scan(ctx, &v)
//
func (oq *OAuthQuery) Select(fields ...string) *OAuthSelect {
	oq.fields = append(oq.fields, fields...)
	return &OAuthSelect{OAuthQuery: oq}
}

func (oq *OAuthQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oq.fields {
		if !oauth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OAuthQuery) sqlAll(ctx context.Context) ([]*OAuth, error) {
	var (
		nodes = []*OAuth{}
		_spec = oq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OAuth{config: oq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oq *OAuthQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OAuthQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oq *OAuthQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauth.Table,
			Columns: oauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth.FieldID)
		for i := range fields {
			if fields[i] != oauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OAuthQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(oauth.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = oauth.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OAuthGroupBy is the group-by builder for OAuth entities.
type OAuthGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OAuthGroupBy) Aggregate(fns ...AggregateFunc) *OAuthGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scans the result into the given value.
func (ogb *OAuthGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogb *OAuthGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OAuthGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogb *OAuthGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogb *OAuthGroupBy) StringX(ctx context.Context) string {
	v, err := ogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OAuthGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogb *OAuthGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogb *OAuthGroupBy) IntX(ctx context.Context) int {
	v, err := ogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OAuthGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogb *OAuthGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogb *OAuthGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OAuthGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogb *OAuthGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OAuthGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogb *OAuthGroupBy) BoolX(ctx context.Context) bool {
	v, err := ogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogb *OAuthGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ogb.fields {
		if !oauth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OAuthGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql.Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
		for _, f := range ogb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ogb.fields...)...)
}

// OAuthSelect is the builder for selecting fields of OAuth entities.
type OAuthSelect struct {
	*OAuthQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (os *OAuthSelect) Scan(ctx context.Context, v interface{}) error {
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	os.sql = os.OAuthQuery.sqlQuery(ctx)
	return os.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (os *OAuthSelect) ScanX(ctx context.Context, v interface{}) {
	if err := os.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Strings(ctx context.Context) ([]string, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OAuthSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (os *OAuthSelect) StringsX(ctx context.Context) []string {
	v, err := os.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = os.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (os *OAuthSelect) StringX(ctx context.Context) string {
	v, err := os.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Ints(ctx context.Context) ([]int, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OAuthSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (os *OAuthSelect) IntsX(ctx context.Context) []int {
	v, err := os.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = os.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (os *OAuthSelect) IntX(ctx context.Context) int {
	v, err := os.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OAuthSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (os *OAuthSelect) Float64sX(ctx context.Context) []float64 {
	v, err := os.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = os.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (os *OAuthSelect) Float64X(ctx context.Context) float64 {
	v, err := os.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OAuthSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (os *OAuthSelect) BoolsX(ctx context.Context) []bool {
	v, err := os.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (os *OAuthSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = os.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth.Label}
	default:
		err = fmt.Errorf("ent: OAuthSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (os *OAuthSelect) BoolX(ctx context.Context) bool {
	v, err := os.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (os *OAuthSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := os.sql.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
