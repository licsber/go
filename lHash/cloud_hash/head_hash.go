package cloud_hash

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	"encoding/hex"
	"github.com/licsber/go/lFile"
	"io"
	"os"
	"strings"
)

func Cal115SHA1(p string) (string, error) {
	return calHeadHash(p, 128*lFile.KB, crypto.SHA1)
}

func CalBaiduMD5(p string) (string, error) {
	return calHeadHash(p, 256*lFile.KB, crypto.MD5)
}

func calHeadHash(
	p string,
	blockSize int,
	hashFunc crypto.Hash,
) (string, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}

	hashObj := hashFunc.New()
	block := make([]byte, blockSize)

	_, err = f.Read(block)
	if err != nil && err != io.EOF {
		return "", err
	}

	err = f.Close()
	if err != nil {
		return "", err
	}

	hashObj.Write(block)
	bytes := hashObj.Sum(nil)
	hexString := hex.EncodeToString(bytes)
	return strings.ToUpper(hexString), nil
}
