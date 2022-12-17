package lMeta

import (
	"github.com/licsber/go/lJson"
)

func (meta *Meta) JSON() []byte {
	b, _ := lJson.PrettyMarshal(meta)
	return b
}

func (meta *Meta) JSONString() string {
	return string(meta.JSON())
}
