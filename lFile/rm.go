package lFile

import (
	"fmt"
	"os"
)

func Rm(p string) error {
	fmt.Print("rm: ")
	fmt.Println(p)
	return os.Remove(p)
}

func RmDir(p string) error {
	fmt.Print("rm: ")
	fmt.Println(p)
	return os.Remove(p)
}
