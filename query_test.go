package gockle

import (
	"context"
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	var s = newSession(t)
	defer s.Close()
	var exec = func(q string) {
		if err := s.Exec(q); err != nil {
			t.Fatalf("Actual error %v, expected no error", err)
		}
	}
	exec(ksDropIf)
	exec(ksCreate)
	defer exec(ksDrop)
	exec(tabCreate)
	defer exec(tabDrop)
	exec(rowInsert)
	q := s.Query("select * from gockle_test.test").
		PageSize(3).
		WithContext(context.Background()).
		PageState(nil)
	defer q.Release()
	if err := q.Exec(); err != nil {
		t.Fatalf("Actual error %v, expected no error", err)
	}
	// Scan
	var id, n int
	if err := q.Scan(&id, &n); err == nil {
		if id != 1 {
			t.Errorf("Actual id %v, expected 1", id)
		}
		if n != 2 {
			t.Errorf("Actual n %v, expected 2", n)
		}
	} else {
		t.Errorf("Actual error %v, expected no error", err)
	}
	// MapScan
	var actual = map[string]interface{}{}
	var expected = map[string]interface{}{"id": 1, "n": 2}
	if err := q.MapScan(actual); err == nil {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Actual map %v, expected %v", actual, expected)
		}
	} else {
		t.Errorf("Actual error %v, expected no error", err)
	}
}
