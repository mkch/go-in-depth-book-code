package main

import (
	"fmt"
)

func main() {
	var number Number = new(SomeNumber)
	number.Print() // 这里会 panic, 而不是显示"100".
}

type Number interface {
	Value() int
	Print()
}

type NumberBase struct {
}

func (n *NumberBase) Value() int {
	panic("需要子类覆盖")
}

func (b *NumberBase) Print() {
	value := b.Value()
	fmt.Println(value)
}

type SomeNumber struct {
	NumberBase // 错误: 试图继承 NumberBase
}

// 错误: 试图覆盖 NumberBase.Value()
func (n *SomeNumber) Value() int {
	return 100
}
