package gouid_test

import (
	"testing"

	"github.com/twharmon/gouid"
)

func TestNew(t *testing.T) {
	id := gouid.New()
	if len(id) != 32 {
		t.Error("lengh of id was not 32")
	}
	if id == gouid.New() {
		t.Error("collision")
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.New()
	}
}
