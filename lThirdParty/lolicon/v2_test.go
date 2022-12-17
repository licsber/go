package lolicon

import (
	"encoding/json"
	"fmt"
	"github.com/licsber/go/lErr"
	"github.com/licsber/go/lNet"
	"testing"
)

func TestV2(t *testing.T) {
	b, err := json.Marshal(V2Config)
	lErr.PanicErr(err)
	res := V2(&lNet.Client{}, b)
	fmt.Print(res)
}
