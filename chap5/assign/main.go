package main

import (
	"bytes"
	"io"
	"os"
)

func main() {

}

func Assignable2() {
	type S struct{ A int }
	var named S
	var composite struct{ A int }
	named = composite
	composite = named

	type A []int
	var named2 A
	var composite2 = []int{1, 2}
	named2 = composite2
	composite2 = named2
}

func Assignable3() {
	var rw chan int
	var r <-chan int
	r = rw
	_ = r

	type RWC chan int
	type RC <-chan int
	var rc RC
	var rwc RWC

	rc = r
	rc = rw
	rwc = rw

	// 两个命名类型是不同的
	// rc = rwc // !! 语法错误 !!

	_, _ = rc, rwc

}

func Assignable4() {
	// *bytes.Buffer 和 *os.File 都实现了 io.Reader 接口
	var r io.Reader = &bytes.Buffer{}
	r = (*os.File)(nil)
	_ = r
}

func Assignable5() {
	var _ *int = nil
	var _ func() = nil
	var _ []int = nil
	var _ map[int]int = nil
	var _ chan int = nil
	var _ error = nil
}

func Assignable6() {
	const c = 1
	var _ int = c
	var _ byte = c
	var _ float32 = c
	type Int int
	var _ Int = c
	var _ complex128 = c
}

func Assignable7[T []int | *int](a T) {
	_ = a
	a = nil // []int 和 *int 都可以表示为 nil
}

type S []int
type S2 []int

func Assignable8[T S | S2](a T) {
	_ = a
	a = []int{1, 2, 3} // S 和 S1 的底层类型都为 []int
	_ = a
}

func Assignable9[V S | S2](x V) {
	var s []int = x // S1 和 S2 都可赋值给 []int
	_ = s
}

func Convertable2() {
	type (
		A struct{ A int }
		B A
		C struct {
			A int `example:"tag"`
		}
	)
	// A, B, C 的底层类型一致(不考虑标签), 可以相互转换
	var a A = A(B{})
	var b B = B(a)
	var c C = C(b)

	_, _, _ = a, b, c

}

func Convertable3() {
	type (
		A struct{ A int }
		B A
		C struct {
			A int `example:"tag"`
		}
	)
	// A, B, C 的底层类型一致(不考虑标签)
	// *A, *B, *C 可以相互转换
	var a *A = (*A)(&B{})
	var b *B = (*B)(a)
	var c *C = (*C)(b)

	_, _, _ = a, b, c
}

func Convertable4() {
	var a int = 1
	var b uint = uint(a)
	var c float32 = float32(a)
	a = int(c)

	var d complex64 = 1.0 + 2i
	var e complex128 = complex128(d)
	d = complex64(e)

	_, _, _ = a, b, c
}

func Convertable5() {
	var a string = string(97)          // "a"
	var b string = string('a')         // "a"
	var c string = string([]byte{97})  // "a"
	var d string = string([]rune{'a'}) // "a"
	_, _, _, _ = a, b, c, d
}

func Convertable6() {
	var a []byte = []byte("a") // [97]
	var b []rune = []rune("a") // [97]

	_, _ = a, b
}

func Convertable7() {
	var a = []byte{1, 2, 3}
	// 如果数组的长度大于切片长度, 会引发运行时 panic
	var b [3]byte = ([3]byte)(a)
	var c *[3]byte = (*[3]byte)(a)

	_, _, _ = a, b, c
}

type Bool bool

func (b Bool) F() {
}

type Channel interface {
	chan uint | chan int
}

func F[C Channel](c C) {
	// 非法: 只有核心类型为 chan 的变量才能使用 <- 操作符
	// Channel 没有核心类型, 因为 chan uint 和 chan int 的元素类型不相同
	//c <- 0
}
