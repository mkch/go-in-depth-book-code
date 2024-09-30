package main

import (
	"fmt"
	"net/url"
)

func main() {
	concept()
	implicitIndirection()
	sliceArrayPtr()
	compositeLiteral()
}

func concept() {
	var n int
	p1 := &n          // 通过 & 运算符取得指针
	p2 := new(string) // 通过 new() 取得指针

	// %p 输出指针的值. %T 输出指针的类型.
	fmt.Printf("%p %[1]T\n", p1)
	fmt.Printf("%p %[1]T\n", p2)

	// *p1 代表 p1 指向的变量 n
	// 为 *p1 赋值即是对 n 赋值
	*p1 = 1
	fmt.Println(*p1, n)
	// 同样修改 n 也是修改 *p1
	n = 2
	fmt.Println(*p1, n)
	// *p2 为 string 的 zero value("")
	fmt.Println(*p2)
	// *p2 在代码中没有对应的变量
	// *p2 自身可以当作一个 string 变量使用
	*p2 = "abc"
	fmt.Println(*p2)

	// 指针的零值为 nil
	var p3 *int
	fmt.Printf("%p %[1]v\n", p3)
	// nil 指针不指向任何有效数据,因此不可脱引用(间接访问)
	// 下面代码将引发 panic
	//_ = *p3
	if p3 == nil {
		// p3 指向的值不存在,不可使用 *p3
	} else {
		// 可以使用 *p3
		fmt.Println(*p3)
		*p3 = 1
	}

	pp1 := &p1
	fmt.Printf("%p %[1]T\n", pp1)

	ppp1 := &pp1
	fmt.Printf("%p %[1]T\n", ppp1)

	var _ int = ***ppp1
}

type S struct {
	Field int
}

func (s S) Method() {}

// PS 是一个"定义类型"(defined type)
// 其底层类型为 *S
type PS *S

func implicitIndirection() {
	var s S
	p := &s
	(*p).Field = 0 // 1: 相当于 s.Field = 0
	(*p).Method()  // 2: 相当于 s.Method()

	p.Field = 0 // 3: 1 的简化写法
	p.Method()  // 4: 2 的简化写法

	var p2 PS = &s
	p2.Field = 0 // 5: 和 3 类似
	//p2.Method()    // 6: !!编译错误!!
	(*p2).Method() // 7: 这样是可以的
}

func sliceArrayPtr() {
	var s = []int{1, 2, 3}

	// a 的内容是 s[0:2] 的拷贝
	a := ([2]int)(s)
	a[0] = 10 // 此修改不影响 s
	fmt.Println(s, a)

	// pa 的内容是 s[0:2] 的引用
	// 即 pa 所指数组和 s 共享存储空间
	pa := (*[3]int)(s)
	pa[0] = 10 // 此修改会影响 s 的内容
	s[1] = 20  // 此此修改会影响 *pa 的内容
	fmt.Println(s, *pa)
}

type S2 struct {
	URL url.URL
}

func compositeLiteral() {
	p1 := &[]byte{0}
	// 相当于
	var hidden1 = []byte{0}
	p1 = &hidden1

	type S struct{ F int }
	p2 := &S{F: 1}
	// 相当于
	var hidden2 = S{F: 1}
	p2 = &hidden2

	_, _ = p1, p2

	urlStr := (&url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/path1",
		RawQuery: url.Values{"key": {"a", "b"}}.Encode(),
	}).String()

	//url.URL{ /* ... */ }.String() // !!语法错误!!

	fmt.Println(urlStr)
}
