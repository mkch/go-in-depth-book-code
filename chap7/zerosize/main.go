package main

import (
	"fmt"
)

var s01 struct{}
var s02 struct{}

func main() {

	var a01 [0]byte
	var a02 [0]byte
	var s1 struct {
		F1 int
		F2 struct{}
		F3 struct{}
	}

	fmt.Printf("&s01:%#p &s02:%#p (&s01==&s02):%v\n", &s01, &s02, &s01 == &s02)
	fmt.Printf("&a01:%#p &a02:%#p (&a01==&a02):%v\n", &a01, &a02, &a01 == &a02)
	fmt.Printf("&s1.F2:%#p &s1.F3:%#p (&s1.F2==&s1.F3):%v\n", &s1.F2, &s1.F3, &s1.F2 == &s1.F3)
	fmt.Printf("&s02:%#p &s1.F2:%#p (&s02==&s1.F2):%v\n", &s02, &s1.F2, &s02 == &s1.F2)
}
