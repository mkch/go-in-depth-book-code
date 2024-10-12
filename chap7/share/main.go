package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str1 = "abcd"
	var str2 = str1[:2]
	fmt.Println(str2)
	fmt.Println(unsafe.StringData(str1))
	fmt.Println(unsafe.StringData(str2))
}
