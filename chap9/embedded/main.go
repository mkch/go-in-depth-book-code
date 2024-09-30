package main

import (
	"math"
)

// Point 代表二维坐标系上的一个点
type Point struct {
	X, Y float64
}

// Distance 返回 p 到坐标原点的距离
func (p Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Offset 把 p 移动到偏移 (cx, cy) 的位置处
func (p *Point) Offset(cx, cy float64) {
	p.X += cx
	p.Y += cy
}

// ColorPoint 是一个带颜色的 Point
type ColorPoint struct {
	Point // 嵌入字段
	color string
}

// SetColor 设置 p 的颜色为 clr
// 如果 clr 不为合法的颜色则引发 panic
func (p *ColorPoint) SetColor(clr string) {
	// if InvalidColor(clr) {
	// 	panic(fmt.Errorf("SetColor: %v is not a color", clr))
	// }
	p.color = clr
}

// Offset 调用 p.Point.Offset(cx, cy) 并执行更多操作
func (p *ColorPoint) Offset(cx, cy float64) {
	p.Point.Offset(cx, cy) // 1
	// 其他更多操作
	// ...
}

func main() {
	var pt ColorPoint
	pt.X = 2
	pt.Offset(1, 1)
	pt.Distance()
}
