package main

import (
	"strings"
	"testing"
)

var src = strings.Repeat("a", 1023) + "中"
var dest = make([]byte, 1024)

func BenchmarkTestCopyString2(b *testing.B) {
	for range b.N {
		CopyString2(dest, src)
	}
}

func BenchmarkTestCopyString(b *testing.B) {
	for range b.N {
		CopyString(dest, src)
	}
}

func TestCopyString(t *testing.T) {
	src := "abcd"
	var dest = make([]byte, 4)
	n := CopyString(dest, src)
	if n != 4 || string(dest[:n]) != "abcd" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "abc"
	n = CopyString(dest, src)
	if n != 3 || string(dest[:n]) != "abc" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "abcdef"
	n = CopyString(dest, src)
	if n != 4 || string(dest[:n]) != "abcd" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "abc中"
	n = CopyString(dest, src)
	if n != 3 || string(dest[:n]) != "abc" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "ab中"
	n = CopyString(dest, src)
	if n != 2 || string(dest[:n]) != "ab" {
		t.Fatal(src)
	}
}

func TestCopyString2(t *testing.T) {
	src := "abcd"
	var dest = make([]byte, 4)
	n := CopyString(dest, src)
	if n != 4 || string(dest[:n]) != "abcd" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "abc"
	n = CopyString2(dest, src)
	if n != 3 || string(dest[:n]) != "abc" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "abcdef"
	n = CopyString2(dest, src)
	if n != 4 || string(dest[:n]) != "abcd" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "abc中"
	n = CopyString2(dest, src)
	if n != 3 || string(dest[:n]) != "abc" {
		t.Fatal(src)
	}

	clear(dest[:])
	src = "ab中"
	n = CopyString2(dest, src)
	if n != 2 || string(dest[:n]) != "ab" {
		t.Fatal(src)
	}
}
