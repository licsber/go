package lFile

import (
	"testing"
)

func TestFriendlySize(t *testing.T) {
	testCases := map[int64]string{
		0:          "0.00B",
		100:        "100.00B",
		1024:       "1.00KB",
		99999:      "97.66KB",
		1024000000: "976.56MB",
		-256:       "0.00B",
		-1:         "0.00B",
	}

	for k, v := range testCases {
		if FriendlySize(k) != v {
			t.Error(k, "except", v, "got", FriendlySize(k))
		}
	}
}
