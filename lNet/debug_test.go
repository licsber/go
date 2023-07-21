package lNet

import (
	"testing"
)

func TestDebug(t *testing.T) {
	text := []byte("hello")
	err := Debug(text)
	if err != nil {
		t.Error(err)
	}
}

func TestDebugJSON(t *testing.T) {
	jsonText := []byte(`[
		{"key": 233}
	]`)
	err := DebugJSON(jsonText)
	if err != nil {
		t.Error(err)
	}
}
