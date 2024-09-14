package main

type Point struct{ X, Y int }

var (
	a = [3]int{}
	b = []byte{1, 2, 3}
	c = Point{X: 1, Y: 2}
	d = []Point{{X: 3}, {X: 5, Y: 6}}
	e = map[int]string{1: "One", 2: "two"}
)

func f() {
	var (
		a int
	)
	a = 1 // 赋值语句
	//b = (a = 2) // !! 语法错误, a=2 是一个语句, 不能求值

	a += 2 // 赋值语句
	// b = a += 2  // !! 语法错误, a+=2 是一个语句, 不能求值

	var ch = make(chan int)
	_ = <-ch // 表达式 (<- 是一个单目操作符)
	ch <- 1  // 语句, 不能求值
}
