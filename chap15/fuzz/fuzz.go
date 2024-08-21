package fuzz

import (
	"strings"
)

// SplitLast 在 str 中寻找最后一个 del, 如果找到将以 del 为分隔符
// 把 str 分割为 left 和 right 两部分.
// 如果未找到 del 或 del 超出 ASCII 范围, 则返回的 left 为 str 自身,
// right 为空.
func SplitLast(str string, del byte) (left, right string) {
	if del > 0x7F {
		left = str
		return
	}
	i := strings.LastIndexByte(str, del)
	if i == -1 {
		left = str
		return
	}
	left, right = str[:i], str[i+1:]
	return
}
