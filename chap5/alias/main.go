package main

// Integer 为 int 的别名.
// 所有出现 Integer 的地方都相当于 int.
type Integer = int

// 非法: 无法为 int 添加方法.
// 相当于:
// func (i int) F() {}
//func (i Integer) F() {}

func main() {
	var i Integer // 等价于 var i int
	var i2 int
	i += i2 // i 和 i2 的类型都为 int
}
