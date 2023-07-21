package lFile

import (
	"testing"
)

func TestRm(t *testing.T) {
	err := Rm("debug.txt")
	if err == nil {
		t.Error("rm debug.txt .")
	}
}

func TestRmDir(t *testing.T) {
	err := RmDir("debug")
	if err == nil {
		t.Error("rmdir debug.")
	}
}
