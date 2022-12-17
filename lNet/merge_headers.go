package lNet

import "net/http"

func MergeHeaders(a, b *http.Header) {
	for k := range *b {
		v := b.Get(k)
		a.Set(k, v)
	}
}
