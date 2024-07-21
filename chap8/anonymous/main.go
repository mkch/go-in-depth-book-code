package main

import "fmt"

func main() {
	// 匿名函数是一个表达式
	// 可以保存到变量
	print := func() {
		fmt.Println("anonymous func1")
	}
	// 变量 print 的类型为 func()
	// 可以像函数一样调用
	print()

	// 匿名函数作为参数传递
	call(func() {
		fmt.Println("anonymous func2")
	})
	call(print)
	call(gen())

	// 直接调用匿名函数
	func() {
		fmt.Println("anonymous func3")
	}()
}

func gen() func() {
	// 匿名函数作为返回值
	return func() {
		fmt.Println("anonymous func4")
	}
}

func call(f func()) {
	f() // 调用函数变量
}
