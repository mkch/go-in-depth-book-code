package main

import (
	"bytes"
	"io"
)

func Instantiated() {
	type S[T any] struct{ A T }
	type S2 struct{ A int }

	// S[int] 是定义类型, S2 也是定义类型
	// 两个定义类型是不同的
	// var _ S[int] = S2{} // !! 语法错误 !!

	// S[int] 的底层类型为 struct{A int}
	// 可以赋值
	var _ S[int] = struct{ A int }{}
}

func TypeParamNamed[T1, T2 any](a1 T1, a2 T2) {
	// T1 和 T2 为两个不同的命名类型
	// 因此 a1 和 a1 不可相互赋值
	// a1 = a2 // !! 语法错误 !!
}

func TypeParamUnderlying[T any]() {
	// T2 的底层类型为 any
	// any 类型的值可以赋值给 T2
	type T2 any
	var a T2 = any(1)
	// 虽然 T 的底层类型为 any
	// 但是 any 类型的值不能赋值给 T
	//var b T = any(1) // !! 语法错误 !!
	// nil 也不能赋值给 T
	//var b T = nil // !! 语法错误 !!
	_ = a
}

func TypeParamUnderlying2[T io.Reader](r T) {
	var reader io.Reader
	//r = reader // !! 语法错误 !!
	r = reader.(T)
}

func UseTypeParamUnderlying2() {
	var buf bytes.Buffer
	TypeParamUnderlying2(&buf)

	var Instantiated = func(r *bytes.Buffer) {
		var reader io.Reader
		//r = reader // !! 语法错误 !!
		r = reader.(*bytes.Buffer)
	}
	Instantiated(&buf)
}

type Slice []int
type Slice2 []int

func Commonality[TP *int | []int, TS Slice | Slice2]() {
	var a TP = nil
	var a2 TS = []int{1, 2, 3}
	var s []int = a2
	_, _, _ = a, a2, s
}
