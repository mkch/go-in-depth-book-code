package main

import "fmt"

var c = make(chan string)

func f() {
	c <- "hello world"
}

func main() {
	go f()
	fmt.Println(<-c)
}
