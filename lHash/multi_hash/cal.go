package multi_hash

import (
	"github.com/licsber/go/lFile"
	"io"
	"os"
)

func (mHash *MultiHash) Cal(p string, bufSize int) error {
	if bufSize <= 0 {
		bufSize = 4 * lFile.MB
	}

	f, err := os.Open(p)
	if err != nil {
		return err
	}

	buffer := make([]byte, bufSize)
	for {
		n, eof := f.Read(buffer)
		if eof == io.EOF {
			break
		}

		mHash.Write(buffer[:n])
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
