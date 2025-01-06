package bitmap

import (
	"example/demo/bitmap/bitmapdata"
	"example/demo/canvas"
)

// Bitmap 代表一幅位图.
type Bitmap struct {
	*bitmapdata.BitmapData
}

// Draw 把 b 绘制到 canvas 上的 (x0, y0) 处.
func (b *Bitmap) Draw(x0, y0 int, canvas *canvas.Canvas) {
	cx, cy := b.Dimensions()
	for x := 0; x < cx; x++ {
		for y := 0; y < cy; y++ {
			canvas.DrawPixel(x0+x, y0+y, b.Pixel(x, y))
		}
	}
}

// New 创建一幅宽度为 cx, 高度为 cy 的位图.
func New(cx, cy int) *Bitmap {
	return &Bitmap{
		bitmapdata.New(cx, cy),
	}
}
