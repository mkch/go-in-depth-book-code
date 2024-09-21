package pkg

import "fmt"

// 要实现 Iface 接口, 请在一个 struct 中嵌入 Impl
type Iface interface {
	F()
	defaultImpl()
}

// Impl 提供了 Iface 的某些底层实现
type Impl struct{}

func (Impl) defaultImpl() {
	fmt.Println("Must do this")
}

// Use 使用一个 Iface
func Use(i Iface) {
	i.defaultImpl() // 调用默认实现
	i.F()
}
