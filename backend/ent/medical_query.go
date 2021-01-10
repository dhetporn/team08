// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/medical"
	"github.com/sut63/team08/ent/predicate"
)

// MedicalQuery is the builder for querying Medical entities.
type MedicalQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Medical
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mq *MedicalQuery) Where(ps ...predicate.Medical) *MedicalQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MedicalQuery) Limit(limit int) *MedicalQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MedicalQuery) Offset(offset int) *MedicalQuery {
	mq.offset = &offset
	return mq
}

// Order adds an order step to the query.
func (mq *MedicalQuery) Order(o ...OrderFunc) *MedicalQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// First returns the first Medical entity in the query. Returns *NotFoundError when no medical was found.
func (mq *MedicalQuery) First(ctx context.Context) (*Medical, error) {
	ms, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ms) == 0 {
		return nil, &NotFoundError{medical.Label}
	}
	return ms[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MedicalQuery) FirstX(ctx context.Context) *Medical {
	m, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return m
}

// FirstID returns the first Medical id in the query. Returns *NotFoundError when no id was found.
func (mq *MedicalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{medical.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mq *MedicalQuery) FirstXID(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Medical entity in the query, returns an error if not exactly one entity was returned.
func (mq *MedicalQuery) Only(ctx context.Context) (*Medical, error) {
	ms, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ms) {
	case 1:
		return ms[0], nil
	case 0:
		return nil, &NotFoundError{medical.Label}
	default:
		return nil, &NotSingularError{medical.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MedicalQuery) OnlyX(ctx context.Context) *Medical {
	m, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// OnlyID returns the only Medical id in the query, returns an error if not exactly one id was returned.
func (mq *MedicalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = &NotSingularError{medical.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MedicalQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Medicals.
func (mq *MedicalQuery) All(ctx context.Context) ([]*Medical, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MedicalQuery) AllX(ctx context.Context) []*Medical {
	ms, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ms
}

// IDs executes the query and returns a list of Medical ids.
func (mq *MedicalQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(medical.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MedicalQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MedicalQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MedicalQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MedicalQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MedicalQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MedicalQuery) Clone() *MedicalQuery {
	return &MedicalQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]OrderFunc{}, mq.order...),
		unique:     append([]string{}, mq.unique...),
		predicates: append([]predicate.Medical{}, mq.predicates...),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MedicalName string `json:"Medical_Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Medical.Query().
//		GroupBy(medical.FieldMedicalName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MedicalQuery) GroupBy(field string, fields ...string) *MedicalGroupBy {
	group := &MedicalGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		MedicalName string `json:"Medical_Name,omitempty"`
//	}
//
//	client.Medical.Query().
//		Select(medical.FieldMedicalName).
//		Scan(ctx, &v)
//
func (mq *MedicalQuery) Select(field string, fields ...string) *MedicalSelect {
	selector := &MedicalSelect{config: mq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return selector
}

func (mq *MedicalQuery) prepareQuery(ctx context.Context) error {
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MedicalQuery) sqlAll(ctx context.Context) ([]*Medical, error) {
	var (
		nodes = []*Medical{}
		_spec = mq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &Medical{config: mq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mq *MedicalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MedicalQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mq *MedicalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medical.Table,
			Columns: medical.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medical.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MedicalQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(medical.Table)
	selector := builder.Select(t1.Columns(medical.Columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(medical.Columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MedicalGroupBy is the builder for group-by Medical entities.
type MedicalGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MedicalGroupBy) Aggregate(fns ...AggregateFunc) *MedicalGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mgb *MedicalGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MedicalGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MedicalGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *MedicalGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MedicalGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *MedicalGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MedicalGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *MedicalGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MedicalGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *MedicalGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MedicalGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgb.sqlQuery().Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MedicalGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql
	columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
	columns = append(columns, mgb.fields...)
	for _, fn := range mgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mgb.fields...)
}

// MedicalSelect is the builder for select fields of Medical entities.
type MedicalSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ms *MedicalSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ms.path(ctx)
	if err != nil {
		return err
	}
	ms.sql = query
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MedicalSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MedicalSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *MedicalSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MedicalSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *MedicalSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MedicalSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *MedicalSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MedicalSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ms *MedicalSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medical.Label}
	default:
		err = fmt.Errorf("ent: MedicalSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *MedicalSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MedicalSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sqlQuery().Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ms *MedicalSelect) sqlQuery() sql.Querier {
	selector := ms.sql
	selector.Select(selector.Columns(ms.fields...)...)
	return selector
}
