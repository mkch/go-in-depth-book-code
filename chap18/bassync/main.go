package main

import "fmt"

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	fmt.Println(b)
	fmt.Println(a)
}

func main() {
	go f()
	g()
}
