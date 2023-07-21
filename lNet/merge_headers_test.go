package lNet

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMergeHeaders(t *testing.T) {
	a := http.Header{
		"a": []string{"a"},
	}
	b := http.Header{
		"b": []string{"b"},
	}
	MergeHeaders(a, b)
	fmt.Println(a)
}
