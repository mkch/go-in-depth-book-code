package main

import "fmt"

func main() {
	var factory AbstractFactory = new(ConcreteFactory)
	var a = factory.CreateProductA()
	a.AbstractProductA()
	var b = factory.CreateProductB()
	b.AbstractProductB()
}

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type AbstractProductA interface {
	AbstractProductA()
}

type AbstractProductB interface {
	AbstractProductB()
}

// ProductA 实现了 AbstractProductA.
type ProductA struct{}

func (*ProductA) AbstractProductA() {
	fmt.Println("AbstractProductA")
}

// ProductB 实现了 AbstractProductB.
type ProductB struct{}

func (*ProductB) AbstractProductB() {
	fmt.Println("AbstractProductB")
}

// ConcreteFactory 实现了 AbstractFactory.
type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProductA() AbstractProductA {
	return new(ProductA)
}

func (f *ConcreteFactory) CreateProductB() AbstractProductB {
	return new(ProductB)
}
