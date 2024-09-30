package main

import (
	"fmt"
)

func main() {
	s := []byte("It is OKAY")
	// is 的内容为 "is"
	is := s[3:5]
	isEvil := appendEvil(is)
	// isEvil 的内容为 "is evil"
	fmt.Printf("%s\n", isEvil)
	// s 的内容变成了 "It is evil"
	fmt.Printf("%s\n", s)
}

// appendEvil 把 s1 和 " evil" 连接并返回连接后的结果.
func appendEvil(s1 []byte) []byte {
	return append(s1, []byte(" evil")...)
	// 使用以下三种方法都可以避免 s 被修改
	// return append(slices.Clone(s1), []byte(" evil")...)
	// return append(slices.Clip(s1), []byte(" evil")...)
	// return append(s1[:len(s1):len(s1)], []byte(" evil")...)
}
