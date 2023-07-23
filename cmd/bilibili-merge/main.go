package main

import (
	"fmt"
	"github.com/licsber/go/lArgs"
	"github.com/licsber/go/lErr"
	"os/exec"
	"path/filepath"
	"strings"
)

func FFMPEG(args []string) {
	cmd := exec.Command("ffmpeg", args...)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	lErr.PanicErr(err)
}

func Merge(videoFilepath, audioFilepath, assFilepath string) {
	mergeFilename := strings.Replace(filepath.Base(assFilepath), ".ass", ".mp4", 1)
	mergeFilePath := filepath.Join(filepath.Dir(assFilepath), mergeFilename)
	fmt.Println(mergeFilePath)

	args := []string{
		"-y",
		"-i", videoFilepath,
		"-i", audioFilepath,
		"-c", "copy",
		mergeFilePath,
	}
	FFMPEG(args)

	audioWMVName := strings.Replace(filepath.Base(assFilepath), ".ass", ".wmv", 1)
	audioWMVPath := filepath.Join(filepath.Dir(assFilepath), audioWMVName)
	args = []string{
		"-y",
		"-i", audioFilepath,
		"-c", "copy",
		audioWMVPath,
	}
	FFMPEG(args)
}

func main() {
	path, fi := lArgs.PathOrPWD()
	if !fi.IsDir() {
		fmt.Println("Must run with directory.")
		return
	}

	assPaths, err := filepath.Glob(filepath.Join(path, "*.ass"))
	lErr.PanicErr(err)

	if len(assPaths) != 1 {
		fmt.Println("Must has 1 ass file.")
		return
	}

	m4sPaths, err := filepath.Glob(filepath.Join(path, "*.m4s"))
	lErr.PanicErr(err)

	if len(m4sPaths) != 2 {
		fmt.Println("Must has 2 m4s file.")
		return
	}

	var videoFilepath, audioFilepath string

	for _, fp := range m4sPaths {
		if strings.Contains(fp, "_") {
			audioFilepath = fp
		} else {
			videoFilepath = fp
		}
	}

	Merge(videoFilepath, audioFilepath, assPaths[0])
}
