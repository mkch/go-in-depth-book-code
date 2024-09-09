package main

import "fmt"

func main() {
	var ary = [...]string{"abc", "def"}

	// 将输出
	// 0 abc
	// 1 def
	for i, str := range ary {
		fmt.Println(i, str)
	}

	// 将输出
	// 0
	// 1
	for i := range ary {
		fmt.Println(i)
	}

	var s = []int{1, 2, 3}
	// 将输出
	// 0 1
	// 1 2
	// 2 3
	for i, n := range s {
		fmt.Println(i, n)
	}
}
