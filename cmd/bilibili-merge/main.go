package main

import (
	"fmt"
	"github.com/licsber/go/lArgs"
)

func main() {
	path, fi := lArgs.PathOrPWD()
	if !fi.IsDir() {
		fmt.Println("Must run with directory.")
		return
	}

	videoFilePath := guessVideoPath(path)
	audioFilePath := guessAudioPath(path)
	assFilePath := guessAssPath(path)

	fmt.Print("Video: ")
	fmt.Println(videoFilePath)
	fmt.Print("Audio: ")
	fmt.Println(audioFilePath)
	fmt.Print("Ass: ")
	fmt.Println(assFilePath)
	Merge(videoFilePath, audioFilePath, assFilePath)
}
