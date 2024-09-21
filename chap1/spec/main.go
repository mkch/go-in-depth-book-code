package main

func MinMax() {
	var s = []int{1, 2, 3}
	s = append([]int{0}, s...) // OK
	//n := min(s...)             // 错误
}
