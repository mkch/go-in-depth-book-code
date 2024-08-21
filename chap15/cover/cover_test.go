package cover_test

import (
	"cover"
	"testing"
)

func TestEllipsis(t *testing.T) {
	if ret := cover.Ellipsis("abc", 4); ret != "abc" {
		t.Fatalf("got %v, want %v", ret, "abc")
	}
	if ret := cover.Ellipsis("abc", 2); ret != "a…" {
		t.Fatalf("got %v, want %v", ret, "a…")
	}
	if ret := cover.Ellipsis("abc", 1); ret != "abc" {
		t.Fatalf("got %v, want %v", ret, "abc")
	}
}

func TestF(t *testing.T) {
	if cover.F(true) != "a" {
		t.Fail()
	}
	cover.F(false)
}
