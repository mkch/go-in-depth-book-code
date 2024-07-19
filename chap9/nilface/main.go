package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var nilptr *os.File = nil
	var iface io.Reader = nilptr
	fmt.Println(iface == nilptr) // 1 true
	fmt.Println(iface == nil)    // 2 false

	if r := F(-1); r != nil {
		io.ReadAll(r) // 1 panic
	}
}

// F 创建一个 io.Reader
// 如果创建失败, 返回 nil
func F(n int) io.Reader {
	var b *strings.Reader
	if n >= 0 {
		b = strings.NewReader("abc")
	}
	return b
}
