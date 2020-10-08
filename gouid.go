package gouid

import (
	"crypto/rand"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

// New returns a string with the given size. Ids are
// made with cryptographically secure random bytes and
// umatch /^[a-z0-9]$/.
func New(size int) string {
	b := make([]byte, size, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]/8]
	}
	return *(*string)(unsafe.Pointer(&b))
}
