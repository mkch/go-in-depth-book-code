package main

import (
	"net/http"
	"net/url"
)

type Map interface {
	Get(string) string // 如果删掉此行,下面 m.Get()无法编译
	// 虽然 Values 和 Header 都包含 Get 方法
	// 但这并不增加 Map 的方法集
	url.Values | http.Header
}

func F[T Map](m T, key string) {
	_ = m[""]
	_ = m.Get(key)
}
