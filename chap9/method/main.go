package main

import (
	"fmt"
	"math"
)

// Point 代表二维坐标系上的一个点.
type Point struct {
	X, Y float64
}

// Distance 返回 p 到坐标原点的距离.
func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Offset 把 p 移动到偏移 (cx, cy) 的位置处.
func (p *Point) Offset(cx, cy float64) {
	p.X += cx
	p.Y += cy
}

func main() {
	pt := Point{3, 4}
	fmt.Println(pt.Distance())
	pt.Offset(1, -1)

	// 绑定到值的方法可以用在临时变量上
	Point{3, 4}.Distance()
	// 编译错误！绑定到指针的方法不可用在临时变量上
	//func() Point { return Point{3, 4} }().Offset(1, 1)
}
