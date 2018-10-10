package gockle

import (
	"context"
	"github.com/gocql/gocql"
)

// Query represents a CQL query.
type Query interface {
	// Consistency sets the consistency level for this query. If no consistency
	// level have been set, the default consistency level of the cluster
	// is used.
	Consistency(c gocql.Consistency) Query
	// PageSize will tell the iterator to fetch the result in pages of size n.
	// This is useful for iterating over large result sets, but setting the
	// page size too low might decrease the performance. This feature is only
	// available in Cassandra 2 and onwards.
	PageSize(n int) Query
	// WithContext will set the context to use during a query, it will be used to
	// timeout when waiting for responses from Cassandra.
	WithContext(ctx context.Context) Query
	// PageState sets the paging state for the query to resume paging from a specific
	// point in time. Setting this will disable to query paging for this query, and
	// must be used for all subsequent pages.
	PageState(state []byte) Query
	// Exec executes the query without returning any rows.
	Exec() error
	// Iter executes the query and returns an iterator capable of iterating
	// over all results.
	Iter() Iterator
	// MapScan executes the query, copies the columns of the first selected
	// row into the map pointed at by m and discards the rest. If no rows
	// were selected, ErrNotFound is returned.
	MapScan(m map[string]interface{}) error
	// Scan executes the query, copies the columns of the first selected
	// row into the values pointed at by dest and discards the rest. If no rows
	// were selected, ErrNotFound is returned.
	Scan(dest ...interface{}) error
	// Release releases a query back into a pool of queries. Released Queries
	// cannot be reused.
	//
	// Example:
	//              qry := session.Query("SELECT * FROM my_table")
	//              qry.Exec()
	//              qry.Release()
	Release()
}

var (
	_ Query = QueryMock{}
	_ Query = query{}
)

type query struct {
	q *gocql.Query
}

func (q query) Consistency(c gocql.Consistency) Query {
	return &query{q: q.q.Consistency(c)}
}

func (q query) PageSize(n int) Query {
	return &query{q: q.q.PageSize(n)}
}

func (q query) WithContext(ctx context.Context) Query {
	return &query{q: q.q.WithContext(ctx)}
}

func (q query) PageState(state []byte) Query {
	return &query{q: q.q.PageState(state)}
}

func (q query) Exec() error {
	return q.q.Exec()
}

func (q query) Iter() Iterator {
	return &iterator{i: q.q.Iter()}
}

func (q query) MapScan(m map[string]interface{}) error {
	return q.q.MapScan(m)
}

func (q query) Scan(dest ...interface{}) error {
	return q.q.Scan(dest...)
}

func (q query) Release() {
	q.q.Release()
}
