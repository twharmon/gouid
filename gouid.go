package gouid

import (
	"crypto/rand"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

// String returns a string with the given size. Ids are
// made with cryptographically secure random bytes and
// umatch /^[a-z0-9]$/.
func String(size int) string {
	b := make([]byte, size, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]/8]
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes returns cryptographically secure random bytes.
func Bytes(size int) []byte {
	b := make([]byte, size, size)
	rand.Read(b)
	return b
}
