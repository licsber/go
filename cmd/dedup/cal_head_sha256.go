package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/licsber/go/lErr"
	"os"
)

// HeadSha256Size 1MiB
const HeadSha256Size = 1024 * 1024

func CalHeadSha256(path string) []byte {
	f, err := os.Open(path)
	lErr.PanicErr(err)

	defer func() {
		err = f.Close()
		lErr.PanicErr(err)
	}()

	buf := make([]byte, HeadSha256Size)
	_, err = f.Read(buf)
	lErr.PanicErr(err)

	h := sha256.New()
	h.Write(buf)
	return h.Sum(nil)
}

func CalHeadSha256Hex(path string) string {
	return hex.EncodeToString(CalSha256(path))
}
