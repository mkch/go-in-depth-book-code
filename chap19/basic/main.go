package main

import (
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"slices"
)

// Int 为有符号整数类型.
type Int interface {
	~int8 | ~int | ~int16 | ~int32 | ~int64
}

// Sum 计算参数之和
func Sum[T Int](values ...T) (r T) {
	for _, v := range values {
		r += v
	}
	return
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// List 是一个链表
type List[T any] struct {
	Head *Node[T]
}

// Size 返回 l 的长度
func (l *List[T]) Size() (r int) {
	for p := l.Head; p != nil; p = p.Next {
		r++
	}
	return
}

type IntStringer interface {
	String() string
	Int
}

func SumStr[T IntStringer](values ...T) string {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum.String()
}

// Scale 返回 s 的一个拷贝, 其中每个元素都乘以 c
func Scale[S ~[]E, E Int](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

// Point 为一个多维的点
type Point []int32

func (p Point) String() string {
	return fmt.Sprintf("Point %v", ([]int32)(p))
}

func main() {
	var v = []int32{1, 2, 3}
	v2 := Scale(v, 2)
	fmt.Println(v2)

	var pt = Point{1, 2, 3}
	pt2 := Scale(pt, 2)
	fmt.Println(pt2.String())
}

func UseSlicesSort() {
	var intSlice = []int{3, 4, 1}
	slices.Sort[[]int, int](intSlice)
	slices.Sort(intSlice)
}

func Copy(dest http.Header, src url.Values) {
	maps.Copy(dest, src)
}
