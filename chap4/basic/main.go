package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

func main() {
	Composite()
	If()
	For()
	Switch()
	LabelFor()
	MethodExp()
}
func Composite() {
	a1 := [3]int{0, 10, 0}
	a2 := [3]int{1: 10}
	a3 := [...]int{1: 10, 2: 0}
	fmt.Println(a1, a2, a3)
}

func MethodExp() {
	l := (*bytes.Buffer).Len
	// l 的类型为 func(*bytes.Buffer) int
	_ = l(&bytes.Buffer{})

	r := io.Reader.Read
	// r 的类型为 func(io.Reader, []byte)(int, error)
	_, _ = r(&bytes.Buffer{}, nil)
}

func MethodValue() {
	buf := &bytes.Buffer{}
	ml := buf.Len
	// ml 的类型为 func() int
	_ = ml()

	mr := buf.Read
	// mr 的类型为 func([]byte) (int, error)
	_, _ = mr(nil)
}

func If() {
	if n, err := strconv.Atoi("1"); err != nil {
		// 处理错误
	} else {
		_ = n // 使用 n
	}

	{
		n, err := strconv.Atoi("1")
		if err != nil {
			// 处理错误
		} else {
			_ = n // 使用 n
		}
	}
}

func For() {
	var cond bool
	if cond {
		// 循环体
	}

	for i := 0; i < 10; i++ {
		// 循环体
		fmt.Println(i)
	}
	for i := range 10 {
		fmt.Println(i)
	}
}

func F() string { return "" }

func Switch() {
	var str string = F()
	switch str {
	case "a":
		// 处理 str == "a" 的情况
	case "b":
		// 处理 str == "b" 的情况
		fallthrough
	case "c":
		// 处理 str == "b" 或 str == "c" 的情况
	default:
		// 处理其他情况
	}

	switch str := F(); str {
	// ...
	}

	var v any = F()
	switch v.(type) {
	case int:
		// 处理 v 为 int 的情况
	case string:
		// 处理 v 为 string 的情况
	}

	switch {
	case v == 1:
		// ...
	case str == "":
		// ...
	}
	// 等价于
	if v == 1 {
		// ...
	} else if str == "" {
		// ...
	}
}

func LabelFor() {
Loop: // 为 for 循环定义一个标签
	for i := range 10 {
		switch i {
		case 1, 3, 5:
			// 处理 1,3,5
		case 2, 4:
			break Loop // 跳出 for 循环
		default:
			// 处理其他情况
		}
	}
}

func Select() {
	var ch1 chan int = make(chan int)
	var ch2 chan int = make(chan int)
	select {
	case v1 := <-ch1:
		_ = v1 // 使用 v1
	case v2, ok := <-ch2:
		if ok {
			_ = v2 // 使用
		}
	case ch1 <- 1:
		// 在成功向 ch1 发送 1 后执行
	default:
		// 当以上所有分支均会阻塞时执行
	}
}

type State uint

const (
	State1    State = iota                       // 0 (iota = 0)
	State2                                       // 1 (iota = 1)
	State3    State = 255                        // 255 (iota = 2, 未用)
	State4          = iota + iota                // 6 (iota = 3)
	State5, x       = 1 << iota, 1 << (iota + 1) // 16,32 (iota = 4)
	_, _                                         // (iota = 5, 未用)
	State6, y                                    // 64, 128 (iota = 6)
)
