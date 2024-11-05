// Package example 演示如何使用 Example 函数
package example

// Add 返回 a+b 的结果.
func Add(a, b int) int {
	return a + b
}

// Adder 是一个包含 Add(a, b int) int 方法的接口.
type Adder interface {
	Add(a, b int) int
}

// AdderFunc 接口可以用来把一个 func(a, b int) int 转换为 Adder.
type AdderFunc func(a, b int) int

// Add 实现了 Adder 接口的 Add 方法.
func (f AdderFunc) Add(a, b int) int {
	return f(a, b)
}
