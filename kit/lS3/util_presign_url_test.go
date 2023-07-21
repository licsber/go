package lS3

import "testing"

func TestPreSignUrl(t *testing.T) {
	defer func() {
		recover()
	}()
	_, _ = PreSignUrl(nil, "", "", 1800)
}
