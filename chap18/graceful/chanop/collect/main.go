package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 良好的编码习惯

	// shutdown 关闭所有 worker 并等待其退出
	shutdown := func() {
		cancel()
		group.Wait()
	}

	startWorker(ctx, 0)
	startWorker(ctx, 1)
	fmt.Printf("Received %v\n", <-results)
	shutdown()
	fmt.Println("exit")
}

var group sync.WaitGroup

// results 为 worker 回报结果的通道
var results = make(chan int)

// startWorker 启动一个 worker goroutine
func startWorker(ctx context.Context, id int) {
	group.Add(1)
	go func() {
		defer group.Done()
		select {
		case <-ctx.Done():
			// ctx 已完成
		case results <- doWork(id):
		}
		// 执行善后
		doCleanup(id)
		// 正常退出
	}()
}

func doCleanup(id int) {
	fmt.Printf("#%v: cleaning up...\n", id)
}

func doWork(id int) int {
	n := id*100 + 1
	fmt.Printf("#%v working: result %v...\n", id, n)
	return n
}
