package lFile

import (
	"github.com/licsber/go/lErr"
	"math"
	"os"
)

func LargestFile(filePaths []string) (res string) {
	size := int64(math.MinInt64)
	for _, path := range filePaths {
		fi, err := os.Stat(path)
		lErr.PanicErr(err)

		fileSize := fi.Size()
		if fileSize > size {
			size = fileSize
			res = path
		}
	}

	return
}
