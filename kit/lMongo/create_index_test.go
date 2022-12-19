package lMongo

import (
	"testing"
)

func create(keys []string) int {
	return len(keys)
}

func TestLen(t *testing.T) {
	res := create(nil)
	if res != 0 {
		t.Error()
	}
}

func TestCreateSimpleIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error()
		}
	}()
	CreateSimpleIndex(nil, []string{"1", "2", "3"})
}
