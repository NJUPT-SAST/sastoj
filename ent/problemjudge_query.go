// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"sastoj/ent/group"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"
	"sastoj/ent/problemjudge"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemJudgeQuery is the builder for querying ProblemJudge entities.
type ProblemJudgeQuery struct {
	config
	ctx          *QueryContext
	order        []problemjudge.OrderOption
	inters       []Interceptor
	predicates   []predicate.ProblemJudge
	withGroups   *GroupQuery
	withProblems *ProblemQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProblemJudgeQuery builder.
func (pjq *ProblemJudgeQuery) Where(ps ...predicate.ProblemJudge) *ProblemJudgeQuery {
	pjq.predicates = append(pjq.predicates, ps...)
	return pjq
}

// Limit the number of records to be returned by this query.
func (pjq *ProblemJudgeQuery) Limit(limit int) *ProblemJudgeQuery {
	pjq.ctx.Limit = &limit
	return pjq
}

// Offset to start from.
func (pjq *ProblemJudgeQuery) Offset(offset int) *ProblemJudgeQuery {
	pjq.ctx.Offset = &offset
	return pjq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pjq *ProblemJudgeQuery) Unique(unique bool) *ProblemJudgeQuery {
	pjq.ctx.Unique = &unique
	return pjq
}

// Order specifies how the records should be ordered.
func (pjq *ProblemJudgeQuery) Order(o ...problemjudge.OrderOption) *ProblemJudgeQuery {
	pjq.order = append(pjq.order, o...)
	return pjq
}

// QueryGroups chains the current query on the "groups" edge.
func (pjq *ProblemJudgeQuery) QueryGroups() *GroupQuery {
	query := (&GroupClient{config: pjq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pjq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problemjudge.Table, problemjudge.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problemjudge.GroupsTable, problemjudge.GroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pjq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProblems chains the current query on the "problems" edge.
func (pjq *ProblemJudgeQuery) QueryProblems() *ProblemQuery {
	query := (&ProblemClient{config: pjq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pjq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problemjudge.Table, problemjudge.FieldID, selector),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problemjudge.ProblemsTable, problemjudge.ProblemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pjq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProblemJudge entity from the query.
// Returns a *NotFoundError when no ProblemJudge was found.
func (pjq *ProblemJudgeQuery) First(ctx context.Context) (*ProblemJudge, error) {
	nodes, err := pjq.Limit(1).All(setContextOp(ctx, pjq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{problemjudge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) FirstX(ctx context.Context) *ProblemJudge {
	node, err := pjq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProblemJudge ID from the query.
// Returns a *NotFoundError when no ProblemJudge ID was found.
func (pjq *ProblemJudgeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pjq.Limit(1).IDs(setContextOp(ctx, pjq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{problemjudge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) FirstIDX(ctx context.Context) int {
	id, err := pjq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProblemJudge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProblemJudge entity is found.
// Returns a *NotFoundError when no ProblemJudge entities are found.
func (pjq *ProblemJudgeQuery) Only(ctx context.Context) (*ProblemJudge, error) {
	nodes, err := pjq.Limit(2).All(setContextOp(ctx, pjq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{problemjudge.Label}
	default:
		return nil, &NotSingularError{problemjudge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) OnlyX(ctx context.Context) *ProblemJudge {
	node, err := pjq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProblemJudge ID in the query.
// Returns a *NotSingularError when more than one ProblemJudge ID is found.
// Returns a *NotFoundError when no entities are found.
func (pjq *ProblemJudgeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pjq.Limit(2).IDs(setContextOp(ctx, pjq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{problemjudge.Label}
	default:
		err = &NotSingularError{problemjudge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pjq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProblemJudges.
func (pjq *ProblemJudgeQuery) All(ctx context.Context) ([]*ProblemJudge, error) {
	ctx = setContextOp(ctx, pjq.ctx, "All")
	if err := pjq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProblemJudge, *ProblemJudgeQuery]()
	return withInterceptors[[]*ProblemJudge](ctx, pjq, qr, pjq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) AllX(ctx context.Context) []*ProblemJudge {
	nodes, err := pjq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProblemJudge IDs.
func (pjq *ProblemJudgeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pjq.ctx.Unique == nil && pjq.path != nil {
		pjq.Unique(true)
	}
	ctx = setContextOp(ctx, pjq.ctx, "IDs")
	if err = pjq.Select(problemjudge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) IDsX(ctx context.Context) []int {
	ids, err := pjq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pjq *ProblemJudgeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pjq.ctx, "Count")
	if err := pjq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pjq, querierCount[*ProblemJudgeQuery](), pjq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) CountX(ctx context.Context) int {
	count, err := pjq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pjq *ProblemJudgeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pjq.ctx, "Exist")
	switch _, err := pjq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pjq *ProblemJudgeQuery) ExistX(ctx context.Context) bool {
	exist, err := pjq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProblemJudgeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pjq *ProblemJudgeQuery) Clone() *ProblemJudgeQuery {
	if pjq == nil {
		return nil
	}
	return &ProblemJudgeQuery{
		config:       pjq.config,
		ctx:          pjq.ctx.Clone(),
		order:        append([]problemjudge.OrderOption{}, pjq.order...),
		inters:       append([]Interceptor{}, pjq.inters...),
		predicates:   append([]predicate.ProblemJudge{}, pjq.predicates...),
		withGroups:   pjq.withGroups.Clone(),
		withProblems: pjq.withProblems.Clone(),
		// clone intermediate query.
		sql:  pjq.sql.Clone(),
		path: pjq.path,
	}
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (pjq *ProblemJudgeQuery) WithGroups(opts ...func(*GroupQuery)) *ProblemJudgeQuery {
	query := (&GroupClient{config: pjq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pjq.withGroups = query
	return pjq
}

// WithProblems tells the query-builder to eager-load the nodes that are connected to
// the "problems" edge. The optional arguments are used to configure the query builder of the edge.
func (pjq *ProblemJudgeQuery) WithProblems(opts ...func(*ProblemQuery)) *ProblemJudgeQuery {
	query := (&ProblemClient{config: pjq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pjq.withProblems = query
	return pjq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GroupID int `json:"group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProblemJudge.Query().
//		GroupBy(problemjudge.FieldGroupID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pjq *ProblemJudgeQuery) GroupBy(field string, fields ...string) *ProblemJudgeGroupBy {
	pjq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProblemJudgeGroupBy{build: pjq}
	grbuild.flds = &pjq.ctx.Fields
	grbuild.label = problemjudge.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GroupID int `json:"group_id,omitempty"`
//	}
//
//	client.ProblemJudge.Query().
//		Select(problemjudge.FieldGroupID).
//		Scan(ctx, &v)
func (pjq *ProblemJudgeQuery) Select(fields ...string) *ProblemJudgeSelect {
	pjq.ctx.Fields = append(pjq.ctx.Fields, fields...)
	sbuild := &ProblemJudgeSelect{ProblemJudgeQuery: pjq}
	sbuild.label = problemjudge.Label
	sbuild.flds, sbuild.scan = &pjq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProblemJudgeSelect configured with the given aggregations.
func (pjq *ProblemJudgeQuery) Aggregate(fns ...AggregateFunc) *ProblemJudgeSelect {
	return pjq.Select().Aggregate(fns...)
}

func (pjq *ProblemJudgeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pjq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pjq); err != nil {
				return err
			}
		}
	}
	for _, f := range pjq.ctx.Fields {
		if !problemjudge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pjq.path != nil {
		prev, err := pjq.path(ctx)
		if err != nil {
			return err
		}
		pjq.sql = prev
	}
	return nil
}

func (pjq *ProblemJudgeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProblemJudge, error) {
	var (
		nodes       = []*ProblemJudge{}
		withFKs     = pjq.withFKs
		_spec       = pjq.querySpec()
		loadedTypes = [2]bool{
			pjq.withGroups != nil,
			pjq.withProblems != nil,
		}
	)
	if pjq.withGroups != nil || pjq.withProblems != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, problemjudge.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProblemJudge).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProblemJudge{config: pjq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pjq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pjq.withGroups; query != nil {
		if err := pjq.loadGroups(ctx, query, nodes, nil,
			func(n *ProblemJudge, e *Group) { n.Edges.Groups = e }); err != nil {
			return nil, err
		}
	}
	if query := pjq.withProblems; query != nil {
		if err := pjq.loadProblems(ctx, query, nodes, nil,
			func(n *ProblemJudge, e *Problem) { n.Edges.Problems = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pjq *ProblemJudgeQuery) loadGroups(ctx context.Context, query *GroupQuery, nodes []*ProblemJudge, init func(*ProblemJudge), assign func(*ProblemJudge, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProblemJudge)
	for i := range nodes {
		if nodes[i].group_problem_judges == nil {
			continue
		}
		fk := *nodes[i].group_problem_judges
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_problem_judges" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pjq *ProblemJudgeQuery) loadProblems(ctx context.Context, query *ProblemQuery, nodes []*ProblemJudge, init func(*ProblemJudge), assign func(*ProblemJudge, *Problem)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProblemJudge)
	for i := range nodes {
		if nodes[i].problem_problem_judges == nil {
			continue
		}
		fk := *nodes[i].problem_problem_judges
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(problem.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "problem_problem_judges" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pjq *ProblemJudgeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pjq.querySpec()
	_spec.Node.Columns = pjq.ctx.Fields
	if len(pjq.ctx.Fields) > 0 {
		_spec.Unique = pjq.ctx.Unique != nil && *pjq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pjq.driver, _spec)
}

func (pjq *ProblemJudgeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(problemjudge.Table, problemjudge.Columns, sqlgraph.NewFieldSpec(problemjudge.FieldID, field.TypeInt))
	_spec.From = pjq.sql
	if unique := pjq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pjq.path != nil {
		_spec.Unique = true
	}
	if fields := pjq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problemjudge.FieldID)
		for i := range fields {
			if fields[i] != problemjudge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pjq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pjq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pjq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pjq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pjq *ProblemJudgeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pjq.driver.Dialect())
	t1 := builder.Table(problemjudge.Table)
	columns := pjq.ctx.Fields
	if len(columns) == 0 {
		columns = problemjudge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pjq.sql != nil {
		selector = pjq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pjq.ctx.Unique != nil && *pjq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pjq.predicates {
		p(selector)
	}
	for _, p := range pjq.order {
		p(selector)
	}
	if offset := pjq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pjq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProblemJudgeGroupBy is the group-by builder for ProblemJudge entities.
type ProblemJudgeGroupBy struct {
	selector
	build *ProblemJudgeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pjgb *ProblemJudgeGroupBy) Aggregate(fns ...AggregateFunc) *ProblemJudgeGroupBy {
	pjgb.fns = append(pjgb.fns, fns...)
	return pjgb
}

// Scan applies the selector query and scans the result into the given value.
func (pjgb *ProblemJudgeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pjgb.build.ctx, "GroupBy")
	if err := pjgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProblemJudgeQuery, *ProblemJudgeGroupBy](ctx, pjgb.build, pjgb, pjgb.build.inters, v)
}

func (pjgb *ProblemJudgeGroupBy) sqlScan(ctx context.Context, root *ProblemJudgeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pjgb.fns))
	for _, fn := range pjgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pjgb.flds)+len(pjgb.fns))
		for _, f := range *pjgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pjgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pjgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProblemJudgeSelect is the builder for selecting fields of ProblemJudge entities.
type ProblemJudgeSelect struct {
	*ProblemJudgeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pjs *ProblemJudgeSelect) Aggregate(fns ...AggregateFunc) *ProblemJudgeSelect {
	pjs.fns = append(pjs.fns, fns...)
	return pjs
}

// Scan applies the selector query and scans the result into the given value.
func (pjs *ProblemJudgeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pjs.ctx, "Select")
	if err := pjs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProblemJudgeQuery, *ProblemJudgeSelect](ctx, pjs.ProblemJudgeQuery, pjs, pjs.inters, v)
}

func (pjs *ProblemJudgeSelect) sqlScan(ctx context.Context, root *ProblemJudgeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pjs.fns))
	for _, fn := range pjs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pjs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pjs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
