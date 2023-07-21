package lNet

import "testing"

func TestJSONArray(t *testing.T) {
	a, err := JSONArray(nil, nil)
	if err == nil {
		t.Error(a)
	}
}
