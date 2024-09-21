package main

import (
	"fmt"
	"os"
)

func main() {
	concept()
	methodvalue()
	closure()
}

func concept() {
	f, _ := os.Open("main.go")
	// 1 fc 保存了 f.Close 的方法值
	var fc func() error = f.Close // 注意此处不是调用 f.Close()
	// 2 相当于调用 f.Close()
	fc()
}

func methodvalue() {
	f, _ := os.Open("main.go")
	defer f.Close() // 1
	// 相当于
	// fc := f.Close
	// defer fc()

	f = nil // 2
}

func closure() {
	f, _ := os.Open("main.go")
	defer func() {
		fmt.Println(f)
		fmt.Println(f.Close())
	}() // 1

	f = nil // 2
}
