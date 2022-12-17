package multi_hash

func (mHash *MultiHash) AllMap() map[string]string {
	res := make(map[string]string)
	for hashName := range mHash.hashFuncMap {
		res[hashName] = mHash.Get(hashName)
	}
	return res
}

func (mHash *MultiHash) AllMapUpper() map[string]string {
	res := make(map[string]string)
	for hashName := range mHash.hashFuncMap {
		res[hashName] = mHash.GetUpper(hashName)
	}
	return res
}
