// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"sastoj/ent/submissioncase"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemCaseQuery is the builder for querying ProblemCase entities.
type ProblemCaseQuery struct {
	config
	ctx                 *QueryContext
	order               []problemcase.OrderOption
	inters              []Interceptor
	predicates          []predicate.ProblemCase
	withSubmissionCases *SubmissionCaseQuery
	withProblem         *ProblemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProblemCaseQuery builder.
func (pcq *ProblemCaseQuery) Where(ps ...predicate.ProblemCase) *ProblemCaseQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *ProblemCaseQuery) Limit(limit int) *ProblemCaseQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *ProblemCaseQuery) Offset(offset int) *ProblemCaseQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *ProblemCaseQuery) Unique(unique bool) *ProblemCaseQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *ProblemCaseQuery) Order(o ...problemcase.OrderOption) *ProblemCaseQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QuerySubmissionCases chains the current query on the "submission_cases" edge.
func (pcq *ProblemCaseQuery) QuerySubmissionCases() *SubmissionCaseQuery {
	query := (&SubmissionCaseClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problemcase.Table, problemcase.FieldID, selector),
			sqlgraph.To(submissioncase.Table, submissioncase.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, problemcase.SubmissionCasesTable, problemcase.SubmissionCasesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProblem chains the current query on the "problem" edge.
func (pcq *ProblemCaseQuery) QueryProblem() *ProblemQuery {
	query := (&ProblemClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problemcase.Table, problemcase.FieldID, selector),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problemcase.ProblemTable, problemcase.ProblemColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProblemCase entity from the query.
// Returns a *NotFoundError when no ProblemCase was found.
func (pcq *ProblemCaseQuery) First(ctx context.Context) (*ProblemCase, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{problemcase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *ProblemCaseQuery) FirstX(ctx context.Context) *ProblemCase {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProblemCase ID from the query.
// Returns a *NotFoundError when no ProblemCase ID was found.
func (pcq *ProblemCaseQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{problemcase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *ProblemCaseQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProblemCase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProblemCase entity is found.
// Returns a *NotFoundError when no ProblemCase entities are found.
func (pcq *ProblemCaseQuery) Only(ctx context.Context) (*ProblemCase, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{problemcase.Label}
	default:
		return nil, &NotSingularError{problemcase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *ProblemCaseQuery) OnlyX(ctx context.Context) *ProblemCase {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProblemCase ID in the query.
// Returns a *NotSingularError when more than one ProblemCase ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *ProblemCaseQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{problemcase.Label}
	default:
		err = &NotSingularError{problemcase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *ProblemCaseQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProblemCases.
func (pcq *ProblemCaseQuery) All(ctx context.Context) ([]*ProblemCase, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProblemCase, *ProblemCaseQuery]()
	return withInterceptors[[]*ProblemCase](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *ProblemCaseQuery) AllX(ctx context.Context) []*ProblemCase {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProblemCase IDs.
func (pcq *ProblemCaseQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(problemcase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *ProblemCaseQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *ProblemCaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*ProblemCaseQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *ProblemCaseQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *ProblemCaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *ProblemCaseQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProblemCaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *ProblemCaseQuery) Clone() *ProblemCaseQuery {
	if pcq == nil {
		return nil
	}
	return &ProblemCaseQuery{
		config:              pcq.config,
		ctx:                 pcq.ctx.Clone(),
		order:               append([]problemcase.OrderOption{}, pcq.order...),
		inters:              append([]Interceptor{}, pcq.inters...),
		predicates:          append([]predicate.ProblemCase{}, pcq.predicates...),
		withSubmissionCases: pcq.withSubmissionCases.Clone(),
		withProblem:         pcq.withProblem.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithSubmissionCases tells the query-builder to eager-load the nodes that are connected to
// the "submission_cases" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *ProblemCaseQuery) WithSubmissionCases(opts ...func(*SubmissionCaseQuery)) *ProblemCaseQuery {
	query := (&SubmissionCaseClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withSubmissionCases = query
	return pcq
}

// WithProblem tells the query-builder to eager-load the nodes that are connected to
// the "problem" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *ProblemCaseQuery) WithProblem(opts ...func(*ProblemQuery)) *ProblemCaseQuery {
	query := (&ProblemClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withProblem = query
	return pcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Point int16 `json:"point,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProblemCase.Query().
//		GroupBy(problemcase.FieldPoint).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *ProblemCaseQuery) GroupBy(field string, fields ...string) *ProblemCaseGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProblemCaseGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = problemcase.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Point int16 `json:"point,omitempty"`
//	}
//
//	client.ProblemCase.Query().
//		Select(problemcase.FieldPoint).
//		Scan(ctx, &v)
func (pcq *ProblemCaseQuery) Select(fields ...string) *ProblemCaseSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &ProblemCaseSelect{ProblemCaseQuery: pcq}
	sbuild.label = problemcase.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProblemCaseSelect configured with the given aggregations.
func (pcq *ProblemCaseQuery) Aggregate(fns ...AggregateFunc) *ProblemCaseSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *ProblemCaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !problemcase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *ProblemCaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProblemCase, error) {
	var (
		nodes       = []*ProblemCase{}
		_spec       = pcq.querySpec()
		loadedTypes = [2]bool{
			pcq.withSubmissionCases != nil,
			pcq.withProblem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProblemCase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProblemCase{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withSubmissionCases; query != nil {
		if err := pcq.loadSubmissionCases(ctx, query, nodes,
			func(n *ProblemCase) { n.Edges.SubmissionCases = []*SubmissionCase{} },
			func(n *ProblemCase, e *SubmissionCase) { n.Edges.SubmissionCases = append(n.Edges.SubmissionCases, e) }); err != nil {
			return nil, err
		}
	}
	if query := pcq.withProblem; query != nil {
		if err := pcq.loadProblem(ctx, query, nodes, nil,
			func(n *ProblemCase, e *Problem) { n.Edges.Problem = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *ProblemCaseQuery) loadSubmissionCases(ctx context.Context, query *SubmissionCaseQuery, nodes []*ProblemCase, init func(*ProblemCase), assign func(*ProblemCase, *SubmissionCase)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*ProblemCase)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(submissioncase.FieldProblemCaseID)
	}
	query.Where(predicate.SubmissionCase(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(problemcase.SubmissionCasesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProblemCaseID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "problem_case_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pcq *ProblemCaseQuery) loadProblem(ctx context.Context, query *ProblemQuery, nodes []*ProblemCase, init func(*ProblemCase), assign func(*ProblemCase, *Problem)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ProblemCase)
	for i := range nodes {
		fk := nodes[i].ProblemID
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
			return fmt.Errorf(`unexpected foreign-key "problem_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pcq *ProblemCaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *ProblemCaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(problemcase.Table, problemcase.Columns, sqlgraph.NewFieldSpec(problemcase.FieldID, field.TypeInt64))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problemcase.FieldID)
		for i := range fields {
			if fields[i] != problemcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pcq.withProblem != nil {
			_spec.Node.AddColumnOnce(problemcase.FieldProblemID)
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *ProblemCaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(problemcase.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = problemcase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProblemCaseGroupBy is the group-by builder for ProblemCase entities.
type ProblemCaseGroupBy struct {
	selector
	build *ProblemCaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *ProblemCaseGroupBy) Aggregate(fns ...AggregateFunc) *ProblemCaseGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *ProblemCaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProblemCaseQuery, *ProblemCaseGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *ProblemCaseGroupBy) sqlScan(ctx context.Context, root *ProblemCaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProblemCaseSelect is the builder for selecting fields of ProblemCase entities.
type ProblemCaseSelect struct {
	*ProblemCaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *ProblemCaseSelect) Aggregate(fns ...AggregateFunc) *ProblemCaseSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *ProblemCaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProblemCaseQuery, *ProblemCaseSelect](ctx, pcs.ProblemCaseQuery, pcs, pcs.inters, v)
}

func (pcs *ProblemCaseSelect) sqlScan(ctx context.Context, root *ProblemCaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
