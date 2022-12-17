package lCookie

import (
	"fmt"
	"os"
	"testing"
)

func TestExtractWinCookie(t *testing.T) {
	b, err := os.ReadFile("debug-win.txt")
	if err != nil {
		t.Error(err)
	}

	err = ExtractWin(b)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Cookie Extract below:")
	fmt.Println(Cookie)
	fmt.Println("Authorization Extract below:")
	fmt.Println(Authorization)
}
