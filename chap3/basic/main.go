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

	RawString()
	InterpretedString()

	ShortDecl()

}

func Temp() {

	// p1 为一个临时变量的地址
	// 相当于:
	// var temp int
	// var p1 = &temp
	var p1 = new(int)

	// p2 为一个临时变量的地址
	// 相当于:
	// var temp = []int{1, 2, 3}
	// var p2 = &temp
	var p2 = &[]int{1, 2, 3}

	_, _ = p1, p2
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

func RawString() {
	str := `a
b\n中`
	fmt.Printf("%X\n", str)
}

func InterpretedString() {
	str := "中文"                        // 字面量
	str2 := `中文`                       // 原始字面量
	str3 := "\u4E2D\u6587"             // UTF-8 码位
	str4 := "\U00004E2D\U00006587"     // UTF-8 码位
	str5 := "\xE4\xB8\xAD\xE6\x96\x87" // 16 进制 UTF-8 字节序列
	str6 := "\344\270\255\346\226\207" // 8 进制 UTF-8 字节序列
	fmt.Println(str, str2, str3, str4, str5, str6)
}

func ShortDecl() {
	var a = 1
	if true {
		// 声明了两个新变量 a, b
		// 其中 a 遮蔽了外层声明的 a
		a, b := 10, 2
		fmt.Println(a)
		_ = b
	}
	fmt.Println(a)
}
