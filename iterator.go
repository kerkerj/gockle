package gockle

import (
	"github.com/gocql/gocql"
)

// Iterator iterates CQL query result rows.
type Iterator interface {
	// Close closes the Iterator.
	Close() error

	// Scan puts the current result row in results and returns whether there are
	// more result rows.
	Scan(results ...interface{}) bool

	// ScanMap puts the current result row in results and returns whether there are
	// more result rows.
	ScanMap(results map[string]interface{}) bool

	// WillSwitchPage detects if iterator reached end of current page and the
	// next page is available.
	WillSwitchPage() bool
	// PageState return the current paging state for a query which can be used
	// for subsequent quries to resume paging this point.
	PageState() []byte

	// SliceMap is a helper function to make the API easier to use
	// returns the data from the query in the form of []map[string]interface{}
	SliceMap() ([]map[string]interface{}, error)
}

var (
	_ Iterator = &IteratorMock{}
	_ Iterator = iterator{}
)

type iterator struct {
	i *gocql.Iter
}

func (i iterator) Close() error {
	return i.i.Close()
}

func (i iterator) Scan(results ...interface{}) bool {
	return i.i.Scan(results...)
}

func (i iterator) ScanMap(results map[string]interface{}) bool {
	return i.i.MapScan(results)
}

func (i iterator) WillSwitchPage() bool {
	return i.i.WillSwitchPage()
}

func (i iterator) PageState() []byte {
	return i.i.PageState()
}

func (i iterator) SliceMap() ([]map[string]interface{}, error) {
	return i.i.SliceMap()
}
