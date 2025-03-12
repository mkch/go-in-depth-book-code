package main

import (
	"fmt"
	"math/bits"
	"unicode/utf8"
)

func main() {
	encoding()
}

func encoding() {
	// 'x'(其中 x 为一个字符) 的类型为 rune, 值为字符的 code point
	// "x"(其中 x 为一个字符) 的类型为 string, 值为字符的 UTF-8 编码
	fmt.Printf("%X %X\n", 'a', "a")
	fmt.Printf("%X %X\n", '中', "中")
}

// CopyString 类似内建的 copy 函数,
// 但要求复制到 dest 中的内容必须为有效的 UTF-8 序列.
func CopyString2(dest []byte, src string) int {
	// 如果 dest 可容纳 src, 则直接 copy
	if len(src) <= len(dest) {
		return copy(dest, src)
	}

	// 取出 src 中对应 dest 尾部的字节
	// 如果该字节第一位是 0, 是单字节字符, 不会截断, 直接 copy
	if src[len(dest)-1]&0b10000000 == 0 {
		return copy(dest, src)
	}

	// 把 src 从 len(dest)-1 处截断, 以下代码保证不出现半个字符
	for i := len(dest) - 1; i >= 0; i-- {
		// 向前找到字符边界
		if utf8.RuneStart(src[i]) {
			// 取出此字符的字节数(即字符首个字节中前导 1 的个数)
			// bits.LeadingZeros32() 是取前导 0 的个数, 所以要用 ^ 取反
			// << 24 把此字节移到 uint32 的首部
			count := bits.LeadingZeros32(uint32(^src[i]) << 24)
			// 如果此字符没有被截断, 直接复制
			if len(dest)-i == count {
				return copy(dest, src)
			} else { // 如果截断了, 则只复制 i 前面的部分
				return copy(dest, src[:i])
			}
		}
	}
	return 0
}

func CopyString(dest []byte, src string) int {
	// 如果 dest 可容纳 src, 则直接 copy
	if len(src) <= len(dest) {
		return copy(dest, src)
	}
	// 在 src 中遍历字符
	for i, r := range src {
		last := i + utf8.RuneLen(r)
		// 如果 dest 尾部处的字符没有截断, 直接复制
		if last == len(dest) {
			return copy(dest, src)
		} else if last > len(dest) { // 截断了, 只复制 i 前面的部分
			return copy(dest, src[:i])
		}
	}
	return 0
}
