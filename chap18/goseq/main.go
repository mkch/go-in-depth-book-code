package main

import "fmt"

var a string

func f() {
	fmt.Println(a)
}

func hello() {
	a = "hello, world"
	go f()
}
