package multi_hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/licsber/go/lHash/ed2k"
	"testing"
)

var testBytes = []byte("licsber")

func TestMultiHash(t *testing.T) {
	m := New(nil)
	m.Write(testBytes)
	fmt.Println(m.AllMapUpper())
}

func BenchmarkMultiHash_Write(b *testing.B) {
	m := New([]string{"ED2K", "MD5", "SHA1", "SHA256"})
	bytes := make([]byte, 100*1024*1024)
	b.SetBytes(int64(len(bytes)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Write(bytes)
		m.AllMapUpper()
	}
}
func TestMD5(t *testing.T) {
	m := New([]string{"MD5"})
	m.Write(testBytes)

	md5Obj := md5.New()
	md5Obj.Write(testBytes)
	md5Sum := fmt.Sprintf("%x", md5Obj.Sum(nil))

	if m.Get("MD5") != md5Sum {
		t.Errorf("Error, 'licsber' MD5 got %s .", md5Sum)
	}
}

func TestSmallMD4(t *testing.T) {
	// 小文件的ED2K就是其MD4
	m := New([]string{"ED2K"})
	m.Write(testBytes)
	m.Reset()
	m.Write(testBytes)

	e := ed2k.New(false)
	_, err := e.Write(testBytes)
	if err != nil {
		panic("ed2k: cannot write.")
	}

	md4Str := m.Get("ED2K")
	ed2kStr := hex.EncodeToString(e.Sum(nil))
	if md4Str != ed2kStr {
		t.Errorf("%s != %s", ed2kStr, md4Str)
	}
}

func TestCloseNone(t *testing.T) {
	m := New(nil)
	_ = m.Close()
}
