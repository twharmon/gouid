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

func TestByteArr16(t *testing.T) {
	id := gouid.ByteArr16()
	if len(id) != 16 {
		t.Error("lengh of id was not 16")
	}
	if id == gouid.ByteArr16() {
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

func TestByteArr16Marshal(t *testing.T) {
	m := map[string]interface{}{
		"id": gouid.GOUID16([16]byte{1}),
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	want := `{"id":"01000000000000000000000000000000"}`
	if string(b) != want {
		t.Errorf("bad json marshal: %s != %s", string(b), want)
	}
}

func TestByteArr16Unmarshal(t *testing.T) {
	m := make(map[string]gouid.GOUID16)
	err := json.Unmarshal([]byte(`{"id":"01000000000000000000000000000000"}`), &m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	a := m["id"]
	b := [16]byte{1}
	if bytes.Compare(a[:], b[:]) != 0 {
		t.Errorf("bad json unmarshal: %v != %v", m["id"], [16]byte{1})
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

func BenchmarkByteArr16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.ByteArr16()
	}
}

func BenchmarkBytes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.Bytes(32)
	}
}
