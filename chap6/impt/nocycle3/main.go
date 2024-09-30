package main

import (
	"example/demo/bitmap"
	"example/demo/canvas"
)

func main() {
	c := canvas.New(1024, 720)
	b := bitmap.New(100, 100)
	b.SetPixel(0, 0, 0x1A2B3C)
	c.DrawBitmap(10, 10, b)

}
