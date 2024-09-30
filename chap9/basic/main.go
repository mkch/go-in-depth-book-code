package main

import (
	"fmt"
)

type S struct {
	F1 string
	F2 int16
}

func main() {
	var a S // 定义一个类型为 S 的变量 a
	// 使用 s.F 的形式来访问结构体的字段
	fmt.Println(a.F1, a.F2) // 结构体的成员会被自动初始化为零值
	// 对结构体成员赋值
	a.F1 = "abc"
	a.F2 = 100
}

// 演示有效的字段类型
type T struct {
	F1 []T
	F2 *T
	F3 [10][]T
	F4 func() T
	F6 struct{ F *T }
}
