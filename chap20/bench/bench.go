package bench

import (
	"strings"
)

// LastDot 返回 str 中最后一个"."字符的索引
// 如果找不到, 返回 -1
func LastDot(str string) int {
	return strings.LastIndexByte(str, '.')
}
