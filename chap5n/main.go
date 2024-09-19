package main

// Integer 为 int 的别名
// 所有出现 Integer 的地方都相当于 int
type Integer = int

// 非法: 无法为预定义类型 int 添加方法
// 相当于
// func (i int) F() {}
//func (i Integer) F() {}

func main() {
	var i Integer // 等价于 var i int
	var i2 int
	i += i2 // i 和 i2 的类型都为 int
}

func Assignable2() {
	type S struct{ A int }
	var named S
	var composite struct{ A int }
	named = composite
	composite = named
}

func Assignable3() {
	var rw chan int
	var r <-chan int
	r = rw
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
	var _ complex128 = 1 + 2i
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
