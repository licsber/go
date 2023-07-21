package lWebp

import "testing"

func TestEncode(t *testing.T) {
	b, err := Encode(nil)
	if err == nil {
		t.Error(b)
	}
}
