package main

import (
	"fmt"
	"github.com/licsber/go/lEnv"
	"github.com/licsber/go/lErr"
	"github.com/licsber/go/lThirdParty/sct"
	"os"
	"strings"
)

var SCTKey string

func main() {
	if SCTKey == "" {
		env := lEnv.Parse([]string{"L_SCT_KEY"})
		SCTKey = env["L_SCT_KEY"]
	}

	switch len(os.Args) {
	case 1:
		fmt.Println("Usage: sct title [text]")
		host, err := os.Hostname()
		lErr.PanicErr(err)

		err = sct.SendPlainText(SCTKey, host, strings.Join(os.Environ(), "\n"))
		lErr.PanicErr(err)
	case 2:
		err := sct.SendPlainText(SCTKey, os.Args[1], os.Args[1])
		lErr.PanicErr(err)
	default:
		err := sct.SendPlainText(SCTKey, os.Args[1], strings.Join(os.Args[2:], "\n"))
		lErr.PanicErr(err)
	}
}
