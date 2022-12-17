package lNet

import (
	"github.com/licsber/go/lJson"
	"os"
)

func Debug(b []byte) error {
	return os.WriteFile("debug.html", b, 0660)
}

func DebugJSON(b []byte) error {
	return lJson.PrettySave(b, "debug.json")
}
