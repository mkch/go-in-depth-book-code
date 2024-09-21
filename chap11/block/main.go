package main

import "fmt"

func main() {
	var ch = make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()

	fmt.Println(<-ch, <-ch)
}
