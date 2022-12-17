package multi_hash

import (
	"encoding/hex"
	"strings"
)

func (mHash *MultiHash) Get(hashName string) string {
	return hex.EncodeToString(mHash.hashFuncMap[hashName].Sum(nil))
}

func (mHash *MultiHash) GetUpper(hashName string) string {
	return strings.ToUpper(mHash.Get(hashName))
}
