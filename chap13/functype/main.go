package main

import "fmt"

// PrintFunc 是一个函数类型
type PrintFunc func(...any) (int, error)

func main() {
	var f PrintFunc   // f 为一个函数类型的变量
	f = gen()         // 可以把函数值保存在变量中
	_, _ = f("test1") // 函数类型变量可以调用
	call(f)           // 函数类型变量可以作为参数传递
	call(fmt.Println) // 或者直接把函数作为参数传递

}

func call(f PrintFunc) { // 可以把函数作为参数传递
	f("test")
}

func gen() PrintFunc {
	return fmt.Println // 可以返回一个函数
}
