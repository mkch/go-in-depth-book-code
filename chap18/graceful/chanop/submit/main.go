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
		cancel()     // 通知所有 worker
		group.Wait() // 等待所有 worker 退出
	}

	startWorker(ctx, 0)
	startWorker(ctx, 1)
	tasks <- 100
	shutdown()
	fmt.Println("exit")
}

var group sync.WaitGroup

// tasks 为向worker提交任务的通道.
var tasks = make(chan int)

// startWorker 启动一个 worker goroutine.
func startWorker(ctx context.Context, id int) {
	group.Add(1)
	go func() {
		defer group.Done()
		select {
		case <-ctx.Done():
			// ctx 已完成
		case task := <-tasks:
			doWork(id, task)
		}
		// 执行善后
		doCleanup(id)
		// 正常退出
	}()
}

func doCleanup(id int) {
	fmt.Printf("#%v: cleaning up...\n", id)
}

func doWork(id int, task int) {
	fmt.Printf("#%v working: task %v...\n", id, task)
}
