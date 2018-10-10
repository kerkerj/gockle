package gockle

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

var mySession = &SessionMock{}

func ExampleIterator_ScanMap() {
	var iteratorMock = &IteratorMock{}

	iteratorMock.On("ScanMap", mock.Anything).Return(func(m map[string]interface{}) bool {
		m["id"] = 1
		m["name"] = "alex"
		return false
	})

	iteratorMock.On("Close").Return(nil)

	var sessionMock = &SessionMock{}

	const query = "select * from users"

	sessionMock.On("ScanIterator", query, mock.Anything).Return(iteratorMock)
	sessionMock.On("Close").Return(nil)

	var session Session = sessionMock
	var iterator = session.ScanIterator(query)
	var row = map[string]interface{}{}

	for more := true; more; {
		more = iterator.ScanMap(row)

		fmt.Printf("id = %v, name = %v\n", row["id"], row["name"])
	}

	if err := iterator.Close(); err != nil {
		fmt.Println(err)
	}

	session.Close()

	// Output: id = 1, name = alex
}

func ExampleSession_Batch() {
	var batchMock = &BatchMock{}

	batchMock.On("Add", "insert into users (id, name) values (1, 'alex')", mock.Anything).Return()
	batchMock.On("Exec").Return(fmt.Errorf("invalid"))

	var sessionMock = &SessionMock{}

	sessionMock.On("Batch", BatchLogged).Return(batchMock)
	sessionMock.On("Close").Return()

	var session Session = sessionMock
	var batch = session.Batch(BatchLogged)

	batch.Add("insert into users (id, name) values (1, 'alex')")

	if err := batch.Exec(); err != nil {
		fmt.Println(err)
	}

	session.Close()

	// Output: invalid
}

func ExampleSession_ScanMapSlice() {
	var sessionMock = &SessionMock{}

	const query = "select * from users"

	sessionMock.On("ScanMapSlice", query, mock.Anything).Return([]map[string]interface{}{{"id": 1, "name": "alex"}}, nil)
	sessionMock.On("Close").Return()

	var session Session = sessionMock
	var rows, _ = session.ScanMapSlice(query)

	for _, row := range rows {
		fmt.Printf("id = %v, name = %v\n", row["id"], row["name"])
	}

	session.Close()

	// Output: id = 1, name = alex
}
