package main

import "fmt"

func main() {
	// 生成两个累加器
	var acc1 = NewAccumulator()
	var acc2 = NewAccumulator()
	// 两个累加器拥有各自独立的内部状态 sum
	fmt.Println(acc1(1), acc1(2), acc1(3))
	fmt.Println(acc2(10), acc2(20), acc2(30))

	// acc3 和 acc4 将共享内部状态 sub
	acc3, acc4 := NewAccumulator2()
	fmt.Println(acc3(1), acc3(2), acc3(3))
	fmt.Println(acc4(10), acc4(20), acc4(30))

}

// Accumulator 时一个累加器
// 每次调用 Accumulator 都会将参数 n 累加到一个内部状态
// 并返回迄今为止的累加值
type Accumulator func(n int) (acc int)

// NewAccumulator() 返回一个新的累加器
func NewAccumulator() Accumulator {
	var sum int                    // 第 21 行
	return func(n int) (acc int) { // 第 22 行
		sum += n // 第 23 行
		return sum
	}
}

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
