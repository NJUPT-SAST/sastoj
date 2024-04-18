// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"sastoj/ent/loginsession"
	"sastoj/ent/predicate"
	"sastoj/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginSessionQuery is the builder for querying LoginSession entities.
type LoginSessionQuery struct {
	config
	ctx        *QueryContext
	order      []loginsession.OrderOption
	inters     []Interceptor
	predicates []predicate.LoginSession
	withUsers  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoginSessionQuery builder.
func (lsq *LoginSessionQuery) Where(ps ...predicate.LoginSession) *LoginSessionQuery {
	lsq.predicates = append(lsq.predicates, ps...)
	return lsq
}

// Limit the number of records to be returned by this query.
func (lsq *LoginSessionQuery) Limit(limit int) *LoginSessionQuery {
	lsq.ctx.Limit = &limit
	return lsq
}

// Offset to start from.
func (lsq *LoginSessionQuery) Offset(offset int) *LoginSessionQuery {
	lsq.ctx.Offset = &offset
	return lsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lsq *LoginSessionQuery) Unique(unique bool) *LoginSessionQuery {
	lsq.ctx.Unique = &unique
	return lsq
}

// Order specifies how the records should be ordered.
func (lsq *LoginSessionQuery) Order(o ...loginsession.OrderOption) *LoginSessionQuery {
	lsq.order = append(lsq.order, o...)
	return lsq
}

// QueryUsers chains the current query on the "users" edge.
func (lsq *LoginSessionQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: lsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loginsession.Table, loginsession.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, loginsession.UsersTable, loginsession.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(lsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LoginSession entity from the query.
// Returns a *NotFoundError when no LoginSession was found.
func (lsq *LoginSessionQuery) First(ctx context.Context) (*LoginSession, error) {
	nodes, err := lsq.Limit(1).All(setContextOp(ctx, lsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loginsession.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lsq *LoginSessionQuery) FirstX(ctx context.Context) *LoginSession {
	node, err := lsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoginSession ID from the query.
// Returns a *NotFoundError when no LoginSession ID was found.
func (lsq *LoginSessionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lsq.Limit(1).IDs(setContextOp(ctx, lsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loginsession.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lsq *LoginSessionQuery) FirstIDX(ctx context.Context) int {
	id, err := lsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoginSession entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoginSession entity is found.
// Returns a *NotFoundError when no LoginSession entities are found.
func (lsq *LoginSessionQuery) Only(ctx context.Context) (*LoginSession, error) {
	nodes, err := lsq.Limit(2).All(setContextOp(ctx, lsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loginsession.Label}
	default:
		return nil, &NotSingularError{loginsession.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lsq *LoginSessionQuery) OnlyX(ctx context.Context) *LoginSession {
	node, err := lsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoginSession ID in the query.
// Returns a *NotSingularError when more than one LoginSession ID is found.
// Returns a *NotFoundError when no entities are found.
func (lsq *LoginSessionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lsq.Limit(2).IDs(setContextOp(ctx, lsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loginsession.Label}
	default:
		err = &NotSingularError{loginsession.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lsq *LoginSessionQuery) OnlyIDX(ctx context.Context) int {
	id, err := lsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoginSessions.
func (lsq *LoginSessionQuery) All(ctx context.Context) ([]*LoginSession, error) {
	ctx = setContextOp(ctx, lsq.ctx, "All")
	if err := lsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoginSession, *LoginSessionQuery]()
	return withInterceptors[[]*LoginSession](ctx, lsq, qr, lsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lsq *LoginSessionQuery) AllX(ctx context.Context) []*LoginSession {
	nodes, err := lsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoginSession IDs.
func (lsq *LoginSessionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lsq.ctx.Unique == nil && lsq.path != nil {
		lsq.Unique(true)
	}
	ctx = setContextOp(ctx, lsq.ctx, "IDs")
	if err = lsq.Select(loginsession.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lsq *LoginSessionQuery) IDsX(ctx context.Context) []int {
	ids, err := lsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lsq *LoginSessionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lsq.ctx, "Count")
	if err := lsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lsq, querierCount[*LoginSessionQuery](), lsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lsq *LoginSessionQuery) CountX(ctx context.Context) int {
	count, err := lsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lsq *LoginSessionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lsq.ctx, "Exist")
	switch _, err := lsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lsq *LoginSessionQuery) ExistX(ctx context.Context) bool {
	exist, err := lsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoginSessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lsq *LoginSessionQuery) Clone() *LoginSessionQuery {
	if lsq == nil {
		return nil
	}
	return &LoginSessionQuery{
		config:     lsq.config,
		ctx:        lsq.ctx.Clone(),
		order:      append([]loginsession.OrderOption{}, lsq.order...),
		inters:     append([]Interceptor{}, lsq.inters...),
		predicates: append([]predicate.LoginSession{}, lsq.predicates...),
		withUsers:  lsq.withUsers.Clone(),
		// clone intermediate query.
		sql:  lsq.sql.Clone(),
		path: lsq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (lsq *LoginSessionQuery) WithUsers(opts ...func(*UserQuery)) *LoginSessionQuery {
	query := (&UserClient{config: lsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lsq.withUsers = query
	return lsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LoginSession.Query().
//		GroupBy(loginsession.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lsq *LoginSessionQuery) GroupBy(field string, fields ...string) *LoginSessionGroupBy {
	lsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoginSessionGroupBy{build: lsq}
	grbuild.flds = &lsq.ctx.Fields
	grbuild.label = loginsession.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.LoginSession.Query().
//		Select(loginsession.FieldCreateTime).
//		Scan(ctx, &v)
func (lsq *LoginSessionQuery) Select(fields ...string) *LoginSessionSelect {
	lsq.ctx.Fields = append(lsq.ctx.Fields, fields...)
	sbuild := &LoginSessionSelect{LoginSessionQuery: lsq}
	sbuild.label = loginsession.Label
	sbuild.flds, sbuild.scan = &lsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoginSessionSelect configured with the given aggregations.
func (lsq *LoginSessionQuery) Aggregate(fns ...AggregateFunc) *LoginSessionSelect {
	return lsq.Select().Aggregate(fns...)
}

func (lsq *LoginSessionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lsq); err != nil {
				return err
			}
		}
	}
	for _, f := range lsq.ctx.Fields {
		if !loginsession.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lsq.path != nil {
		prev, err := lsq.path(ctx)
		if err != nil {
			return err
		}
		lsq.sql = prev
	}
	return nil
}

func (lsq *LoginSessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoginSession, error) {
	var (
		nodes       = []*LoginSession{}
		withFKs     = lsq.withFKs
		_spec       = lsq.querySpec()
		loadedTypes = [1]bool{
			lsq.withUsers != nil,
		}
	)
	if lsq.withUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, loginsession.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoginSession).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoginSession{config: lsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lsq.withUsers; query != nil {
		if err := lsq.loadUsers(ctx, query, nodes, nil,
			func(n *LoginSession, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lsq *LoginSessionQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*LoginSession, init func(*LoginSession), assign func(*LoginSession, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*LoginSession)
	for i := range nodes {
		if nodes[i].user_login_sessions == nil {
			continue
		}
		fk := *nodes[i].user_login_sessions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_login_sessions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lsq *LoginSessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lsq.querySpec()
	_spec.Node.Columns = lsq.ctx.Fields
	if len(lsq.ctx.Fields) > 0 {
		_spec.Unique = lsq.ctx.Unique != nil && *lsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lsq.driver, _spec)
}

func (lsq *LoginSessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loginsession.Table, loginsession.Columns, sqlgraph.NewFieldSpec(loginsession.FieldID, field.TypeInt))
	_spec.From = lsq.sql
	if unique := lsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lsq.path != nil {
		_spec.Unique = true
	}
	if fields := lsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginsession.FieldID)
		for i := range fields {
			if fields[i] != loginsession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lsq *LoginSessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lsq.driver.Dialect())
	t1 := builder.Table(loginsession.Table)
	columns := lsq.ctx.Fields
	if len(columns) == 0 {
		columns = loginsession.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lsq.sql != nil {
		selector = lsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lsq.ctx.Unique != nil && *lsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lsq.predicates {
		p(selector)
	}
	for _, p := range lsq.order {
		p(selector)
	}
	if offset := lsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoginSessionGroupBy is the group-by builder for LoginSession entities.
type LoginSessionGroupBy struct {
	selector
	build *LoginSessionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lsgb *LoginSessionGroupBy) Aggregate(fns ...AggregateFunc) *LoginSessionGroupBy {
	lsgb.fns = append(lsgb.fns, fns...)
	return lsgb
}

// Scan applies the selector query and scans the result into the given value.
func (lsgb *LoginSessionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lsgb.build.ctx, "GroupBy")
	if err := lsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoginSessionQuery, *LoginSessionGroupBy](ctx, lsgb.build, lsgb, lsgb.build.inters, v)
}

func (lsgb *LoginSessionGroupBy) sqlScan(ctx context.Context, root *LoginSessionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lsgb.fns))
	for _, fn := range lsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lsgb.flds)+len(lsgb.fns))
		for _, f := range *lsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoginSessionSelect is the builder for selecting fields of LoginSession entities.
type LoginSessionSelect struct {
	*LoginSessionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lss *LoginSessionSelect) Aggregate(fns ...AggregateFunc) *LoginSessionSelect {
	lss.fns = append(lss.fns, fns...)
	return lss
}

// Scan applies the selector query and scans the result into the given value.
func (lss *LoginSessionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lss.ctx, "Select")
	if err := lss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoginSessionQuery, *LoginSessionSelect](ctx, lss.LoginSessionQuery, lss, lss.inters, v)
}

func (lss *LoginSessionSelect) sqlScan(ctx context.Context, root *LoginSessionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lss.fns))
	for _, fn := range lss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
