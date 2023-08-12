package main

import (
	"fmt"
	"github.com/licsber/go/lErr"
	"os/exec"
)

func FFMPEG(args []string) {
	fmt.Print("FFMPEG: ")
	fmt.Println(args)
	cmd := exec.Command("ffmpeg", args...)
	out, err := cmd.CombinedOutput()
	if len(out) != 0 {
		fmt.Println(string(out))
	}

	lErr.PanicErr(err)
}
