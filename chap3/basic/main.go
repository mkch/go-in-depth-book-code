package main

import (
	"fmt"
)

const (
	_ = 'a'          //  U+061 (0x61)
	_ = '中'          // U+4E2D (0x4E2D)
	_ = 'ä'          // U+00E4 (0xE4)
	_ = '\x61'       // U+061 'a'
	_ = '\u4e2d'     // U+4E2D '中'
	_ = '\U0001F600' // U+1F600 '😀'
	_ = '\141'       // U+061 'a'
	_ = '\a'         // U+0007 ^G 响铃
	_ = '\b'         // U+0008 ^H 退格
	_ = '\f'         // U+000C ^L 换页
	_ = '\n'         // U+000A ^J 换行
	_ = '\r'         // U+000D ^M 回车
	_ = '\t'         // U+0009 ^I TAB
	_ = '\v'         // U+000B ^K 垂直 TAB
	_ = '\\'         // U+005C \
	_ = '\''         // U+0027 单引号
)

func main() {
	var n int         // n 类型为 int, 初始值为 0
	var n1 uint64 = 1 // n2 类型为 uint64, 初始值为 1
	var n2 = "abc"    // n3 类型为 string, 初始值为 "abc"
	// 也可以一次性声明多个变量
	var (
		b byte = 0xFF
		c      = '中'
	)

	_, _, _, _, _ = n, n1, n2, b, c

	fmt.Println(State1, State2, State3, State4, State5, x, State6, y)

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
