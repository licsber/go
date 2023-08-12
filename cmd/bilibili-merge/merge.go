package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Merge(videoFilepath, audioFilepath, assFilepath string) {
	mergeFilename := strings.Replace(filepath.Base(assFilepath), ".ass", ".mp4", 1)
	mergeFilePath := filepath.Join(filepath.Dir(assFilepath), mergeFilename)
	fmt.Print("Merge: ")
	fmt.Println(mergeFilePath)

	args := []string{
		"-y",
		"-loglevel", "error",
		"-hide_banner",
		"-i", videoFilepath,
		"-i", audioFilepath,
		"-c", "copy",
		mergeFilePath,
	}
	FFMPEG(args)

	audioAACName := strings.Replace(filepath.Base(assFilepath), ".ass", ".aac", 1)
	audioAACPath := filepath.Join(filepath.Dir(assFilepath), audioAACName)
	args = []string{
		"-y",
		"-loglevel", "error",
		"-hide_banner",
		"-i", audioFilepath,
		"-c", "copy",
		audioAACPath,
	}
	FFMPEG(args)
}
