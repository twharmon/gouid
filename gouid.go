package gouid

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"unsafe"
)

// GOUID is a byte slice.
type GOUID []byte

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
func Bytes(size int) GOUID {
	b := make([]byte, size, size)
	rand.Read(b)
	return b
}

// MarshalJSON .
func (g GOUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(g))
}

// UnmarshalJSON .
func (g *GOUID) UnmarshalJSON(data []byte) error {
	var x string
	err := json.Unmarshal(data, &x)
	if err == nil {
		str, e := hex.DecodeString(x)
		*g = GOUID([]byte(str))
		err = e
	}
	return err
}
