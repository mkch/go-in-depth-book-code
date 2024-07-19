package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str1 = "abc"
	var str2 = "abc"
	fmt.Println(unsafe.StringData(str1))
	fmt.Println(unsafe.StringData(str2))
}

func f() {
	var str = "abc"
	var slice = []byte{'a', 'b', 'c'}

	slice[0] = 0 // slice 是"可变"的
	str = "abc"  // 不可变类型也可以被重新赋值
	//str[0] = 0   // 1: 编译错误!! 不能改变其底层数据
	_ = str

}
