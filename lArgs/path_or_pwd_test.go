package lArgs

import (
	"testing"
)

func TestPathOrPWD(t *testing.T) {
	defer func() {
		recover()
	}()
	PathOrPWD()
}
