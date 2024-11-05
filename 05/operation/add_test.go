package operation_test

import (
	"golang_course_one/05/operation"
	"testing"
)

func TestAdd(t *testing.T) {
	got := operation.Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
