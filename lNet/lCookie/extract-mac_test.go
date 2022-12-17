package lCookie

import (
	"fmt"
	"os"
	"testing"
)

func TestExtractCookie(t *testing.T) {
	// TODO 增加Mac的测试用例.
	b, err := os.ReadFile("debug-mac.txt")
	if err != nil {
		t.Error(err)
	}

	err = ExtractMac(b)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Cookie Extract below:")
	fmt.Println(Cookie)
	fmt.Println("Authorization Extract below:")
	fmt.Println(Authorization)
}
