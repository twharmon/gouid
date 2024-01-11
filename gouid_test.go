package gouid_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/twharmon/gouid"
)

// BenchmarkString8-10     	 2466236	       475.0 ns/op	       8 B/op	       1 allocs/op
// BenchmarkString16-10    	 2397400	       494.9 ns/op	      16 B/op	       1 allocs/op
// BenchmarkString32-10    	 2246272	       534.4 ns/op	      32 B/op	       1 allocs/op
// BenchmarkBytes8-10      	 2552016	       470.2 ns/op	       8 B/op	       1 allocs/op
// BenchmarkBytes16-10     	 2472471	       483.4 ns/op	      16 B/op	       1 allocs/op
// BenchmarkBytes32-10     	 2400997	       500.6 ns/op	      32 B/op	       1 allocs/op

func TestString(t *testing.T) {
	length := 32
	id := gouid.String(length, gouid.Secure64Char)
	if len(id) != length {
		t.Error("lengh of id was not 32")
	}
	if id == gouid.String(length, gouid.Secure64Char) {
		t.Error("collision")
	}
}

func TestStringSecure64Char(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[a-zA-Z0-9-_]{%d}", length)).Match([]byte(gouid.String(length, gouid.Secure64Char))) {
		t.Error("string did not match expected regexp")
	}
}

func TestStringSecure32Char(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[a-z0-5]{%d}", length)).Match([]byte(gouid.String(length, gouid.Secure32Char))) {
		t.Error("string did not match expected regexp")
	}
}

func TestBytes(t *testing.T) {
	length := 16
	id := gouid.Bytes(length)
	if len(id) != length {
		t.Error("lengh of id was not 16")
	}
	if bytes.Equal(id, gouid.Bytes(length)) {
		t.Error("collision")
	}
}

func TestBytesMarshal(t *testing.T) {
	m := map[string]interface{}{
		"id": gouid.GOUID([]byte{1}),
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	want := `{"id":"01"}`
	if string(b) != want {
		t.Errorf("bad json marshal: %s != %s", string(b), want)
	}
}

func TestBytesUnmarshal(t *testing.T) {
	m := make(map[string]gouid.GOUID)
	err := json.Unmarshal([]byte(`{"id":"01"}`), &m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !bytes.Equal(m["id"], []byte{1}) {
		t.Errorf("bad json unmarshal: %v != %v", m["id"], []byte{1})
	}
}

func BenchmarkString8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(8, gouid.Secure64Char)
	}
}

func BenchmarkString16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(16, gouid.Secure64Char)
	}
}

func BenchmarkString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(32, gouid.Secure64Char)
	}
}
func BenchmarkBytes8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.Bytes(8)
	}
}

func BenchmarkBytes16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.Bytes(16)
	}
}

func BenchmarkBytes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.Bytes(32)
	}
}
