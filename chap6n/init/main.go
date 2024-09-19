package main

import "fmt"

var x = c

// f() 返回两个值
var a, b = f() // a 和 b 在 x 之前同时初始化

func f() (int, int) {
	return 0, 0
}

var (
	c = e + d // == 9
	d = h()   // == 4
	e = h()   // == 5
	g = 3     // == 5
)

func h() int {
	g++
	return g
}

var y int //= I(T{}).jk() // y 隐性依赖 j 和 k, 依赖分析无法发现此依赖
func init() {
	y = I(T{}).jk()
}

var j = k
var k = 1

type I interface{ jk() int }
type T struct{}

func (T) jk() int { return j + k }

func main() {
	fmt.Println(y, j, k)
}
