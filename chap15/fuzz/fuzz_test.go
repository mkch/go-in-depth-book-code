package fuzz

import (
	"testing"
	"unicode/utf8"
)

func TestSplitLast(t *testing.T) {
	tests := []struct {
		name      string
		arg1      string
		arg2      byte
		wantLeft  string
		wantRight string
	}{
		{"dot", "abc-=&^%d-efg.", '.', "abc-=&^%d-efg", ""},
		{"dot2", "abcdef.g", '.', "abcdef", "g"},
		{"not-found", "abcdefg", '|', "abcdefg", ""},
		{"EMPTY", "", '-', "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := SplitLast(tt.arg1, tt.arg2)
			if gotLeft != tt.wantLeft {
				t.Errorf("SplitLast() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if gotRight != tt.wantRight {
				t.Errorf("SplitLast() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}

func FuzzSplitLast(f *testing.F) {
	f.Add("abc-=&^%d-efg.", byte('.'))
	f.Add("abcdef.g", byte('.'))
	f.Add("abcdefg", byte('|'))
	f.Add("", byte('-'))
	f.Fuzz(func(t *testing.T, str string, del byte) {
		if !utf8.ValidString(str) {
			t.Skip("not valid input")
		}
		l, r := SplitLast(str, del)
		if !utf8.ValidString(l) {
			t.Fatalf("str=%v(%v), del=%v, left=%v(%v) not a valid string", str, []byte(str), del, l, []byte(l))
		}
		if !utf8.ValidString(r) {
			t.Fatalf("str=%v(%v), del=%v, right=%v(%v) not a valid string", str, []byte(str), del, r, []byte(r))
		}
	})
}
