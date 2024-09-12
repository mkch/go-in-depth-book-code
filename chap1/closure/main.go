package main

import "fmt"

func F() (get func() int, set func(int)) {
	var n int
	// get 和 set 通过闭包共享变量 n
	get = func() int { return n }
	set = func(i int) { n = i }
	return
}

func MyPrint(a ...any) (n int, err error) {
	// 实现代码省略
	return
}

func Call(f func(...any) (int, error), args ...any) {
	f(args...)
}

func main() {
	// 函数赋值给变量
	f := fmt.Println // f 的类型为 func(...any) (int, error)
	// 函数类型的变量可以直接调用
	f(1, 2, 3)
	f = MyPrint
	// 函数类型的变量可作为参数传递
	Call(f)
	// 函数自身也可作为参数传递
	Call(fmt.Print, 1, 2, 3)
	Call(MyPrint, 1, 2, 3)

	get, set := F()
	set(1)
	fmt.Println(get())
	set(2)
	fmt.Println(get())
}
