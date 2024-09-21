package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func F(int) (int, error) { return 0, nil }

// 如果 err != nil 则引发 panic, 否则返回 a
func Must[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}
	return a
}

func UseMust() {
	var arg1 int

	// n, err := F(arg1)
	// if err != nil {
	// 	panic(err)
	// }
	// // 使用 n
	// _ = n

	n := Must(F(arg1))
	//使用 n
	_ = n

	n = Must(strconv.Atoi("10"))
	// 使用 n
	_ = n

	file := Must(os.Open("file1"))
	defer file.Close()
	// 使用 file
}

// DecodeBody 把 r 的内容解码后放入 dest 所指向的变量中
// 类似 json.Unmarshal()
func DecodeBody(r io.Reader, dest any) error {
	// 实现代码省略
	return nil
}

// DecodeQuey 把 m 解码后放入 dest 所指向的变量中
// 类似 json.Unmarshal()
func DecodeQuey(m url.Values, dest any) error {
	// 实现代码省略
	return nil
}

// DecodeQuey 把 m 解码后放入 dest 所指向的变量中
// 类似 json.Unmarshal()
func DecodeHeader(m http.Header, dest any) error {
	// 实现代码省略
	return nil
}

// Validate 对 v 进行校验
// 如果 v 是 struct, 校验将根据 v 的 field tag 进行
// https://github.com/go-playground/validator 就是一个类似的校验库
func Validate(v any) error {
	return nil
}

func Decode[T any](w http.ResponseWriter, decode func(T, any) error, arg T, dest any) (ok bool) {
	var err error
	if err = decode(arg, dest); err == nil {
		err = Validate(dest)
	}
	if ok = err == nil; !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}

func Handle(w http.ResponseWriter, r *http.Request) {
	type Body struct{}
	type Query struct{}
	type Header struct{}

	var body Body
	if err := DecodeBody(r.Body, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err = Validate(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var query Query
	if err := DecodeQuey(r.URL.Query(), &query); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err = Validate(&query); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var header Header
	if err := DecodeHeader(r.Header, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err = Validate(&header); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// ...
}

func Handle2(w http.ResponseWriter, r *http.Request) {
	type Body struct{}
	type Query struct{}
	type Header struct{}

	var body Body
	var query Query
	var header Header
	if !Decode[io.Reader](w, DecodeBody, r.Body, &body) ||
		!Decode(w, DecodeQuey, r.URL.Query(), &query) ||
		!Decode(w, DecodeHeader, r.Header, &header) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// ...
}
