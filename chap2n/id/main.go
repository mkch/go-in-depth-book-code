package main

import (
	"main/pkg"
)

func main() {
	var 变量1 = pkg.Δ()
	_ = 变量1

	var max = 1 // 遮蔽了内建的 max()
	_ = max
	// var break =2 // 非法
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
