package main

import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	// 时间点1 (赋值完成的时间)
	c <- 0
	// 时间点2 (发送完成的时间)
}

func main() {
	go f()
	<-c
	// 时间点3 (接收完成的时间)
	fmt.Println(a)
}
