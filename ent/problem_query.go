// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/predicate"
	"sastoj/ent/problem"
	"sastoj/ent/problemtype"
	"sastoj/ent/submission"
	"sastoj/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemQuery is the builder for querying Problem entities.
type ProblemQuery struct {
	config
	ctx              *QueryContext
	order            []problem.OrderOption
	inters           []Interceptor
	predicates       []predicate.Problem
	withSubmission   *SubmissionQuery
	withContest      *ContestQuery
	withOwner        *UserQuery
	withProblemType  *ProblemTypeQuery
	withAdjudicators *GroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProblemQuery builder.
func (pq *ProblemQuery) Where(ps ...predicate.Problem) *ProblemQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProblemQuery) Limit(limit int) *ProblemQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProblemQuery) Offset(offset int) *ProblemQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProblemQuery) Unique(unique bool) *ProblemQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProblemQuery) Order(o ...problem.OrderOption) *ProblemQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QuerySubmission chains the current query on the "submission" edge.
func (pq *ProblemQuery) QuerySubmission() *SubmissionQuery {
	query := (&SubmissionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, selector),
			sqlgraph.To(submission.Table, submission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, problem.SubmissionTable, problem.SubmissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContest chains the current query on the "contest" edge.
func (pq *ProblemQuery) QueryContest() *ContestQuery {
	query := (&ContestClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, selector),
			sqlgraph.To(contest.Table, contest.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.ContestTable, problem.ContestColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (pq *ProblemQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.OwnerTable, problem.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProblemType chains the current query on the "problem_type" edge.
func (pq *ProblemQuery) QueryProblemType() *ProblemTypeQuery {
	query := (&ProblemTypeClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, selector),
			sqlgraph.To(problemtype.Table, problemtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.ProblemTypeTable, problem.ProblemTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAdjudicators chains the current query on the "adjudicators" edge.
func (pq *ProblemQuery) QueryAdjudicators() *GroupQuery {
	query := (&GroupClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, problem.AdjudicatorsTable, problem.AdjudicatorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Problem entity from the query.
// Returns a *NotFoundError when no Problem was found.
func (pq *ProblemQuery) First(ctx context.Context) (*Problem, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{problem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProblemQuery) FirstX(ctx context.Context) *Problem {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Problem ID from the query.
// Returns a *NotFoundError when no Problem ID was found.
func (pq *ProblemQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{problem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProblemQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Problem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Problem entity is found.
// Returns a *NotFoundError when no Problem entities are found.
func (pq *ProblemQuery) Only(ctx context.Context) (*Problem, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{problem.Label}
	default:
		return nil, &NotSingularError{problem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProblemQuery) OnlyX(ctx context.Context) *Problem {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Problem ID in the query.
// Returns a *NotSingularError when more than one Problem ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProblemQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{problem.Label}
	default:
		err = &NotSingularError{problem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProblemQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Problems.
func (pq *ProblemQuery) All(ctx context.Context) ([]*Problem, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Problem, *ProblemQuery]()
	return withInterceptors[[]*Problem](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProblemQuery) AllX(ctx context.Context) []*Problem {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Problem IDs.
func (pq *ProblemQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(problem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProblemQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProblemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProblemQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProblemQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProblemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProblemQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProblemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProblemQuery) Clone() *ProblemQuery {
	if pq == nil {
		return nil
	}
	return &ProblemQuery{
		config:           pq.config,
		ctx:              pq.ctx.Clone(),
		order:            append([]problem.OrderOption{}, pq.order...),
		inters:           append([]Interceptor{}, pq.inters...),
		predicates:       append([]predicate.Problem{}, pq.predicates...),
		withSubmission:   pq.withSubmission.Clone(),
		withContest:      pq.withContest.Clone(),
		withOwner:        pq.withOwner.Clone(),
		withProblemType:  pq.withProblemType.Clone(),
		withAdjudicators: pq.withAdjudicators.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithSubmission tells the query-builder to eager-load the nodes that are connected to
// the "submission" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProblemQuery) WithSubmission(opts ...func(*SubmissionQuery)) *ProblemQuery {
	query := (&SubmissionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withSubmission = query
	return pq
}

// WithContest tells the query-builder to eager-load the nodes that are connected to
// the "contest" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProblemQuery) WithContest(opts ...func(*ContestQuery)) *ProblemQuery {
	query := (&ContestClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withContest = query
	return pq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProblemQuery) WithOwner(opts ...func(*UserQuery)) *ProblemQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// WithProblemType tells the query-builder to eager-load the nodes that are connected to
// the "problem_type" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProblemQuery) WithProblemType(opts ...func(*ProblemTypeQuery)) *ProblemQuery {
	query := (&ProblemTypeClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withProblemType = query
	return pq
}

// WithAdjudicators tells the query-builder to eager-load the nodes that are connected to
// the "adjudicators" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProblemQuery) WithAdjudicators(opts ...func(*GroupQuery)) *ProblemQuery {
	query := (&GroupClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAdjudicators = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProblemTypeID int64 `json:"problem_type_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Problem.Query().
//		GroupBy(problem.FieldProblemTypeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProblemQuery) GroupBy(field string, fields ...string) *ProblemGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProblemGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = problem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProblemTypeID int64 `json:"problem_type_id,omitempty"`
//	}
//
//	client.Problem.Query().
//		Select(problem.FieldProblemTypeID).
//		Scan(ctx, &v)
func (pq *ProblemQuery) Select(fields ...string) *ProblemSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProblemSelect{ProblemQuery: pq}
	sbuild.label = problem.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProblemSelect configured with the given aggregations.
func (pq *ProblemQuery) Aggregate(fns ...AggregateFunc) *ProblemSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProblemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !problem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProblemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Problem, error) {
	var (
		nodes       = []*Problem{}
		_spec       = pq.querySpec()
		loadedTypes = [5]bool{
			pq.withSubmission != nil,
			pq.withContest != nil,
			pq.withOwner != nil,
			pq.withProblemType != nil,
			pq.withAdjudicators != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Problem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Problem{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withSubmission; query != nil {
		if err := pq.loadSubmission(ctx, query, nodes,
			func(n *Problem) { n.Edges.Submission = []*Submission{} },
			func(n *Problem, e *Submission) { n.Edges.Submission = append(n.Edges.Submission, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withContest; query != nil {
		if err := pq.loadContest(ctx, query, nodes, nil,
			func(n *Problem, e *Contest) { n.Edges.Contest = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withOwner; query != nil {
		if err := pq.loadOwner(ctx, query, nodes, nil,
			func(n *Problem, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withProblemType; query != nil {
		if err := pq.loadProblemType(ctx, query, nodes, nil,
			func(n *Problem, e *ProblemType) { n.Edges.ProblemType = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withAdjudicators; query != nil {
		if err := pq.loadAdjudicators(ctx, query, nodes,
			func(n *Problem) { n.Edges.Adjudicators = []*Group{} },
			func(n *Problem, e *Group) { n.Edges.Adjudicators = append(n.Edges.Adjudicators, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProblemQuery) loadSubmission(ctx context.Context, query *SubmissionQuery, nodes []*Problem, init func(*Problem), assign func(*Problem, *Submission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Problem)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(submission.FieldProblemID)
	}
	query.Where(predicate.Submission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(problem.SubmissionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProblemID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "problem_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProblemQuery) loadContest(ctx context.Context, query *ContestQuery, nodes []*Problem, init func(*Problem), assign func(*Problem, *Contest)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Problem)
	for i := range nodes {
		fk := nodes[i].ContestID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(contest.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "contest_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProblemQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Problem, init func(*Problem), assign func(*Problem, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Problem)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProblemQuery) loadProblemType(ctx context.Context, query *ProblemTypeQuery, nodes []*Problem, init func(*Problem), assign func(*Problem, *ProblemType)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Problem)
	for i := range nodes {
		fk := nodes[i].ProblemTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(problemtype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "problem_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProblemQuery) loadAdjudicators(ctx context.Context, query *GroupQuery, nodes []*Problem, init func(*Problem), assign func(*Problem, *Group)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Problem)
	nids := make(map[int64]map[*Problem]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(problem.AdjudicatorsTable)
		s.Join(joinT).On(s.C(group.FieldID), joinT.C(problem.AdjudicatorsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(problem.AdjudicatorsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(problem.AdjudicatorsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*Problem]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Group](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "adjudicators" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pq *ProblemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProblemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(problem.Table, problem.Columns, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt64))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problem.FieldID)
		for i := range fields {
			if fields[i] != problem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withContest != nil {
			_spec.Node.AddColumnOnce(problem.FieldContestID)
		}
		if pq.withOwner != nil {
			_spec.Node.AddColumnOnce(problem.FieldUserID)
		}
		if pq.withProblemType != nil {
			_spec.Node.AddColumnOnce(problem.FieldProblemTypeID)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProblemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(problem.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = problem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProblemGroupBy is the group-by builder for Problem entities.
type ProblemGroupBy struct {
	selector
	build *ProblemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProblemGroupBy) Aggregate(fns ...AggregateFunc) *ProblemGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProblemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProblemQuery, *ProblemGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProblemGroupBy) sqlScan(ctx context.Context, root *ProblemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProblemSelect is the builder for selecting fields of Problem entities.
type ProblemSelect struct {
	*ProblemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProblemSelect) Aggregate(fns ...AggregateFunc) *ProblemSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProblemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProblemQuery, *ProblemSelect](ctx, ps.ProblemQuery, ps, ps.inters, v)
}

func (ps *ProblemSelect) sqlScan(ctx context.Context, root *ProblemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
