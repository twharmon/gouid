package gouid_test

import (
	"bytes"
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
	t.Fatal(id)
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
