package gockle

import (
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// SessionMock is a mock Session. See github.com/maraino/go-mock.
type SessionMock struct {
	mock.Mock
}

// Batch implements Session.
func (m SessionMock) Batch(kind BatchKind) Batch {
	return m.Called(kind).Get(0).(Batch)
}

// Close implements Session.
func (m SessionMock) Close() {
	m.Called()
}

// Columns implements Session.
func (m SessionMock) Columns(keyspace, table string) (map[string]gocql.TypeInfo, error) {
	var r = m.Called(keyspace, table)

	return r.Get(0).(map[string]gocql.TypeInfo), r.Error(1)
}

// Exec implements Session.
func (m SessionMock) Exec(statement string, arguments ...interface{}) error {
	return m.Called(statement, arguments).Error(0)
}

// Scan implements Session.
func (m SessionMock) Scan(statement string, results []interface{}, arguments ...interface{}) error {
	return m.Called(statement, results, arguments).Error(0)
}

// ScanIterator implements Session.
func (m SessionMock) ScanIterator(statement string, arguments ...interface{}) Iterator {
	return m.Called(statement, arguments).Get(0).(Iterator)
}

// ScanMap implements Session.
func (m SessionMock) ScanMap(statement string, results map[string]interface{}, arguments ...interface{}) error {
	return m.Called(statement, results, arguments).Error(0)
}

// ScanMapSlice implements Session.
func (m SessionMock) ScanMapSlice(statement string, arguments ...interface{}) ([]map[string]interface{}, error) {
	var r = m.Called(statement, arguments)

	return r.Get(0).([]map[string]interface{}), r.Error(1)
}

// ScanMapTx implements Session.
func (m SessionMock) ScanMapTx(statement string, results map[string]interface{}, arguments ...interface{}) (bool, error) {
	var r = m.Called(statement, results, arguments)

	return r.Bool(0), r.Error(1)
}

// Tables implements Session.
func (m SessionMock) Tables(keyspace string) ([]string, error) {
	var r = m.Called(keyspace)

	return r.Get(0).([]string), r.Error(1)
}

// Query implements Session.
func (m SessionMock) Query(statement string, arguments ...interface{}) Query {
	return m.Called(statement, arguments).Get(0).(Query)
}
