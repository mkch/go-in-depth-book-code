package main

import "fmt"

func main() {
	var result int64 = Add(1, 2)
	_ = result

	fmt.Printf("some string\n")                 // 不提供 a
	fmt.Printf("value is %v\n", 1)              // 为 a 提供 1个参数
	fmt.Printf("value1:%v value2:%v\n", 2, "a") // 为 a 提供 2 个参数

	fmt.Println(avg(), avg(1), avg(5, 10, 15))

	var n int
	n = avg(1, 3, 5) // n 为 3
	fmt.Println(n)
	num := []int{1, 3, 5}
	// num... 把 num 展开, 等效于 avg(1,3,5)
	n = avg(num...) // n 为 3
	fmt.Println(n)
	n = avg() // n 为 0
	fmt.Println(n)
}

// Add 把 a+b 的结果作为 int64 返回
func Add(a int, b int) int64 {
	return int64(a + b)
}

// avg 返回所有参数的平均值
func avg(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	var sum int
	for _, i := range args {
		sum += i
	}
	return sum / len(args)
}
