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

type Bool bool

func (b Bool) F() {
}

type Channel interface {
	chan uint | chan int
}

func F[C Channel](c C) {
	// 非法: 只有核心类型为 chan 的变量才能使用 <- 操作符
	// Channel 没有核心类型, 因为 chan uint 和 chan int 的元素类型不相同
	//c <- 0
}
