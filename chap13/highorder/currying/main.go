package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Add 返回 a+b 的和.
func Add(a, b int) int {
	return a + b
}

// AddCurrying 是柯里化的 Add.
// AddCurrying(a)返回一个函数 f, 调用 f(b) 将得到 Add(a,b) 的结果.
func AddCurrying(a int) (f func(b int) int) {
	return func(n int) int {
		return Add(a, n)
	}
}

func main() {
	sum1 := Add(1, 2)
	sum2 := AddCurrying(1)(2)

	fmt.Println(sum1, sum2)

	http.HandleFunc("/", HelloHandler(log.Default()))
	//http.ListenAndServe(":8080", nil)
}

// HelloWithLog 记录 r 中的路径信息到 logger, 并向 w 中写入 200 状态码和 "Hello".
func HelloWithLog(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	// 记录日志
	logger.Printf("path=%v\n", r.URL.Path)
	// 处理请求
	io.WriteString(w, "Hello")
}

// HelloHandler 使用 logger 对 HelloWithLog 进行柯里化,
// 得到一个可用于 http.HandleFunc() 的函数.
func HelloHandler(logger *log.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		HelloWithLog(logger, w, r)
	}
}
