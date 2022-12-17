package tests

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		// [0, 5)
		fmt.Println(rand.Intn(5))
	}
}
