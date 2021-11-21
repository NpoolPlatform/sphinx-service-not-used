// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/keystore"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// KeyStoreQuery is the builder for querying KeyStore entities.
type KeyStoreQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.KeyStore
	// eager-loading edges.
	withCoin *CoinInfoQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeyStoreQuery builder.
func (ksq *KeyStoreQuery) Where(ps ...predicate.KeyStore) *KeyStoreQuery {
	ksq.predicates = append(ksq.predicates, ps...)
	return ksq
}

// Limit adds a limit step to the query.
func (ksq *KeyStoreQuery) Limit(limit int) *KeyStoreQuery {
	ksq.limit = &limit
	return ksq
}

// Offset adds an offset step to the query.
func (ksq *KeyStoreQuery) Offset(offset int) *KeyStoreQuery {
	ksq.offset = &offset
	return ksq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ksq *KeyStoreQuery) Unique(unique bool) *KeyStoreQuery {
	ksq.unique = &unique
	return ksq
}

// Order adds an order step to the query.
func (ksq *KeyStoreQuery) Order(o ...OrderFunc) *KeyStoreQuery {
	ksq.order = append(ksq.order, o...)
	return ksq
}

// QueryCoin chains the current query on the "coin" edge.
func (ksq *KeyStoreQuery) QueryCoin() *CoinInfoQuery {
	query := &CoinInfoQuery{config: ksq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ksq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(keystore.Table, keystore.FieldID, selector),
			sqlgraph.To(coininfo.Table, coininfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, keystore.CoinTable, keystore.CoinColumn),
		)
		fromU = sqlgraph.SetNeighbors(ksq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KeyStore entity from the query.
// Returns a *NotFoundError when no KeyStore was found.
func (ksq *KeyStoreQuery) First(ctx context.Context) (*KeyStore, error) {
	nodes, err := ksq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keystore.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ksq *KeyStoreQuery) FirstX(ctx context.Context) *KeyStore {
	node, err := ksq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KeyStore ID from the query.
// Returns a *NotFoundError when no KeyStore ID was found.
func (ksq *KeyStoreQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ksq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keystore.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ksq *KeyStoreQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ksq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KeyStore entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one KeyStore entity is not found.
// Returns a *NotFoundError when no KeyStore entities are found.
func (ksq *KeyStoreQuery) Only(ctx context.Context) (*KeyStore, error) {
	nodes, err := ksq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keystore.Label}
	default:
		return nil, &NotSingularError{keystore.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ksq *KeyStoreQuery) OnlyX(ctx context.Context) *KeyStore {
	node, err := ksq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KeyStore ID in the query.
// Returns a *NotSingularError when exactly one KeyStore ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ksq *KeyStoreQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ksq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = &NotSingularError{keystore.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ksq *KeyStoreQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ksq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KeyStores.
func (ksq *KeyStoreQuery) All(ctx context.Context) ([]*KeyStore, error) {
	if err := ksq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ksq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ksq *KeyStoreQuery) AllX(ctx context.Context) []*KeyStore {
	nodes, err := ksq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KeyStore IDs.
func (ksq *KeyStoreQuery) IDs(ctx context.Context) ([]int32, error) {
	var ids []int32
	if err := ksq.Select(keystore.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ksq *KeyStoreQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ksq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ksq *KeyStoreQuery) Count(ctx context.Context) (int, error) {
	if err := ksq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ksq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ksq *KeyStoreQuery) CountX(ctx context.Context) int {
	count, err := ksq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ksq *KeyStoreQuery) Exist(ctx context.Context) (bool, error) {
	if err := ksq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ksq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ksq *KeyStoreQuery) ExistX(ctx context.Context) bool {
	exist, err := ksq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeyStoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ksq *KeyStoreQuery) Clone() *KeyStoreQuery {
	if ksq == nil {
		return nil
	}
	return &KeyStoreQuery{
		config:     ksq.config,
		limit:      ksq.limit,
		offset:     ksq.offset,
		order:      append([]OrderFunc{}, ksq.order...),
		predicates: append([]predicate.KeyStore{}, ksq.predicates...),
		withCoin:   ksq.withCoin.Clone(),
		// clone intermediate query.
		sql:  ksq.sql.Clone(),
		path: ksq.path,
	}
}

// WithCoin tells the query-builder to eager-load the nodes that are connected to
// the "coin" edge. The optional arguments are used to configure the query builder of the edge.
func (ksq *KeyStoreQuery) WithCoin(opts ...func(*CoinInfoQuery)) *KeyStoreQuery {
	query := &CoinInfoQuery{config: ksq.config}
	for _, opt := range opts {
		opt(query)
	}
	ksq.withCoin = query
	return ksq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Address string `json:"address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.KeyStore.Query().
//		GroupBy(keystore.FieldAddress).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ksq *KeyStoreQuery) GroupBy(field string, fields ...string) *KeyStoreGroupBy {
	group := &KeyStoreGroupBy{config: ksq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ksq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Address string `json:"address,omitempty"`
//	}
//
//	client.KeyStore.Query().
//		Select(keystore.FieldAddress).
//		Scan(ctx, &v)
//
func (ksq *KeyStoreQuery) Select(fields ...string) *KeyStoreSelect {
	ksq.fields = append(ksq.fields, fields...)
	return &KeyStoreSelect{KeyStoreQuery: ksq}
}

func (ksq *KeyStoreQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ksq.fields {
		if !keystore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ksq.path != nil {
		prev, err := ksq.path(ctx)
		if err != nil {
			return err
		}
		ksq.sql = prev
	}
	return nil
}

func (ksq *KeyStoreQuery) sqlAll(ctx context.Context) ([]*KeyStore, error) {
	var (
		nodes       = []*KeyStore{}
		withFKs     = ksq.withFKs
		_spec       = ksq.querySpec()
		loadedTypes = [1]bool{
			ksq.withCoin != nil,
		}
	)
	if ksq.withCoin != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, keystore.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &KeyStore{config: ksq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ksq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ksq.withCoin; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*KeyStore)
		for i := range nodes {
			if nodes[i].coin_info_keys == nil {
				continue
			}
			fk := *nodes[i].coin_info_keys
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(coininfo.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coin_info_keys" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Coin = n
			}
		}
	}

	return nodes, nil
}

func (ksq *KeyStoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ksq.querySpec()
	return sqlgraph.CountNodes(ctx, ksq.driver, _spec)
}

func (ksq *KeyStoreQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ksq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ksq *KeyStoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keystore.Table,
			Columns: keystore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: keystore.FieldID,
			},
		},
		From:   ksq.sql,
		Unique: true,
	}
	if unique := ksq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ksq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keystore.FieldID)
		for i := range fields {
			if fields[i] != keystore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ksq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ksq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ksq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ksq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ksq *KeyStoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ksq.driver.Dialect())
	t1 := builder.Table(keystore.Table)
	columns := ksq.fields
	if len(columns) == 0 {
		columns = keystore.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ksq.sql != nil {
		selector = ksq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ksq.predicates {
		p(selector)
	}
	for _, p := range ksq.order {
		p(selector)
	}
	if offset := ksq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ksq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeyStoreGroupBy is the group-by builder for KeyStore entities.
type KeyStoreGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ksgb *KeyStoreGroupBy) Aggregate(fns ...AggregateFunc) *KeyStoreGroupBy {
	ksgb.fns = append(ksgb.fns, fns...)
	return ksgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ksgb *KeyStoreGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ksgb.path(ctx)
	if err != nil {
		return err
	}
	ksgb.sql = query
	return ksgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ksgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KeyStoreGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) StringsX(ctx context.Context) []string {
	v, err := ksgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ksgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) StringX(ctx context.Context) string {
	v, err := ksgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KeyStoreGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) IntsX(ctx context.Context) []int {
	v, err := ksgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ksgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) IntX(ctx context.Context) int {
	v, err := ksgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KeyStoreGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ksgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ksgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ksgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ksgb.fields) > 1 {
		return nil, errors.New("ent: KeyStoreGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ksgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ksgb *KeyStoreGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ksgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ksgb *KeyStoreGroupBy) BoolX(ctx context.Context) bool {
	v, err := ksgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ksgb *KeyStoreGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ksgb.fields {
		if !keystore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ksgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ksgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ksgb *KeyStoreGroupBy) sqlQuery() *sql.Selector {
	selector := ksgb.sql.Select()
	aggregation := make([]string, 0, len(ksgb.fns))
	for _, fn := range ksgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ksgb.fields)+len(ksgb.fns))
		for _, f := range ksgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ksgb.fields...)...)
}

// KeyStoreSelect is the builder for selecting fields of KeyStore entities.
type KeyStoreSelect struct {
	*KeyStoreQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (kss *KeyStoreSelect) Scan(ctx context.Context, v interface{}) error {
	if err := kss.prepareQuery(ctx); err != nil {
		return err
	}
	kss.sql = kss.KeyStoreQuery.sqlQuery(ctx)
	return kss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kss *KeyStoreSelect) ScanX(ctx context.Context, v interface{}) {
	if err := kss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Strings(ctx context.Context) ([]string, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KeyStoreSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kss *KeyStoreSelect) StringsX(ctx context.Context) []string {
	v, err := kss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kss *KeyStoreSelect) StringX(ctx context.Context) string {
	v, err := kss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Ints(ctx context.Context) ([]int, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KeyStoreSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kss *KeyStoreSelect) IntsX(ctx context.Context) []int {
	v, err := kss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kss *KeyStoreSelect) IntX(ctx context.Context) int {
	v, err := kss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KeyStoreSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kss *KeyStoreSelect) Float64sX(ctx context.Context) []float64 {
	v, err := kss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kss *KeyStoreSelect) Float64X(ctx context.Context) float64 {
	v, err := kss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(kss.fields) > 1 {
		return nil, errors.New("ent: KeyStoreSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := kss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kss *KeyStoreSelect) BoolsX(ctx context.Context) []bool {
	v, err := kss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (kss *KeyStoreSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keystore.Label}
	default:
		err = fmt.Errorf("ent: KeyStoreSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kss *KeyStoreSelect) BoolX(ctx context.Context) bool {
	v, err := kss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kss *KeyStoreSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := kss.sql.Query()
	if err := kss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
