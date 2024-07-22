package main

import (
	"fmt"
)

func f(n int) {
	defer fmt.Println("defer1") // 1
	if n == 1 {
		return // 3
	} else if n == 0 {
		panic(0) // 4
	}
	defer fmt.Println("defer2") // 2
}

func main() {
	fmt.Println("f(1)")
	f(1)
	fmt.Println("f(2)")
	f(2)
	fmt.Println("f(0)")
	f(0)
}
