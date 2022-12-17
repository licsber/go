package tests

import (
	"fmt"
	"github.com/licsber/go/lEnv"
	"testing"
)

func TestImport(t *testing.T) {
	fmt.Print(lEnv.Get("123", "123"))
}
