package gouid

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"unsafe"
)

// GOUID is a byte slice.
type GOUID []byte

// Charsets for string gouids.
var (
	LowerCaseAlphaNum = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	UpperCaseAlphaNum = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	MixedCaseAlphaNum = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	LowerCaseAlpha    = []byte("abcdefghijklmnopqrstuvwxyz")
	UpperCaseAlpha    = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	MixedCaseAlpha    = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// String returns a string with the given size. Ids are
// made with cryptographically secure random bytes and
// umatch /^[a-z0-9]$/.
func String(size int, charset []byte) string {
	b := make([]byte, size, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		charCnt := byte(len(charset))
		index := (b[i] / (255 / charCnt)) % charCnt
		b[i] = charset[index]
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes returns cryptographically secure random bytes.
func Bytes(size int) GOUID {
	b := make([]byte, size, size)
	rand.Read(b)
	return b
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
