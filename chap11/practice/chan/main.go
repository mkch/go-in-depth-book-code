package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

// 等待从文件系统中删除的资源
var queue = make(chan string, 64)

func main() {
	var server = http.Server{Addr: ":8080"}

	var wg sync.WaitGroup
	wg.Add(1)
	// 负责删除文件的 goroutine
	go func() {
		defer wg.Done()
		for res := range queue {
			deleteFromFS(res)
		}
	}()

	http.HandleFunc("/delete", handleDelete)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// 用于接收 os.Interrupt 信号 (比如 Ctrl+C)
	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
	// 优雅地关闭 server
	server.Shutdown(context.Background())
	close(queue)
	// 等待负责删除文件的 goroutine 结束
	wg.Wait()
}

// handleDelete 处理删除请求.
// 被删除的资源名称从 URL 的 "name" query 中读取.
func handleDelete(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Query().Get("name")
	deleteFromDB(res)
	// 把 res 放入队列等待删除
	queue <- res
}

// deleteFromDB 从数据库中删除资源 res.
func deleteFromDB(res string) { /* 代码省略*/ }

// deleteFromFS 从文件系统中删除 res 对应的所有文件.
// 可能耗时较长.
func deleteFromFS(res string) {
	log.Printf("Deleting %v\n", res)
	/* 代码省略*/
}
