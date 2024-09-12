package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int         // n 类型为 int, 初始值为 0
	var n1 uint64 = 1 // n2 类型为 uint64, 初始值为 1
	var n2 = "abc"    // n3 类型为 string, 初始值为 "abc"
	// 也可以一次性声明多个变量
	var (
		b byte = 0xFF
		c      = '中'
	)

	_, _, _, _, _ = n, n1, n2, b, c

	fmt.Println(State1, State2, State3, State4, State5, x, State6, y)

	var ary [3]int
	// s 和 s2 共享底层数组 ary
	s := ary[0:2]
	s2 := s[0:1]
	// 此 append 同时影响 s2, s 和 ary 的值
	s2 = append(s2, 2)
	fmt.Println(s2, s, ary)

	var m map[int]string
	l := len(m) // 0 (nil map 的长度为 0)
	_ = l
	day := m[2] // "" (nil map 相当于只读空 map)
	_ = day
	m = make(map[int]string)
	m[1] = "Monday"
	m[2] = "Tuesday"
	l = len(m) // 2
	_ = l
	day, ok := m[2] // "Tuesday", true
	_, _ = day, ok
	wrong, ok := m[100] // "", false
	_, _ = wrong, ok

	var ch chan string
	// 读或写 nil chan 将永远阻塞
	// ch <- "" // 永远阻塞
	// <-ch     // 永远阻塞
	ch = make(chan string)
	go func() { // goroutine 1
		msg := <-ch
		fmt.Println(msg)
	}()

	go func() { // goroutine 2
		ch <- "abc"
	}()

	var ready = make(chan struct{})
	go func() {
		<-ready // 等待准备完毕
		fmt.Println("Go!")
	}()
	go func() {
		<-ready // 等待准备完毕
		fmt.Println("Go 2!")
	}()

	// 进行准备工作
	// ...
	close(ready) // 发出 ready 信号

	If()
	For()
	Switch()
	LabelFor()
}

func If() {
	if n, err := strconv.Atoi("1"); err != nil {
		// 处理错误
	} else {
		_ = n // 使用 n
	}

	n, err := strconv.Atoi("1")
	if err != nil {
		// 处理错误
	} else {
		_ = n // 使用 n
	}
}

func For() {
	var cond bool
	if cond {
		// 循环体
	}

	for i := 0; i < 10; i++ {
		// 循环体
		fmt.Println(i)
	}
	for i := range 10 {
		fmt.Println(i)
	}
}

func F() string { return "" }

func Switch() {
	str := F()
	switch str {
	case "a":
		// 处理 str == "a" 的情况
	case "b":
		// 处理 str == "b" 的情况
		fallthrough
	case "c":
		// 处理 str == "b" 或者 处理 str == "c" 的情况
	default:
		// 处理其他情况
	}

	switch str := F(); str {
	// ...
	}

	var v any = F()
	switch v.(type) {
	case int:
		// 处理 v 为 int 的情况
	case string:
		// 处理 v 为 string 的情况
	}

	switch {
	case v == 1:
		// ...
	case str == "":
		// ...
	}
	// 等价于
	if v == 1 {
		// ...
	} else if str == "" {
		// ...
	}
}

func LabelFor() {
Loop: // 为 for 循环定义一个标签
	for i := range 10 {
		switch i {
		case 1, 3, 5:
			// 处理 1,3,5
		case 2, 4:
			break Loop // 跳出 for 循环
		default:
			// 处理其他情况
		}
	}
}

func Select() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	select {
	case v1 := <-ch1:
		_ = v1 // 使用 v1
	case v2, ok := <-ch2:
		if ok {
			_ = v2 // 使用
		}
	case ch1 <- 1:
		// 在成功向 ch1 发送 1 后执行
	default:
		// 当以上所有分支均会阻塞时执行
	}
}

type State uint

const (
	State1    State = iota                       // 0 (iota = 0)
	State2                                       // 1 (iota = 1)
	State3    State = 255                        // 255 (iota = 2, 未用)
	State4          = iota + iota                // 6 (iota = 3)
	State5, x       = 1 << iota, 1 << (iota + 1) // 16,32 (iota = 4)
	_, _                                         // (iota = 5, 未用)
	State6, y                                    // 64, 128 (iota = 6)
)
