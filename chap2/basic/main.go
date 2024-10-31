package main

import (
	"fmt"
	"io"
	"main/pkg"
)

func main() {
	var 变量1 = pkg.Δ()
	_ = 变量1

	var max = 1 // 遮蔽了内建的 max()
	_ = max
	// var break =2 // 非法

	_, err := fmt.Print("Hello\n")
	if err != nil {
		// 处理错误
	}

	Shift()
	BinaryBitwise()
	UnaryBitwise()
}

func Shift() {
	var i = 192         //      二进制 11000000
	n1 := int8(i) >> 1  // -32 (二进制 11100000)
	n2 := uint8(i) >> 1 // 96, (二进制 01100000)
	fmt.Printf("%b %b\n", uint8(n1), n2)
}

func BinaryBitwise() {
	var a int8 = 10 //     二进制 1010
	var b int8 = 12 //     二进制 1100
	and := a & b    // 8  (二进制 1000)
	or := a | b     // 14 (二进制 1110)
	xor := a ^ b    // 6  (二进制 0110)

	fmt.Printf("%v: %[1]b  %v: %[2]b  %v: %[3]b\n", and, or, xor)
}

func UnaryBitwise() {
	var a uint8 = 10 //       二进制 00001010
	ca := ^a         // 245  (二进制 11110101)

	var b int8 = -10 //       二进制 11110110
	cb := ^b         // 9    (二进制 00001001)

	fmt.Printf("%v: %[1]b %v: %[1]b\n", ca, cb)
}

func f() io.Reader {
	return nil
}

func TypeAssertion() {
	var x any = 1
	n := x.(int) // n 的类型将为 int, 值为 1

	var r io.Reader = f()
	// 如果 r 的动态类型没有实现 io.Writer
	// 此语句将引发 panic
	rw := r.(io.Writer)

	var v, ok interface{} = x.(int) // dynamic types of v and ok are T and bool

	_, _, _, _ = n, rw, v, ok
}

func Precedence() {
	var a = 100
	a <<= 3
	_ = 'a' //  U+061
	_ = '中' // U+4E2D
	_ = 'ä' // U+00E4

	_ = '\x61'       // U+061 'a'
	_ = '\u4e2d'     // U+4E2D '中'
	_ = '\U0001F600' // U+1F600 '😀'
	_ = '\141'       // U+061 'a'

	_ = '\a' // U+0007 ^G 响铃
	_ = '\b' // U+0008 ^H 退格
	_ = '\f' // U+000C ^L 换页
	_ = '\n' // U+000A ^J 换行
	_ = '\r' // U+000D ^M 回车
	_ = '\t' // U+0009 ^I TAB
	_ = '\v' // U+000B ^K 垂直 TAB
	_ = '\\' // U+005C 字符\
	_ = '\'' // U+0027 单引号字符
}
