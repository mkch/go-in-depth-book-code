package session

import (
	"context"
	"net/http"
)

type ctxKey struct{}

// 1
var key ctxKey

type Handler struct {
	http.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 2
	// 从 r.Cookies() 中取出 session data
	// 校验 session data
	// 通过 w 刷新 session
	var session = "session data"
	r = r.WithContext(context.WithValue(r.Context(), key, session))
	h.Handler.ServeHTTP(w, r)
}

// Get 得到此次请求所对应的 Session 数据
func Get(r *http.Request) any {
	return r.Context().Value(key) // 3
}
