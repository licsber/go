package lJson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/licsber/go/lFile"
	"strings"
)

var Prefix = ""
var Indent = strings.Repeat(" ", 4)

func Pretty(b []byte) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	err = json.Indent(buf, b, Prefix, Indent)
	return
}

func PrettyMarshal(v any) (b []byte, err error) {
	b, err = json.MarshalIndent(v, Prefix, Indent)
	return
}

func PrettyPrint(b []byte) error {
	buf, err := Pretty(b)
	if err != nil {
		return err
	}

	fmt.Println(buf.String())
	return nil
}

func PrettySave(b []byte, savePath string) error {
	buf, err := Pretty(b)
	if err != nil {
		return err
	}

	return lFile.Save(savePath, buf.Bytes())
}
