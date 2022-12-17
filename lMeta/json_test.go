package lMeta

import (
	"fmt"
	"testing"
)

func TestMeta_JSONString(t *testing.T) {
	meta := &Meta{}
	fmt.Println(meta.JSONString())
}
