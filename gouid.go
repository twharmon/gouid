// Package guoid provides cryptographically secure unique identifiers \
// of type string and type []byte.
//
// On Linux, FreeBSD, Dragonfly and Solaris, getrandom(2) is used if
// available, /dev/urandom otherwise.
// On OpenBSD and macOS, getentropy(2) is used.
// On other Unix-like systems, /dev/urandom is used.
// On Windows systems, the RtlGenRandom API is used.
// On Wasm, the Web Crypto API is used.
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

// String returns a string with the given size made up of
// characters from the given charset. Ids are made with
// cryptographically secure random bytes. The length of
// charset must not exceed 256.
func String(size int, charset []byte) string {
	b := make([]byte, size)
	randBytes(b)
	charCnt := byte(len(charset))
	for i := range b {
		b[i] = charset[b[i]%charCnt]
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes returns cryptographically secure random bytes.
func Bytes(size int) GOUID {
	b := make([]byte, size)
	randBytes(b)
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

func randBytes(buf []byte) {
	var n int
	var err error
	for n < len(buf) && err == nil {
		var nn int
		nn, err = rand.Reader.Read(buf[n:])
		n += nn
	}
}
