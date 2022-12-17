package lMeta

import (
	"github.com/licsber/go/lErr"
	"github.com/licsber/go/lFile"
	"github.com/licsber/go/lHash/cloud_hash"
	"github.com/licsber/go/lHash/multi_hash"
	"github.com/licsber/go/lTime"
	"os"
)

var _ReadFileBufferSize = 4 * lFile.MB

func CalFile(p string) (m *Meta) {
	m = &Meta{}
	info, err := os.Stat(p)
	if err != nil {
		panic("lMeta.CalFile.Stat: " + p + " file not exist?")
	}

	m.BaseName = info.Name()
	m.Size = info.Size()
	m.FriendlySize = lFile.FriendlySize(m.Size)
	m.ModTime = lTime.TimeStampSec(info.ModTime())

	m.Head115, _ = cloud_hash.Cal115SHA1(p)
	m.HeadBaidu, _ = cloud_hash.CalBaiduMD5(p)

	mHash := multi_hash.New(_RequiredHashes)
	err = mHash.Cal(p, _ReadFileBufferSize)
	lErr.PanicErr(err)

	m.ED2K = mHash.GetUpper("ED2K")
	m.MD5 = mHash.GetUpper("MD5")
	m.SHA1 = mHash.GetUpper("SHA1")
	m.SHA256 = mHash.GetUpper("SHA256")

	err = mHash.Close()
	if err != nil {
		panic("lMeta.CalFile.Close: close MultiHash error.")
	}

	return
}
