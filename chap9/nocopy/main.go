package main

import "fmt"

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type S struct {
	noCopy noCopy // S 不应该被拷贝
}

func main() {
	var s S
	// go vet 会对此行输出做出反应
	fmt.Println(s)
}
