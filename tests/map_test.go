package tests

import (
	"fmt"
	"testing"
)

func f(a map[string]string) {
	a["123"] = "233"

}

func TestMap(t *testing.T) {
	m := make(map[string]string)
	f(m)
	// map是值传递 没错 但是传递的值是个地址 可以修改本体对象
	fmt.Println(m)
}
