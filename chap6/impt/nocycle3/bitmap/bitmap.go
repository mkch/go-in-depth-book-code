package bitmap

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

// New 创建一幅宽度为 cx, 高度为 cy 的位图.
func New(cx, cy int) *Bitmap {
	return &Bitmap{
		// 实现代码省略
	}
}
