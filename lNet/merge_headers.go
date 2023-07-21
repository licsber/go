package lNet

import (
	"net/http"
)

func MergeHeaders(a, b http.Header) {
	for k, v := range b {
		a[k] = v
	}
}
