package bitmapdata

// BitmapData 代表一幅位图的数据.
type BitmapData struct {
	// 实现代码省略
}

// Dimensions 返回 b 的宽度和高度.
func (b *BitmapData) Dimensions() (cx, cy int) {
	return 0, 0 // 实现代码省略
}

// Pixel 取坐标 (x,y) 处的像素颜色.
func (b *BitmapData) Pixel(x, y int) uint32 {
	return 0 // 实现代码省略
}

// SetPixel 设置坐标 (x,y) 处的像素颜色.
func (b *BitmapData) SetPixel(x, y int, color uint32) {
	// 实现代码省略
}

// New 创建宽度为 cx, 高度为 cy 的位图数据.
func New(cx, cy int) *BitmapData {
	return &BitmapData{
		// 实现代码省略
	}
}
