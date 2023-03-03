package lNet

import (
	"github.com/licsber/go/lFile"
	"github.com/licsber/go/lJson"
)

func Debug(b []byte) error {
	return lFile.Save("debug.html", b)
}

func DebugJSON(b []byte) error {
	return lJson.PrettySave(b, "debug.json")
}
