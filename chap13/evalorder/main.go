package main

import "fmt"

func main() {
	// 求值顺序为: g(), h(), i(), k()
	f(g(), h(), k(i()))
}

var a int

func g() int      { a++; fmt.Println("g"); return a }
func h() int      { a *= 3; fmt.Println("h"); return a }
func i() int      { a += 5; fmt.Println("i"); return a }
func k(n int) int { a += n; fmt.Println("k"); return a }

func f(n1 int, n2 ...int) {
	fmt.Println(n1, n2)
}

func f0() {
	return
}

func f1() int {
	return 0
}

func f2() (int, string) {
	return 1, fmt.Sprintf("%#v", f2)
}

func f3() (int, string) {
	return f2()
}

func f4() (n int, err error) {
	n = 1
	err = nil
	return
}

func f5() (n int, err error) {
	if _, err := f4(); err != nil {
		//return // !语法错误! 此处局部变量 err 遮蔽(shadows)了返回值 err
	}
	return
}
