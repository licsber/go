package lErr

import (
	"fmt"
	_go "github.com/licsber/go"
)

func init() {
	fmt.Println("LicsberLib " + _go.VERSION)
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
