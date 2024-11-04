package main

import (
	"fmt"
	"time"
)

func main() {
	//cap0()
	cap1()
}

// cap0 演示无缓冲管道并发写操作.
func cap0() {
	var ch = make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()

	fmt.Println(<-ch, <-ch)
}

// cap0 演示有缓冲管道并发写操作.
func cap1() {
	var ch = make(chan int, 1)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()

	go func() {
		fmt.Println(<-ch)
	}()

	go func() {
		fmt.Println(<-ch)
	}()

	go func() {
		fmt.Println(<-ch)
	}()

	time.Sleep(time.Second * 1)
}
