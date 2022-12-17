package lFile

import "os"

func Exists(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return false
	}

	return true
}
