package main

import (
	"ctxval/session"
	"log"
	"net"
	"net/http"
)

func main() {
	// NOP.
}

func Session() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := session.Get(r)
		_ = s // 使用 s
	})

	http.ListenAndServe(":8888",
		&session.Handler{Handler: http.DefaultServeMux})
}

func HttpContextKey() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server := r.Context().Value(http.ServerContextKey).(*http.Server)
		localAddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)
		// 使用 server 和 localAddr
		log.Printf("%p, %v\n", server, localAddr)
	})

	http.ListenAndServe(":8888", nil)
}
