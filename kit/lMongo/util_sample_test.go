package lMongo

import "testing"

func TestSample(t *testing.T) {
	defer func() {
		recover()
	}()
	_, _ = Sample(nil, 0, nil, nil)
}
