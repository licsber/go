package main

import (
	"flag"
)

var ForceFlag = flag.Bool("f", false, "Force(no prompt)")

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		panic("usage: dedup srcPath comparePath")
	}

	DeDup(flag.Args()[0], flag.Args()[1])
}
