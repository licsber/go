package lFile

import (
	"os"
	"strings"
)

const macMetaStr1 = "com.apple.quarantine"
const macMetaStr2 = "This resource fork intentionally left blank"

func IsMacMetaFile(p string) bool {
	b, err := os.ReadFile(p)
	if err != nil {
		return false
	}

	if strings.Contains(string(b), macMetaStr1) {
		return true
	}

	return strings.Contains(string(b), macMetaStr2)
}
