package gouid_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/twharmon/gouid"
)

func TestString(t *testing.T) {
	length := 32
	id := gouid.String(length)
	if len(id) != length {
		t.Error("lengh of id was not 32")
	}
	if id == gouid.String(length) {
		t.Error("collision")
	}
}

func TestBytes(t *testing.T) {
	length := 16
	id := gouid.Bytes(length)
	if len(id) != length {
		t.Error("lengh of id was not 16")
	}
	if bytes.Compare(id, gouid.Bytes(length)) == 0 {
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
	if bytes.Compare(m["id"], []byte{1}) != 0 {
		t.Errorf("bad json unmarshal: %v != %v", m["id"], []byte{1})
	}
}

func BenchmarkString8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(8)
	}
}

func BenchmarkString16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(16)
	}
}

func BenchmarkString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(32)
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
