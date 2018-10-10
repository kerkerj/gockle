package gockle

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

// QueryMock is a mock Query. See github.com/maraino/go-mock.
type QueryMock struct {
	mock.Mock
}

// Consistency implements Query.
func (m QueryMock) Consistency(c gocql.Consistency) Query {
	return m.Called(c).Get(0).(Query)
}

// PageSize implements Query.
func (m QueryMock) PageSize(n int) Query {
	return m.Called(n).Get(0).(Query)
}

// WithContext implements Query.
func (m QueryMock) WithContext(ctx context.Context) Query {
	return m.Called(ctx).Get(0).(Query)
}

// PageState implements Query.
func (m QueryMock) PageState(state []byte) Query {
	return m.Called(state).Get(0).(Query)
}

// Exec implements Query.
func (m QueryMock) Exec() error {
	return m.Called().Error(0)
}

// Iter implements Query.
func (m QueryMock) Iter() Iterator {
	return m.Called().Get(0).(Iterator)
}

// MapScan implements Query.
func (m QueryMock) MapScan(mm map[string]interface{}) error {
	return m.Called(mm).Error(0)
}

// Scan implements Query.
func (m QueryMock) Scan(dest ...interface{}) error {
	return m.Called(dest).Error(0)
}

// Release implements Query.
func (m QueryMock) Release() {
	m.Called()
}
