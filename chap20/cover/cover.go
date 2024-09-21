package cover

// Ellipsis 返回 str 首部最多 max 个字符组成的字符串.
// 如果 str 的字符数不大于 max, 则返回 str 自身.
// 否则返回 str 前 max-1 个字符和省略号(…)组成的 string.
// 如果 max 小于等于 1, 则返回 str 自身.
func Ellipsis(str string, max int) string {
	if max <= 1 {
		return str
	}
	runes := []rune(str)
	if len(runes) <= max {
		return str
	}
	return string(append(runes[:max-1], '…'))
}

func F(cond bool) string {
	var str string
	if cond {
		str = "abc"
	}
	return str[:1]
}
