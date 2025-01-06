package main

import "fmt"

type T struct{}

func (tv T) Mv(a int) int          { return a + 1 } // 接收者为 T
func (tp *T) Mp(f float32) float32 { return 1.0 }   // 接收者为 *T

var t T

func main() {
	T.Mv(t, 1)       // 相当于 t.Mv(1)
	(*T).Mp(&t, 1.0) // 相当于 (&t).Mp(1.0)
	(*T).Mv(&t, 1)   // 相当于 (&t).Mv(1)
	// T.Mp(&t, 1.0)

	f2()
}

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
}

type X struct {
	v int
}

func (x *X) Print() {
	fmt.Println(x.v)
}

func f2() {
	var x = &X{v: 1}
	f := x.Print // 接收者 x 求值并保存在 f 内
	f()          // 输出 1
	x.v = 2      // 修改 x.v 会影响 f
	f()          // 输出 2
	x = &X{v: 3} // 修改 x 不影响保存在 f 内的接收者
	f()          // 输出 2
}
