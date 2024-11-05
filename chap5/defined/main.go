package main

type MyInt int
type MyInt2 MyInt

func Test() {
	var n int = 1
	var mn MyInt = 2
	var mn2 MyInt2 = 3
	// int MyInt MyInt2 的底层类型都为 int
	n = int(mn)
	n = int(mn2)
	mn = MyInt(mn2)
	mn = MyInt(n)

	var n1 MyInt = 1
	var _ MyInt = n1.Add(2)
}

func (n MyInt) Add(n2 MyInt) MyInt {
	return n + n2
}

func Underlying() {
	type Int int
	type A bool
	type B A
	type C [3]int
	type D [3]int
	type E1 map[int]string
	type E E1
	type F struct{ A, B int }
	type G func(int) int
}

func TestStruct() {
	var s1 struct {
		A, B int
	}

	var s2 struct {
		B, A int
	}

	var s3 struct {
		B int `example:"tag"`
		A int
	}

	// s1, s2, s3 的类型不相同
	// s1 = s2 // !! 语法错误 !!
	// s2 = s3 // !! 语法错误 !!
	// s1 = s3 // !! 语法错误 !!

	// 类型转换时忽略标签
	s2 = struct{ B, A int }(s3)

	_, _, _ = s1, s2, s3
}
