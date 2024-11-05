package main

import "fmt"

func main() {
	var concreteMethod = func() Product { return new(ContreteProduct) }
	var creator = &Creator{FactoryMethod: concreteMethod}
	creator.AnOperation()
}

type Product interface {
	Product()
}

type Creator struct {
	FactoryMethod func() Product
	// ...
}

func (c *Creator) AnOperation() {
	product := c.FactoryMethod()
	product.Product()
}

// ContreteProduct 实现了 Product/
type ContreteProduct struct{}

func (*ContreteProduct) Product() {
	fmt.Println("Product")
}
