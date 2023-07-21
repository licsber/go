package lTime

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeCost(t *testing.T) {
	_startTime := time.Now()
	defer func() {
		recover()
	}()
	defer func() {
		fmt.Print("Time cost: ")
		fmt.Println(time.Now().Sub(_startTime))
	}()
	TimeCost()
}
