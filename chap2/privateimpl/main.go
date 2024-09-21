package main

import (
	"fmt"

	"example.com/privateimpl/pkg"
)

type MyIface struct {
	pkg.Impl // 为了实现 pkg.Iface
}

func (i MyIface) F() {
	fmt.Println("MyFace.F")
}

func main() {
	pkg.Use(MyIface{})
}
