package main

import (
	"fmt"
)

func main() {
	basicOp()
	makeMap()
	iteration()
	dynamicKey()
}

func makeMap() {
	m1 := make(map[string]int)
	l1 := len(m1) // 0
	m2 := make(map[string]int, 3)
	l2 := len(m2) // 0

	fmt.Println(l1, l2)
	_, _ = m1, m2

	delete(map[string]int(nil), "")
	clear(map[string]int(nil))
}

// basicOp 演示 map 基本操作
func basicOp() {
	// 1
	var colors = map[string]int{
		"Red":   0xFF0000,
		"Green": 0x00FF00,
		"Blue":  0x0000FF,
	}
	l := len(colors) // l 的值为 3
	// 2: 读取
	red := colors["Red"]         // read 的值为 0xFF0000
	brown, ok := colors["Brown"] // brown 的值为 0, ok 的值为 false
	// 3: 写入(覆盖)
	colors["Blue"] = 255
	// 4: 添加
	colors["Gray"] = 0x808080
	l = len(colors) // l 的值为 4
	// 5: 删除
	delete(colors, "Gray")
	l = len(colors) // l 的值为 3
	// 6: 清空
	clear(colors)
	l = len(colors) // l 的值为 0

	_, _, _, _ = l, red, brown, ok
}

// iteration 演示 map 的遍历.
func iteration() {
	var colors = map[string]int{
		"Red":   0xFF0000,
		"Green": 0x00FF00,
		"Blue":  0x0000FF,
	}
	for key, value := range colors {
		fmt.Printf("%v:%#x\n", key, value)
	}
	for key := range colors {
		fmt.Println(key)
	}
	for _, value := range colors {
		fmt.Println(value)
	}
}

func dynamicKey() {
	var m = make(map[any]int)
	m[1] = 1
	m["2"] = 2
	m[struct{}{}] = 3
	m[(*func())(nil)] = 4
	m[[...]int{1, 2, 3}] = 5
	fmt.Println(m)
}
