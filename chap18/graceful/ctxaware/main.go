package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 良好的编码习惯

	result := startWorker(ctx, "https://baidu.com")
	time.Sleep(time.Millisecond * 200)
	cancel()              // 关闭 worker
	fmt.Println(<-result) // 读出结果(成功或失败), 兼具等待 worker 退出
	fmt.Println("exit")
}

// startWorker 启动一个 worker goroutine 使用 HTTP GET 方法请求 url.
// result 为 worker 回报结果的通道.
// 如果请求成功, worker 向 result 中写入页面内容,
// 如果请求失败, worker 向 result 中写入错误原因.
func startWorker(ctx context.Context, url string) (result chan string) {
	result = make(chan string)
	go func() {
		req, err := http.NewRequestWithContext(ctx,
			http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			result <- err.Error()
			return
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			result <- err.Error()
			return
		}
		result <- string(body)
	}()
	return
}
