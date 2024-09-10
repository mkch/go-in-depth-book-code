package main

import (
	"fmt"
	"iter"
	"slices"
)

// Values 返回整数序列 1, 2, 3 的迭代器.
func Values123() iter.Seq[int] {
	return slices.Values([]int{1, 2, 3})
}

// Values123Words 返回序列 (1, "One"), (2, "two"), (3, "three") 的迭代器.
func Values123Words() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		if !yield(1, "one") {
			return
		}
		if !yield(2, "two") {
			return
		}
		if !yield(3, "three") {
			return
		}
	}
}

func Values123V2() iter.Seq[int] {
	return func(yield func(int) bool) {
		if !yield(1) {
			return
		}
		if !yield(2) {
			return
		}
		if !yield(3) {
			return
		}
	}
}

// Zip seq1 和 seq2 "缝合" 为一个 iter.Seq2
func Zip[K, V any](seq1 iter.Seq[K], seq2 iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			k, ok1 := next1()
			if !ok1 {
				return
			}
			v, ok2 := next2()
			if !ok2 {
				return
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func ZipNaive[K, V any](seq1 iter.Seq[K], seq2 iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range seq1 {
			var v V
			// 无法实现按需取出 seq2 的值
			// v = ...
			yield(k, v)
		}
	}
}

func main() {
	seq := Values123()
	for n := range seq {
		fmt.Println(n)
	}

	// 内联展开:
	seq = func(yield func(int) bool) {
		for _, v := range []int{1, 2, 3} {
			if !yield(v) {
				return
			}
		}
	}
	seq(func(v int) bool {
		fmt.Println(v)
		return true
	})

	// 去掉seq变量
	func(yield func(int) bool) {
		for _, v := range []int{1, 2, 3} {
			if !yield(v) {
				return
			}
		}
	}(func(v int) bool {
		fmt.Println(v)
		return true
	})

	// 去掉匿名函数
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
		var ok = true
		if !ok {
			return
		}
	}

	// 去掉 ok 变量
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}

	for n := range Values123V2() {
		fmt.Println(n)
	}
	// 相当于
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)

	for n, w := range Values123Words() {
		fmt.Println(n, w)
	}

	keys := slices.Values([]int{1, 2, 3})
	values := slices.Values([]string{"a", "b", "c"})
	for k, v := range Zip(keys, values) {
		fmt.Println(k, v)
	}
}
