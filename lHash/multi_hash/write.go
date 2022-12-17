package multi_hash

import (
	"hash"
	"sync"
)

func (mHash *MultiHash) Write(p []byte) {
	wg := &sync.WaitGroup{}
	for _, h := range mHash.hashFuncMap {
		wg.Add(1)
		go func(h hash.Hash) {
			h.Write(p)
			wg.Done()
		}(h)
	}
	wg.Wait()
}
