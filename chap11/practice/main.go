package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/delete", handleDelete)
	log.Println(http.ListenAndServe(":8080", nil))
}

// handleDelete 处理删除请求.
// 被删除的资源名称从 URL 的 "name" query 中读取.
func handleDelete(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Query().Get("name")
	deleteFromDB(res)
	go deleteFromFS(res)
}

// deleteFromDB 从数据库中删除资源 res.
func deleteFromDB(res string) { /* 代码省略*/ }

// deleteFromFS 从文件系统中删除 res 对应的所有文件.
// 可能耗时较长.
func deleteFromFS(res string) { /* 代码省略*/ }
