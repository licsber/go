package lFile

import "testing"

func TestExists(t *testing.T) {
	ok := Exists("exists.go")
	if !ok {
		t.Error("exists.go not exist.")
	}
}
