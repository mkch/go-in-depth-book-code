package main

import (
	"math"
	"testing"
)

func TestAvgArea(t *testing.T) {
	var want = (float32(1)*2 + math.Pi*3*3) / 2
	if got := AvgArea([]Shape{&Rectangle{1, 2}, &Circle{3}}); got != want {
		t.Fatalf("want: %v got: %v", want, got)
	}
}
