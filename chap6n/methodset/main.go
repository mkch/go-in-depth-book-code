package main

import "fmt"

type I interface {
	I()
}

type A int

func (A) I() {}

func (*A) P() {}

type B int

func (*B) I() {}
func (*B) P() {}

func f() {
	var a A
	// A 的方法集中包含 I()
	var _ I = a
	// *A 的方法集中也包含 I()
	var _ I = &a

	var b B
	// *B 的方法集中包含 I()
	var _ I = &b
	// B 的方法集为空
	// var _ I = b // !! 语法错误 !!

	a.P() // 自动寻址, 相当于 (&a).P()
	b.I() // 自动寻址, 相当于 (&b).I()

	f := func() A { return A(1) }
	// 临时变量不能寻址
	// f().P() // !! 语法错误 !!
	_ = f
}

type Point struct {
	x, y float64
}

func (p *Point) SetX(a float64) {
	p.x = a
}

func (p *Point) SetY(a float64) {
	p.y = a
}

type X struct {
	v int
}

func (x *X) Print() {
	fmt.Println(x.v)
}

func main() {
	var x = &X{v: 1}
	f := x.Print // 接收者 *X 计算并保存在 f 内
	f()          // 输出 1
	x = &X{v: 2} // 修改 x 不影响保存在 f 内的接收者
	f()          // 输出1
}
