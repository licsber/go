package lFile

import (
	"os"
	"path/filepath"
)

func Save(p string, b []byte) error {
	err := os.MkdirAll(filepath.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(p, b, os.ModePerm)
}
