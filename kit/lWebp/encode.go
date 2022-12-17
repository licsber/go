package lWebp

import (
	"bytes"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"log"
)

var Opts *encoder.Options

func init() {
	opts, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	Opts = opts
}

func Encode(img image.Image) (b []byte, err error) {
	var res bytes.Buffer
	err = webp.Encode(&res, img, Opts)
	if err != nil {
		return
	}

	b = res.Bytes()
	return
}
