package main

import "fmt"

var A = 0

const B = len("abc")

type C struct{}

func D() {}

func E[T any](a T) {
	var b T = a
	fmt.Println(b)
}

func F(a int) {
	var b = max(a, 1)
	fmt.Println(b)
	if b == 0 {
		var max = float32(a) + 9.9
		fmt.Println(max)
		{
			max := func() {}
			max()
		}
		fmt.Println(max + 1)
	}
}

func Label() func() {
	fmt.Println("start")
	var n = 9
loop:
	for i := range n {
		if i > 10 {
			switch i {
			case 11:
				break loop // 跳出 for 循环而不是 switch
			}
		}
	}
	return func() {
		// goto loop // !! 语法错误
	}
}
