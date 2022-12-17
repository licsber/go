package lFile

import (
	"os"
	"strings"
)

const macMetaStr = "This resource fork intentionally left blank"

func IsMacMetaFile(p string) bool {
	b, err := os.ReadFile(p)
	if err != nil {
		return false
	}
	return strings.Contains(string(b), macMetaStr)
}
