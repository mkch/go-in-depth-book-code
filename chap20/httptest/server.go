package server

import (
	"io"
	"net/http"
)

// handleConcat 把 r 的 query 中的 a 和 b 连接起来写入 w.
func handleConcat(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	io.WriteString(w, a+b)
}

// mustHttp2 在 r 的 HTTP 版本不为 HTTP/2 时返回 StatusBadRequest.
func mustHttp2(w http.ResponseWriter, r *http.Request) {
	if r.ProtoMajor != 2 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// 使用 HTTP2 的功能
		// ...
		w.WriteHeader(http.StatusOK)
	}
}
