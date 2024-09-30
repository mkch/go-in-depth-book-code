package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a S
	fmt.Println(unsafe.Sizeof(a), unsafe.Offsetof(a.F2))
}

type S struct {
	F1 byte
	_  [2]byte // padding
	F2 byte
	F3 uint32
}
