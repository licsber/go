// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package md4 implements the MD4 hash algorithm as defined in RFC 1320.
//
// Deprecated: MD4 is cryptographically broken and SHOULD ONLY be used
// where compatibility with legacy systems, not security, is the goal. Instead,
// use a secure hash like SHA-256 (from crypto/sha256).

package md4

import (
	"crypto"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.MD4, New)
}

// Size The size of an MD4 checksum in bytes.
const Size = 16

// BlockSize The block size of MD4 in bytes.
const BlockSize = 64

const (
	_Chunk = 64
	_Init0 = 0x67452301
	_Init1 = 0xEFCDAB89
	_Init2 = 0x98BADCFE
	_Init3 = 0x10325476
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	s   [4]uint32
	x   [_Chunk]byte
	nx  int
	len uint64
}

func (d *digest) Reset() {
	d.s[0] = _Init0
	d.s[1] = _Init1
	d.s[2] = _Init2
	d.s[3] = _Init3
	d.nx = 0
	d.len = 0
}

// New returns a new hash.Hash computing the MD4 checksum.
func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 {
		n := len(p)
		if n > _Chunk-d.nx {
			n = _Chunk - d.nx
		}
		for i := 0; i < n; i++ {
			d.x[d.nx+i] = p[i]
		}
		d.nx += n
		if d.nx == _Chunk {
			_Block(d, d.x[0:])
			d.nx = 0
		}
		p = p[n:]
	}
	n := _Block(d, p)
	p = p[n:]
	if len(p) > 0 {
		d.nx = copy(d.x[:], p)
	}
	return
}

func (d *digest) Sum(in []byte) []byte {
	// Make a copy of d0, so that caller can keep writing and summing.
	d0 := new(digest)
	*d0 = *d

	// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.
	l := d0.len
	var tmp [64]byte
	tmp[0] = 0x80
	if l%64 < 56 {
		_, err := d0.Write(tmp[0 : 56-l%64])
		if err != nil {
			panic("md4: write error")
		}
	} else {
		_, err := d0.Write(tmp[0 : 64+56-l%64])
		if err != nil {
			panic("md4: write error")
		}
	}

	// Length in bits.
	l <<= 3
	for i := uint(0); i < 8; i++ {
		tmp[i] = byte(l >> (8 * i))
	}
	_, err := d0.Write(tmp[0:8])
	if err != nil {
		panic("md4: write error")
	}

	if d0.nx != 0 {
		panic("d0.nx != 0")
	}

	for _, s := range d0.s {
		in = append(in, byte(s>>0))
		in = append(in, byte(s>>8))
		in = append(in, byte(s>>16))
		in = append(in, byte(s>>24))
	}
	return in
}
