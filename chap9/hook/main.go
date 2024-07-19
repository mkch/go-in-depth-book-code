package main

import "fmt"

func main() {
	hook := func() Product { return new(ConcreteProduct) }
	factory := &Factory{hook}
	factory.AnOperation()
}

type Product interface {
	ProductOp()
}

type ConcreteProduct struct{}

func (p *ConcreteProduct) ProductOp() {
	fmt.Println("ProductOp")
}

type Factory struct {
	Producer func() Product // 此处使用一个函数代替单方法接口
}

func (f *Factory) AnOperation() {
	product := f.Producer()
	product.ProductOp()
}
