package gockle

import (
	"github.com/maraino/go-mock"
)

// BatchMock is a mock Batch. See github.com/maraino/go-mock.
type BatchMock struct {
	mock.Mock
}

// Add implements Batch.
func (m BatchMock) Add(statement string, arguments ...interface{}) {
	m.Called(statement, arguments)
}

// Exec implements Batch.
func (m BatchMock) Exec() error {
	return m.Called().Error(0)
}

// ExecTx implements Batch.
func (m BatchMock) ExecTx() ([]map[string]interface{}, error) {
	var r = m.Called()

	return r.Get(0).([]map[string]interface{}), r.Error(1)
}
