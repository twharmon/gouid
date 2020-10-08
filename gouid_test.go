package gouid_test

import (
	"testing"

	"github.com/twharmon/gouid"
)

func TestNew(t *testing.T) {
	length := 32
	id := gouid.New(length)
	if len(id) != length {
		t.Error("lengh of id was not 32")
	}
	if id == gouid.New(length) {
		t.Error("collision")
	}
}

func BenchmarkNew8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.New(8)
	}
}

func BenchmarkNew16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.New(16)
	}
}

func BenchmarkNew32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.New(32)
	}
}
