package main

func main() {
	var s = []int{1, 2, 3}
	_ = append([]int(nil), s...) // OK
	//_ = min(s...)                // 错误
}
