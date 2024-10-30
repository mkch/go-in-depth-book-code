package main

import "fmt"

func main() {
	var creator = new(ContreteCreator)
	creator.AnOperation() // 这里会 panic, 而不是显示 "Product".
}

type Product interface {
	Product()
}

type Creator struct {
}

func (c *Creator) FactoryMethod() Product {
	panic("需要子类覆盖") // !! 错误: "子类" 无法覆盖 !!
}

func (c *Creator) AnOperation() {
	// !! 错误: 此处永远调用 (*BaseCreator).FactoryMethod() !!
	product := c.FactoryMethod()
	product.Product()
}

type ContreteCreator struct {
	Creator // !! 嵌入不具有多态性 !!
}

// !! 错误: 试图覆盖 (*BaseCreator).FactoryMethod() !!
func (c *ContreteCreator) FactoryMethod() Product {
	return new(ContreteProduct)
}

// ContreteProduct 实现了 Product
type ContreteProduct struct{}

func (*ContreteProduct) Product() {
	fmt.Println("Product")
}
