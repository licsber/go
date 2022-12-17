package lFile

import (
	"fmt"
	"github.com/licsber/go/lErr"
	"os"
)

func Rm(p string) {
	fmt.Print("rm: ")
	fmt.Println(p)
	err := os.Remove(p)
	lErr.PanicErr(err)
}

func RmDir(p string) {
	fmt.Print("rm: ")
	fmt.Println(p)
	err := os.Remove(p)
	lErr.PanicErr(err)
}
