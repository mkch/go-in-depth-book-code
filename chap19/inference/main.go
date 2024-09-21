package main

import "io"

func F[T any](arg1 T, arg2 T) {}

func F1(arg1 io.Reader, arg2 io.Reader) {}

func main() {
	var r io.Reader
	var rc io.ReadCloser

	//F(r, rc) // !!语法错误!! 类型推导失败
	F[io.Reader](r, rc) // 手工加上类型参数即可.
	F1(r, rc)           // 上一行和此行等价
}
