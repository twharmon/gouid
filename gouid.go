package gouid

import (
	"crypto/rand"
	"encoding/base32"
)

// New returns a string with 32 characters. Ids are 20
// cryptographically secure random bytes encoded with
// standard base32 encoding.
func New() string {
	b := make([]byte, 20)
	rand.Read(b)
	return base32.StdEncoding.EncodeToString(b)
}
