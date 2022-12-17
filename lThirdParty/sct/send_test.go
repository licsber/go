package sct

import (
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	sctKey := os.Getenv("L_SCT_KEY")
	err := SendPlainText(sctKey, "测试", "测试\n测试")
	if err != nil {
		t.Error(err)
	}
}

func TestSendLong(t *testing.T) {
	sctKey := os.Getenv("L_SCT_KEY")
	err := SendPlainText(sctKey, "一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十", "测试\n测试")
	if err != nil {
		t.Error(err)
	}
}
