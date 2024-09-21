package main

import (
	"fmt"
	"os"
)

// PrintSum 打印出从 from 到 to 的整数的和
// 本函数在计算出和 sum 后, 会调用 print(sum) 来执行打印操作
func PrintSum(from, to int, print func(n int)) {
	sum := (from + to) * (to - from + 1) / 2
	print(sum)
}

func main() {
	// 打印到控制台
	PrintSum(1, 100, func(n int) { fmt.Println(n) })
	// 打印到文件
	PrintSum(1, 100, func(n int) {
		f, err := os.OpenFile("output.txt", os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Fprintf(f, "The sum from 1 to 100 is %v\n", n)
	})
}
