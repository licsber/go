package lTime

import (
	"fmt"
	"testing"
	"time"
)

func TestLocalTimeStr(t *testing.T) {
	fmt.Println(LocalTimeStr(time.Now()))
}

func TestShanghaiTimeStr(t *testing.T) {
	fmt.Println(ShanghaiTimeStr(time.Now()))
}
