package main

import (
	"fmt"
	"time"
)

func main() {
	go add(1, 2) // 1: go 函数调用(参数...)
	var s1 S
	go s1.Method1() // 2: go 方法调用()
	go func(n int) {
		fmt.Printf("func(%v)\n", n)
	}(100) // 3: go 匿名函数调用()

	time.Sleep(3 * time.Second) // 强制等待 3 秒钟. 以后会有更好的实现
}

func add(a, b int) int {
	fmt.Printf("f(%v, %v)\n", a, b)
	return a + b
}

type S struct{}

func (s S) Method1() {
	fmt.Println("S.Method1()")
}
