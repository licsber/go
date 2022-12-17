package lMeta

import (
	"encoding/json"
	"os"
)

func Decode(f *os.File) (m *Meta, err error) {
	m = &Meta{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(m)
	return
}
