package main

import (
	"fmt"
)

func main() {
	elemAndIndex()
	arrayType()
	multiD()

	var MZHeader = [...]byte{'M', 'Z'} // [...] 编译器自动推断数组的长度
	_ = &MZHeader

	arrayValue()
	arrayPtr()
	slice1()
	makeSlice()
	copySlice()
	appendSlice()
	slice2D()
}

// 演示数组元素和索引.
func elemAndIndex() {
	// ary 中有3个元素, 类型均为 int, 值均为 int 的零值(0)
	var ary [3]int
	// 可以用 len() 内建函数来取得数组的长度
	var length = len(ary) // length 为 3
	// 读取 ary 的各个元素
	e0, e1, e2 := ary[0], ary[1], ary[2]
	// 为 ary 的各个元素赋值
	ary[0] = 10
	ary[1] = 11
	ary[2] = 12
	_, _, _, _ = length, e0, e1, e2
}

// 演示数组长度是类型的一部分.
func arrayType() {
	// var ary1 [1]int
	// var ary2 [2]int
	// ary2 = ary1  // 编译错误: 类型不匹配
	// e1 = ary1[1] // 编译错误: 索引越界
}

// 演示多维数组.
func multiD() {
	var ary [2][3]int
	ary[0] = [3]int{1, 2, 3}
	ary[1] = [3]int{4, 5, 6}
}

// 演示数组是值类型.
func arrayValue() {
	ary1 := [...]byte{1, 2, 3}
	ary2 := ary1
	fmt.Printf("&ary1=%p &ary1[0]=%p\n", &ary1, &ary1[0])
	fmt.Printf("&ary2=%p &ary2[0]=%p\n", &ary2, &ary2[0])
}

// 演示数组指针.
func arrayPtr() {
	ary1 := [...]byte{1, 2, 3}
	var pary1 *[3]byte = &ary1 // pary1 为 ary1 的指针
	// 可以通过数组指针直接索引其成员
	// 等效于 (*pary1)[0] = 10
	pary1[0] = 10 // 通过指针修改数组成员
	// ary1 的内容被改变了
	fmt.Println(ary1)
}

// 演示切片操作.
func slice1() {
	var ary = [5]int{10, 20, 30, 40, 50}
	var s []int = ary[1:3] // 1
	fmt.Println(s, len(s), cap(s), cap(ary))
	fmt.Println(s[0], s[1])

	var s1 = s[1:2:3] // 2
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s1[0])

	s1[0] = 999 // 3
	fmt.Println(ary, s, s1)
}

// 演示 make().
func makeSlice() {
	// s 的长度为 3, 容量为 3
	var s []int = make([]int, 3)
	// s2 的类型为[]byte
	// 长度为 3, 容量为 6.
	s2 := make([]byte, 3, 6)
	fmt.Println(len(s), cap(s))
	fmt.Println(len(s2), cap(s2))
	copy(s, s)
}

// 演示 copy().
func copySlice() {
	s := []byte{1, 2, 3}
	s1 := make([]byte, len(s))
	n := copy(s1, s) // n 等于 3
	// 至此, s1 是 s 的一个"深拷贝".

	s2 := make([]byte, 5)
	n2 := copy(s2, "ab") // n2 等于 2
	// 至此, s2 的内容为 {'a','b', 0, 0, 0}

	fmt.Println(s1, n, s2, n2)
}

// 演示 append().
func appendSlice() {
	s := []byte{1}
	s = append(s, 2, 3)
	// cap(s) 可能大于 3
	fmt.Println(len(s), cap(s))
}

func slice2D() {
	var s = make([][]int, 2)
	s[0] = []int{1, 2, 3} // 各维均需要独立初始化
	s[1] = []int{4, 5, 6}

	e1 := s[0][0] // 1
	e2 := s[0][1] // 2
	e3 := s[0][2] // 3
	e4 := s[1][0] // 4
	// ...
	fmt.Println(e1, e2, e3, e4)
}
