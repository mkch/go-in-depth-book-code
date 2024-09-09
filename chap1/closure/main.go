package main

import "fmt"

func F() (get func() int, set func(int)) {
	var n int
	// get 和 set 通过闭包共享变量 n
	get = func() int { return n }
	set = func(i int) { n = i }
	return
}

func main() {
	get, set := F()
	set(1)
	fmt.Println(get())
	set(2)
	fmt.Println(get())
}
