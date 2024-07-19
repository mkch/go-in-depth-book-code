package main

type S1 struct {
	D int
	C byte
}

func (s *S1) B() {} // 2

type S2 struct {
	E string
}

func (s S2) D() {} // 4

type S3 struct {
	A int
	S1
}

type I interface {
	C() // 3
	E()
}

type S struct {
	S2
	S3
	I
	A string // 1
}

func main() {
	var s S

	// s.A 访问的是 1 处定义的 A
	s.A = ""

	// s.B 访问的是 2 处定义的 B
	s.B()

	// s.C 访问的是 3 处定义的 C
	s.C()

	// s.D 访问的是 4 处定义的 D
	s.D()

	// 编译错误!
	// s.E 有歧义
	//s.E = ""
	// 只能通过以下形式分别访问两个 E
	s.I.E()
	s.S2.E = ""

	// D() 可以通过 S 和 *S 访问
	func() S { return s }().D()

	// 5 编译错误！
	// B() 只能通过 *S 访问, 而临时变量无法取指针
	//func() S { return s }().B()
}
