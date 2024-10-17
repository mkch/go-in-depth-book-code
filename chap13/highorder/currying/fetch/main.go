package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Fetch 使用 HTTP GET 方法请求 url. timeout 指定超时时间.
func Fetch(url string, timeout time.Duration) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	return io.ReadAll(resp.Body)
}

// FetchFunc 使用 HTTP GET 方法请求 url.
type FetchFunc func(url string) (data []byte, err error)

// NewFetch 返回一个超时时间为 timeout 的 FetchFunc.
func NewFetch(timeout time.Duration) FetchFunc {
	return func(url string) (data []byte, err error) {
		return Fetch(url, timeout)
	}
}

func UseFetch() {
	var url1 string
	var url2 string
	var url3 string
	var url4 string
	var timeout1 time.Duration
	var timeout2 time.Duration

	Fetch(url1, timeout1)
	Fetch(url2, timeout1)
	// ...

	Fetch(url3, timeout2)
	Fetch(url4, timeout2)
	// ...

	fetch1 := NewFetch(timeout1)
	fetch1(url1)
	fetch1(url2)
	// ...

	fetch2 := NewFetch(timeout2)
	fetch2(url3)
	fetch2(url4)
	// ...
}
