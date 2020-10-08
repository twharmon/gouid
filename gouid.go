package gouid

import (
	"crypto/rand"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

// New returns a string with the given size. Ids are
// made with cryptographically secure random bytes.
func New(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		index := b[i] / 8
		letter := alphabet[index]
		b[i] = letter
	}
	return *(*string)(unsafe.Pointer(&b))
}
