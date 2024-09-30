package main

import (
	"fmt"
	"math"
)

func main() {
	basic()
}

func basic() {
	var c1 chan int        // 可以传递 int 的 channel
	var c2 chan string     // 可以传递 string 的 channel
	var c3 chan []int      // 可以传递 []int 的 channel
	var c4 chan chan []int // 相当于 chan (chan []int), 可以传递 (chan []int) 的 channel

	c1 = make(chan int)       // 创建一个容量为 0 的 channel. 无缓冲
	c2 = make(chan string, 2) // 创建一个容量为 2 的channel

	c2 <- "a"    // 不阻塞
	l := len(c2) // len(c2) 取出 c2 缓冲区的占用量, 此处为 1
	c2 <- "b"    // 不阻塞
	l = len(c2)  // 2
	<-c2         // 不阻塞, 将读出 "a"
	l = len(c2)  // 1
	<-c2         //  不阻塞, 将读出 "b"
	l = len(c2)  // 0
	c2 <- "c"    // 不阻塞
	l = len(c2)  // 1
	c2 <- "d"    // 不阻塞
	l = len(c2)  // 2
	//c2 <- "e"    // 阻塞!!

	go func() {
		// 1 启动一个新的 goroutine, 向 c1 中写入 1
		c1 <- 1
	}()
	<-c1 // 2 此处可以读出 1.

	//c1 <- 100             // 向 c1 中写入值 100
	//var str string = <-c2 // 从 c2 中读出一个值并赋值给 str
	_, _, _, _, _ = l, c1, c2, c3, c4

	var cr <-chan int // 只读 channel
	cr = c1           // c1r 是 c1 的 "读端"
	var cw <-chan int // 只写 channel
	cw = c1           // c1w 是 c1 的 "写端"

	c := make(chan<- int, 1)

	_, _, _ = cr, cw, c

	gen := NewIdGenerator()
	fmt.Println(gen())
	fmt.Println(gen())

	closed()
	closeSignal()
	closeSignal2()
}

// NewIdGenerator 返回一个并发安全的顺序 ID 生成器
func NewIdGenerator() func() (g uint64) {
	c := make(chan uint64)
	producer := func(w chan<- uint64) {
		for i := range uint64(math.MaxUint64) {
			w <- i
		}
		panic("out of Id")
	}
	go producer(c)
	return func() uint64 { return <-c }
}

// closed 演示 channel 的 close() 操作
func closed() {
	c := make(chan int)
	go func() {
		c <- 0
		c <- 1
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}

// closeSignal 演示用 close() 作为操作已完成的信号
func closeSignal() {
	c := make(chan int)
	// 生产者
	go func() {
		c <- 1
		c <- 2
		close(c) // 关闭 channel, 通知消费者退出
	}()

	// 消费者
	for {
		v, ok := <-c
		if !ok { // 1: 是否需要退出
			break
		}
		// 使用 v
		fmt.Println(v)
	}
}

func closeSignal2() {
	c := make(chan int)
	// 生产者
	go func() {
		c <- 1
		c <- 2
		close(c) // 关闭 channel, 通知消费者退出
	}()

	// 消费者
	for v := range c {
		// 使用 v
		fmt.Println(v)
	}
}
