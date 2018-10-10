package gockle

import (
	"github.com/maraino/go-mock"
)

// IteratorMock is a mock Iterator. See github.com/maraino/go-mock.
type IteratorMock struct {
	mock.Mock
}

// Close implements Iterator.
func (m IteratorMock) Close() error {
	return m.Called().Error(0)
}

// Scan implements Iterator.
func (m IteratorMock) Scan(results ...interface{}) bool {
	return m.Called(results).Bool(0)
}

// ScanMap implements Iterator.
func (m IteratorMock) ScanMap(results map[string]interface{}) bool {
	return m.Called(results).Bool(0)
}

// WillSwitchPage implements Iterator.
func (m IteratorMock) WillSwitchPage() bool {
	return m.Called().Bool(0)
}

// PageState implements Iterator.
func (m IteratorMock) PageState() []byte {
	return m.Called().Bytes(0)
}
