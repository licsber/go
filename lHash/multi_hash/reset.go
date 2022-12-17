package multi_hash

import (
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"github.com/licsber/go/lHash/ed2k"
	_ "github.com/licsber/go/lHash/md4"
	"hash"
)

func (mHash *MultiHash) Reset() {
	ed2kNeedClose := true
	if mHash.hashFuncMap == nil {
		ed2kNeedClose = false
		mHash.hashFuncMap = make(map[string]hash.Hash)
	}

	for i := range mHash.needHashes {
		hashName := mHash.needHashes[i]
		if hashName == "ED2K" {
			if ed2kNeedClose {
				err := mHash.hashFuncMap["ED2K"].(ed2k.HashCloser).Close()
				if err != nil {
					panic(err)
				}
			}

			mHash.hashFuncMap[hashName] = ed2k.New(false)
			continue
		}

		if hashFunc, ok := _supportHashes[hashName]; ok {
			mHash.hashFuncMap[hashName] = hashFunc.New()
		}
	}
}
