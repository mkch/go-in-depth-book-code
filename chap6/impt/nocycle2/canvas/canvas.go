package canvas

import (
	"example/demo/bitmap/bitmapdata"
)

// Canvas 是一块画布
type Canvas struct {
	backing *bitmapdata.BitmapData
}

// New 创建一个宽度为 cx, 高度为 cy 的 Canvas
func New(cx, cy int) *Canvas {
	return &Canvas{bitmapdata.New(cx, cy)}
}

// DrawPixel 在 c 的坐标 (x,y) 处画一个指定颜色的像素
func (c *Canvas) DrawPixel(x, y int, color uint32) {
	c.backing.SetPixel(x, y, color)
}
