package main

import (
	"github.com/licsber/go/lErr"
	"io/fs"
	"path/filepath"
)

func BuildFileSizeMap(rootPath string) map[int64][]string {
	var res = make(map[int64][]string)

	err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		res[info.Size()] = append(res[info.Size()], path)

		return nil
	})
	lErr.PanicErr(err)

	return res
}
