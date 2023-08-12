package main

import (
	"fmt"
	"github.com/licsber/go/lErr"
	"github.com/licsber/go/lFile"
	"path/filepath"
)

func guessVideoPath(rootPath string) string {
	m4sPaths, err := filepath.Glob(filepath.Join(rootPath, "*.m4s"))
	lErr.PanicErr(err)

	if len(m4sPaths) == 2 {
		return lFile.LargestFile(m4sPaths)
	}

	if len(m4sPaths) != 1 {
		fmt.Println(m4sPaths)
		panic("Cannot find video m4s file.")
	}

	mp4Paths, err := filepath.Glob(filepath.Join(rootPath, "*.mp4"))
	lErr.PanicErr(err)

	if len(mp4Paths) == 1 {
		return mp4Paths[0]
	}

	fmt.Println(m4sPaths)
	fmt.Println(mp4Paths)
	panic("Cannot find video mp4 file.")
}

func guessAudioPath(rootPath string) string {
	m4sPaths, err := filepath.Glob(filepath.Join(rootPath, "*.m4s"))
	lErr.PanicErr(err)

	if len(m4sPaths) == 1 {
		return m4sPaths[0]
	}

	if len(m4sPaths) == 2 {
		return lFile.SmallestFile(m4sPaths)
	}

	panic("Cannot find audio file.")
}

func guessAssPath(rootPath string) string {
	assPaths, err := filepath.Glob(filepath.Join(rootPath, "*.ass"))
	lErr.PanicErr(err)

	if len(assPaths) != 1 {
		fmt.Print("Has more than one ass file.")
		fmt.Println(assPaths)
	}

	return assPaths[0]
}
