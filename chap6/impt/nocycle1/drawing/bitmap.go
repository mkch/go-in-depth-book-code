package drawing

// Bitmap 代表一幅位图.
type Bitmap struct {
	// 实现代码省略
}

// Width 返回 b 的宽度和高度.
func (b *Bitmap) Dimension() (cx, cy int) {
	return 0, 0 // 实现代码省略
}

// Pixel 取坐标 (x,y) 处的像素颜色.
func (b *Bitmap) Pixel(x, y int) uint32 {
	return 0 // 实现代码省略
}

// SetPixel 设置坐标 (x,y) 处的像素颜色.
func (b *Bitmap) SetPixel(x, y int, color uint32) {
	// 实现代码省略
}

// Draw 把 b 绘制到 canvas 上的 (x0, y0) 处.
func (b *Bitmap) Draw(x0, y0 int, canvas *Canvas) {
	cx, cy := b.Dimension()
	for x := 0; x < cx; x++ {
		for y := 0; y < cy; y++ {
			canvas.DrawPixel(x0+x, y0+y, b.Pixel(x, y))
		}
	}
}

// New 创建一幅宽度为 cx, 高度为 cy 的位图.
func NewBitmap(cx, cy int) *Bitmap {
	return &Bitmap{
		// 实现代码省略
	}
}
