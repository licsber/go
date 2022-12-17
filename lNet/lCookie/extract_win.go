package lCookie

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ExtractWin 复制 - 复制为Node.js fetch
func ExtractWin(b []byte) error {
	first := bytes.Index(b, []byte("\"headers\": {")) + len("headers\": ")
	last := bytes.Index(b, []byte("}")) + 1
	b = b[first:last]

	var j map[string]interface{}
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	Authorization, _ = j["Authorization"].(string)
	Cookie, _ = j["cookie"].(string)
	if Authorization == "" && Cookie == "" {
		return errors.New("extract cookie failed")
	}
	return nil
}
