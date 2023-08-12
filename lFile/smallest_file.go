package lFile

import (
	"github.com/licsber/go/lErr"
	"math"
	"os"
)

func SmallestFile(filePaths []string) (res string) {
	size := int64(math.MaxInt64)
	for _, path := range filePaths {
		fi, err := os.Stat(path)
		lErr.PanicErr(err)

		fileSize := fi.Size()
		if fileSize < size {
			size = fileSize
			res = path
		}
	}

	return
}
