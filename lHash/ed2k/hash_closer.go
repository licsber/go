package ed2k

import (
	"hash"
	"io"
)

// HashCloser is a hash.Hash that also needs to be Close()d when done.
type HashCloser interface {
	hash.Hash
	io.Closer
}
