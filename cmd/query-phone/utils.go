package main

import "errors"

func get4(b []byte) int32 {
	if len(b) < 4 {
		return 0
	}
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

func getN(s string) (uint32, error) {
	var n, cutoff, maxVal uint32
	i := 0
	base := 10
	cutoff = (1<<32-1)/10 + 1
	maxVal = 1<<uint(32) - 1
	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			return 0, errors.New("invalid syntax")
		}

		if v >= byte(base) {
			return 0, errors.New("invalid syntax")
		}

		if n >= cutoff {
			// n*base overflows
			n = 1<<32 - 1
			return n, errors.New("value out of range")
		}

		n *= uint32(base)

		n1 := n + uint32(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = 1<<32 - 1
			return n, errors.New("value out of range")
		}

		n = n1
	}

	return n, nil
}
