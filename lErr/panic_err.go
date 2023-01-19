package lErr

import (
	"fmt"
	_go "github.com/licsber/go"
	"os"
)

func init() {
	if os.Getenv("LICSBER") == "" {
		fmt.Println("LicsberLib " + _go.VERSION)
	}
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
