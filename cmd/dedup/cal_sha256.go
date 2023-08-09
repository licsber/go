package main

import (
	"crypto/sha256"
	"encoding/hex"
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
	length, err := io.Copy(h, f)
	if length != 0 {
		lErr.PanicErr(err)
	}

	return h.Sum(nil)
}

func CalSha256Hex(path string) string {
	return hex.EncodeToString(CalSha256(path))
}
