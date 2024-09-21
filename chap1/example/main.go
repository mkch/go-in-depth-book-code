package main

import (
	"fmt"

	"example.com/demo/greeting"
)

func main() {
	fmt.Println(greeting.Message())
	var ch chan int
	fmt.Println(len(ch), cap(ch))
	close(ch)
}
