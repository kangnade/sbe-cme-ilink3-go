package sbe

import "bytes"

// clean method finds the first null byte in a ilink message field, slices it there, and returns it in string
func clean(b []byte) string{
	// func IndexByte(b []byte, c byte) int
	// uses the bytes.IndexByte to return the index of the first instance of c (0 null) in b, or -1 if c is not present in b.
	n := bytes.IndexByte(b, 0)
	if n == -1{
		return string(b)
	}
	return string(b[:n])
}