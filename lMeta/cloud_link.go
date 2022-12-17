package lMeta

import "fmt"

type CloudLink interface {
	LinkED2K() string
	Link115() string
	LinkBaidu() string
	LinkLicsber() string
}

func (meta *Meta) LinkED2K() string {
	return fmt.Sprintf("ed2k://|file|%s|%d|%s|/", meta.BaseName, meta.Size, meta.ED2K)
}

func (meta *Meta) Link115() string {
	return fmt.Sprintf("115://%s|%d|%s|%s", meta.BaseName, meta.Size, meta.SHA1, meta.Head115)
}

func (meta *Meta) LinkBaidu() string {
	return fmt.Sprintf("%s#%s#%d#%s", meta.MD5, meta.HeadBaidu, meta.Size, meta.BaseName)
}

func (meta *Meta) LinkLicsberV1() string {
	return fmt.Sprintf("licsber://v1/%d/%s/%s/%s", meta.Size, meta.SHA1, meta.MD5, meta.BaseName)
}
