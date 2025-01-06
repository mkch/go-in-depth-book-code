package canvas

import "example/demo/bitmap"

// Canvas 是一块画布.
type Canvas struct {
	backing *bitmap.Bitmap
}

// New 创建一个宽度为 cx, 高度为 cy 的 Canvas.
func New(cx, cy int) *Canvas {
	return &Canvas{bitmap.New(cx, cy)}
}

// DrawPixel 在 c 的坐标 (x,y) 处画一个指定颜色的像素.
func (c *Canvas) DrawPixel(x, y int, color uint32) {
	c.backing.SetPixel(x, y, color)
}

// DrawBitmap 在 c 的 (x,y) 坐标处画出 bitmap.
func (c *Canvas) DrawBitmap(x, y int, bitmap *bitmap.Bitmap) {
	cx, cy := bitmap.Dimensions()
	for x1 := 0; x < cx; x++ {
		for y1 := 0; y < cy; y++ {
			c.DrawPixel(x1+x, y1+y, bitmap.Pixel(x, y))
		}
	}
}
