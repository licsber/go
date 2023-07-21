package lMeta

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	m, err := Decode(nil)
	if err == nil {
		fmt.Println(m)
		t.Error(t)
	}
}
