package tests

import (
	"fmt"
	"testing"
)

func TestFloatFormat(t *testing.T) {
	// !!!! float前面数字的长度是总长度 包含小数点了
	fmt.Printf("%07.2f", 123.0)
	// 0123.00
}
