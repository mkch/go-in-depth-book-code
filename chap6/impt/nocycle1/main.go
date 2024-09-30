package main

import "example/demo/drawing"

func main() {
	c := drawing.NewCanvas(1024, 720)
	b := drawing.NewBitmap(100, 100)
	b.SetPixel(0, 0, 0x1A2B3C)
	b.Draw(10, 10, c)

	_ = c
}
