package errors

import (
	"errors"
	"testing"
)

func TestErrorsAdd(t *testing.T) {
	e := Errors{}
	e = e.Add(errors.New("for testing"))
	if e.Len() != 1 {
		t.Logf("should have length 1, but was: %d", e.Len())
		t.Fail()
	}
}
