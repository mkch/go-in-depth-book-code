package session

import (
	"context"
	"net/http"
)

type ctxKey struct{}

var key ctxKey

type Handler struct {
	http.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), key, cookieSession(w, r)))
	h.Handler.ServeHTTP(w, r)
}

// cookieSession 从 cookie 中取出 session 数据
func cookieSession(w http.ResponseWriter, r *http.Request) any {
	// 从 r.Cookie() 中取出 session key
	// 通过 session key 取出对应的 session data
	// 或者使用 http.SetCookie(w, ...) 写出 session key
	return "session data" // 假设这是取出的 session data
}

// Get 得到此次请求所对应的 Session 数据
func Get(r *http.Request) any {
	return r.Context().Value(key)
}
