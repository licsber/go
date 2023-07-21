package lJson

import "testing"

func TestPrettyPrint(t *testing.T) {
	err := PrettyPrint(nil)
	if err == nil {
		t.Error("nil print.")
	}
}
