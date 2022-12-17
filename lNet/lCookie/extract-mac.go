package lCookie

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
)

// ExtractMac 复制请求
func ExtractMac(b []byte) error {
	sc := bufio.NewScanner(bytes.NewReader(b))
	for sc.Scan() {
		text := sc.Text()
		if strings.HasPrefix(text, "Authorization") {
			Authorization = text[len("Authorization: "):]
		} else if strings.HasPrefix(text, "Cookie") {
			Cookie = text[len("Cookie: "):]
		}
	}
	if Authorization == "" && Cookie == "" {
		return errors.New("extract cookie failed")
	}
	return nil
}
