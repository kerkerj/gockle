package gockle

import (
	"fmt"

	"github.com/gocql/gocql"
)

func metadata(s *gocql.Session, keyspace string) (*gocql.KeyspaceMetadata, error) {
	var m, err = s.KeyspaceMetadata(keyspace)

	if err != nil {
		return nil, err
	}

	if !m.DurableWrites && m.Name == keyspace && m.StrategyClass == "" && len(m.StrategyOptions) == 0 && len(m.Tables) == 0 {
		return nil, fmt.Errorf("gockle: keyspace %v invalid", keyspace)
	}

	return m, nil
}

// Session is a Cassandra connection. The Query methods run CQL queries. The
// Columns and Tables methods provide simple metadata.
type Session interface {
	// Batch returns a new Batch for the Session.
	Batch(kind BatchKind) Batch

	// Close closes the Session.
	Close()

	// Columns returns a map from column names to types for keyspace and table.
	// Schema changes during a session are not reflected; you must open a new
	// Session to observe them.
	Columns(keyspace, table string) (map[string]gocql.TypeInfo, error)

	// Exec executes the query for statement and arguments.
	Exec(statement string, arguments ...interface{}) error

	// Scan executes the query for statement and arguments and puts the first
	// result row in results.
	Scan(statement string, results []interface{}, arguments ...interface{}) error

	// ScanIterator executes the query for statement and arguments and returns an
	// Iterator for the results.
	ScanIterator(statement string, arguments ...interface{}) Iterator

	// ScanMap executes the query for statement and arguments and puts the first
	// result row in results.
	ScanMap(statement string, results map[string]interface{}, arguments ...interface{}) error

	// ScanMapSlice executes the query for statement and arguments and returns all
	// the result rows.
	ScanMapSlice(statement string, arguments ...interface{}) ([]map[string]interface{}, error)

	// ScanMapTx executes the query for statement and arguments as a lightweight
	// transaction. If the query is not applied, it puts the current values for the
	// conditional columns in results. It returns whether the query is applied.
	ScanMapTx(statement string, results map[string]interface{}, arguments ...interface{}) (bool, error)

	// Tables returns the table names for keyspace. Schema changes during a session
	// are not reflected; you must open a new Session to observe them.
	Tables(keyspace string) ([]string, error)

	// Query generates a new query object for interacting with the database.
	// Further details of the query may be tweaked using the resulting query
	// value before the query is executed. Query is automatically prepared if
	// it has not previously been executed.
	Query(statement string, arguments ...interface{}) Query
}

var (
	_ Session = &SessionMock{}
	_ Session = session{}
)

// NewSession returns a new Session for s.
func NewSession(s *gocql.Session) Session {
	return session{s: s}
}

// NewSimpleSession returns a new Session for hosts. It uses native protocol
// version 4.
func NewSimpleSession(hosts ...string) (Session, error) {
	var c = gocql.NewCluster(hosts...)

	c.ProtoVersion = 4

	var s, err = c.CreateSession()

	if err != nil {
		return nil, err
	}

	return session{s: s}, nil
}

type session struct {
	s *gocql.Session
}

func (s session) Batch(kind BatchKind) Batch {
	return batch{b: s.s.NewBatch(gocql.BatchType(kind)), s: s.s}
}

func (s session) Close() {
	s.s.Close()
}

func (s session) Columns(keyspace, table string) (map[string]gocql.TypeInfo, error) {
	var m, err = metadata(s.s, keyspace)

	if err != nil {
		return nil, err
	}

	var t, ok = m.Tables[table]

	if !ok {
		return nil, fmt.Errorf("gockle: table %v.%v invalid", keyspace, table)
	}

	var types = map[string]gocql.TypeInfo{}

	for n, c := range t.Columns {
		types[n] = c.Type
	}

	return types, nil
}

func (s session) Exec(statement string, arguments ...interface{}) error {
	return s.s.Query(statement, arguments...).Exec()
}

func (s session) Scan(statement string, results []interface{}, arguments ...interface{}) error {
	return s.s.Query(statement, arguments...).Scan(results...)
}

func (s session) ScanIterator(statement string, arguments ...interface{}) Iterator {
	return iterator{i: s.s.Query(statement, arguments...).Iter()}
}

func (s session) ScanMap(statement string, results map[string]interface{}, arguments ...interface{}) error {
	return s.s.Query(statement, arguments...).MapScan(results)
}

func (s session) ScanMapSlice(statement string, arguments ...interface{}) ([]map[string]interface{}, error) {
	return s.s.Query(statement, arguments...).Iter().SliceMap()
}

func (s session) ScanMapTx(statement string, results map[string]interface{}, arguments ...interface{}) (bool, error) {
	return s.s.Query(statement, arguments...).MapScanCAS(results)
}

func (s session) Tables(keyspace string) ([]string, error) {
	var m, err = metadata(s.s, keyspace)

	if err != nil {
		return nil, err
	}

	var ts []string

	for t := range m.Tables {
		ts = append(ts, t)
	}

	return ts, nil
}

func (s session) Query(statement string, arguments ...interface{}) Query {
	return query{q: s.s.Query(statement, arguments...)}
}
