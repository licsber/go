package multi_hash

import (
	"crypto"
	"hash"
)

var _supportHashes = map[string]crypto.Hash{
	"ED2K":   65535, // 例外实现
	"MD4":    crypto.MD4,
	"MD5":    crypto.MD5,
	"SHA1":   crypto.SHA1,
	"SHA256": crypto.SHA256,
	"SHA512": crypto.SHA512,
}

type MultiHash struct {
	hashFuncMap map[string]hash.Hash
	needHashes  []string
}

func New(hashes []string) *MultiHash {
	if len(hashes) == 0 {
		hashes = []string{"MD5", "SHA1"}
	}

	m := &MultiHash{needHashes: hashes}
	m.Reset()
	return m
}
