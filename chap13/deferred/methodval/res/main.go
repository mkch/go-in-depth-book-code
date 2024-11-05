package main

import (
	"fmt"
)

// Resource 代表某种类型的资源.
type Resource struct {
	name string
	// ...
}

// NewResource 创建一个名称为 name 的资源.
func NewResource(name string) *Resource {
	fmt.Printf("Create resource %v\n", name)
	return &Resource{name}
}

// Close 方法关闭资源 res.
func (res *Resource) Close() {
	fmt.Printf("Close resource %v\n", res.name)
}

func main() {
	res := NewResource("1") // 1 创建资源1
	//defer res.Close()       // 2 保证 res 会被关闭
	defer func() { res.Close() }() // 2 保证 res 会被关闭

	// 使用资源1
	// ...

	// 3 关闭资源1, 并打开资源2
	res.Close()
	res = NewResource("2")
}
