package main

import (
	"crypto/sha256"
	"github.com/licsber/go/lErr"
	"io"
	"os"
)

func CalSha256(path string) []byte {
	f, err := os.Open(path)
	lErr.PanicErr(err)

	defer func() {
		err = f.Close()
		lErr.PanicErr(err)
	}()

	h := sha256.New()
	_, err = io.Copy(h, f)
	lErr.PanicErr(err)

	return h.Sum(nil)
}
