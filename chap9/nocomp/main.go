package main

// NoCompare 类型不可比较.
type NoCompara struct {
	_  [0]func() // 1
	F1 string
}

func main() {
	var s NoCompara
	var s2 NoCompara
	// if s == s2 { // 编译错误! S 不可比较
	// }
	_, _ = s, s2
}
