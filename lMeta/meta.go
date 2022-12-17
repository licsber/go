package lMeta

import (
	"github.com/licsber/go/lTime"
)

var _RequiredHashes = []string{"ED2K", "MD5", "SHA1", "SHA256"}

type Meta struct {
	BaseName     string             `json:"basename"`
	Size         int64              `json:"size"`
	FriendlySize string             `json:"friendly_size"`
	ModTime      lTime.TimeStampSec `json:"mtime"`

	Head115   string `json:"head_115"`
	HeadBaidu string `json:"head_baidu"`
	ED2K      string `json:"ed2k"`

	MD5    string `json:"md5"`
	SHA1   string `json:"sha1"`
	SHA256 string `json:"sha256"`
}
