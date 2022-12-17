package multi_hash

import "github.com/licsber/go/lHash/ed2k"

func (mHash *MultiHash) Close() error {
	if e, ok := mHash.hashFuncMap["ED2K"]; ok {
		return e.(ed2k.HashCloser).Close()
	}

	return nil
}
