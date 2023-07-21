package lFile

import (
	"testing"
)

func TestIsMacMetaFile(t *testing.T) {
	ok := IsMacMetaFile("not exist.")
	if ok {
		t.Error("error.")
	}
}
