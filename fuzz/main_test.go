package demo

import (
	"strings"
	"testing"
)

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func TestUpper(t *testing.T) {
	have := "hello"
	expect := "HELLO"
	got := toUpper(have)
	t.Log(got, expect)
	if expect != got {
		t.Fail()
	}
}

func FuzzToUper(f *testing.F) {
	f.Add("hello")
	f.Fuzz(func(t *testing.T, s string) {
		out := toUpper(s)
		if out != strings.ToUpper(s) {
			t.Fail()
		}
	})
}
