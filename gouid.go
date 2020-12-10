package gouid

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"unsafe"
)

// GOUID is a byte slice.
type GOUID []byte

// GOUID16 is a byte array with 16 bytes.
type GOUID16 [16]byte

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

// ByteArr16 returns 16 cryptographically secure random bytes.
func ByteArr16() GOUID16 {
	var ba [16]byte
	rand.Read(ba[:])
	return ba
}

// MarshalJSON hex encodes the gouid.
func (g GOUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// String implements the Stringer interface.
func (g GOUID) String() string {
	return hex.EncodeToString(g)
}

// UnmarshalJSON decodes a hex encoded string into a gouid.
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

// MarshalJSON hex encodes the gouid.
func (g GOUID16) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// String implements the Stringer interface.
func (g GOUID16) String() string {
	return hex.EncodeToString(g[:])
}

// UnmarshalJSON decodes a hex encoded string into a gouid.
func (g *GOUID16) UnmarshalJSON(data []byte) error {
	var x string
	err := json.Unmarshal(data, &x)
	if err == nil {
		str, e := hex.DecodeString(x)
		var ba [16]byte
		copy(ba[:], []byte(str))
		*g = ba
		err = e
	}
	return err
}
