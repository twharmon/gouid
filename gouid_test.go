package gouid_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/twharmon/gouid"
)

func TestString(t *testing.T) {
	length := 32
	id := gouid.String(length, gouid.MixedCaseAlphaNum)
	if len(id) != length {
		t.Error("lengh of id was not 32")
	}
	if id == gouid.String(length, gouid.MixedCaseAlphaNum) {
		t.Error("collision")
	}
}

func TestStringLowerCaseAlpha(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[a-z]{%d}", length)).Match([]byte(gouid.String(length, gouid.LowerCaseAlpha))) {
		t.Error("string did not match expected regexp")
	}
}

func TestStringUpperCaseAlpha(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[A-Z]{%d}", length)).Match([]byte(gouid.String(length, gouid.UpperCaseAlpha))) {
		t.Error("string did not match expected regexp")
	}
}

func TestStringMixedCaseAlpha(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[a-zA-Z]{%d}", length)).Match([]byte(gouid.String(length, gouid.MixedCaseAlpha))) {
		t.Error("string did not match expected regexp")
	}
}

func TestStringLowerCaseAlphaNum(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[a-z0-9]{%d}", length)).Match([]byte(gouid.String(length, gouid.LowerCaseAlphaNum))) {
		t.Error("string did not match expected regexp")
	}
}

func TestStringUpperCaseAlphaNum(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[A-Z0-9]{%d}", length)).Match([]byte(gouid.String(length, gouid.UpperCaseAlphaNum))) {
		t.Error("string did not match expected regexp")
	}
}

func TestStringMixedCaseAlphaNum(t *testing.T) {
	length := 256
	if !regexp.MustCompile(fmt.Sprintf("[a-zA-Z0-9]{%d}", length)).Match([]byte(gouid.String(length, gouid.MixedCaseAlphaNum))) {
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
		gouid.String(8, gouid.MixedCaseAlphaNum)
	}
}

func BenchmarkString16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(16, gouid.MixedCaseAlphaNum)
	}
}

func BenchmarkString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gouid.String(32, gouid.MixedCaseAlphaNum)
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
