package lMeta

import (
	"os"
	"testing"
)

const emptyED2KLink = "ed2k://|file|empty|0|31D6CFE0D16AE931B73C59D7E0C089C0|/"

//goland:noinspection SpellCheckingInspection
const empty115Link = "115://empty|0|DA39A3EE5E6B4B0D3255BFEF95601890AFD80709|67DFD19F3EB3649D6F3F6631E44D0BD36B8D8D19"
const emptyBaiduLink = "D41D8CD98F00B204E9800998ECF8427E#EC87A838931D4D5D2E94A04644788A55#0#empty"

//goland:noinspection SpellCheckingInspection
const emptyLicsberLink = "licsber://v1/0/DA39A3EE5E6B4B0D3255BFEF95601890AFD80709/D41D8CD98F00B204E9800998ECF8427E/empty"

func TestEmptyMetaLinkGen(t *testing.T) {
	emptyFile, _ := os.CreateTemp("", "")
	meta := CalFile(emptyFile.Name())
	meta.BaseName = "empty"

	flag := true
	flag = flag && meta.LinkED2K() == emptyED2KLink
	flag = flag && meta.Link115() == empty115Link
	flag = flag && meta.LinkBaidu() == emptyBaiduLink
	flag = flag && meta.LinkLicsberV1() == emptyLicsberLink
	if !flag {
		t.Error(meta)
	}
}
