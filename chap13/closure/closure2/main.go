package main

import "fmt"

func main() {
	// acc3 和 acc4 将共享内部状态 sub
	acc3, acc4 := NewAccumulator2()
	fmt.Println(acc3(1), acc3(2), acc3(3))
	fmt.Println(acc4(10), acc4(20), acc4(30))
}

// Accumulator 时一个累加器
// 每次调用 Accumulator 都会将参数 n 累加到一个内部状态
// 并返回迄今为止的累加值
type Accumulator func(n int) (acc int)

// NewAccumulator2 返回两个新的累加器
func NewAccumulator2() (Accumulator, Accumulator) {
	var sum int
	return func(n int) (acc int) {
			sum += n
			return sum
		},
		func(n int) (acc int) {
			sum += n
			return sum
		}
}
