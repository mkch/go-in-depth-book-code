package main

import (
	"fmt"
	"math"
)

func main() {
	var shapes = []Shape{Rectangle{1, 2}, Circle{3}}
	fmt.Println(AvgArea(shapes))
}

// AvgArea 计算 shapes 的面积的平均值
func AvgArea(shapes []Shape) float32 {
	var sum float32
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum / float32(len(shapes))
}

// Rectangle 代表一个矩形
type Rectangle struct {
	Width, Height int // 长和宽
}

// Area 计算 r 的面积
func (r Rectangle) Area() float32 {
	return float32(r.Width) * float32(r.Height)
}

type Circle struct {
	Radius int
}

func (c Circle) Area() float32 {
	return math.Pi * float32(c.Radius*c.Radius)
}

type Shape interface {
	Area() float32
}
