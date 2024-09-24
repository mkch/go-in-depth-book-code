package main

type S1 struct {
	D int
	C byte
}

func (s *S1) B() {}

type S2 struct {
	E string
}

func (s S2) D() {}

type S3 struct {
	A int
	S1
}

type I interface {
	C()
	E()
}

type S struct {
	S2
	S3
	I
	A string
}

func main() {
	var s S

	// s.A 访问的是第 30 行定义的 A
	s.A = ""

	// s.B 访问的是第 8 行定义的 B
	s.B()

	// s.C 访问的是第 22 行定义的 C
	s.C()

	// s.D 访问的是第 14 行定义的 D
	s.D()

	// s.E 有歧义
	//s.E = "" // !! 编译错误 !!
	// 只能通过以下形式分别访问两个 E
	s.I.E()
	s.S2.E = ""

	// D() 可以通过 S 和 *S 访问
	func() S { return s }().D()

	// B() 只能通过 *S 访问, 而临时变量无法取指针
	//func() S { return s }().B()
}
