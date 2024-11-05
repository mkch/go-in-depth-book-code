package main

import (
	"context"
	"fmt"
	"sync"
	"time"
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
	time.Sleep(time.Second * 3)
	shutdown()
	fmt.Println("exit")
}

var group sync.WaitGroup

// startWorker 启动一个 worker goroutine.
func startWorker(ctx context.Context, id int) {
	group.Add(1)
	go func() {
		defer group.Done()
		tick := time.Tick(time.Second)
	loop:
		for {
			select {
			case <-ctx.Done():
				// done 已关闭
				break loop
			case <-tick:
				doWork(id)
			}
		}
		// 执行善后
		doCleanup(id)
		// 正常退出
	}()
}

// startWorkerPrior23 启动一个 worker goroutine.
// 演示在 Go1.23 之前如何正确释放 ticker
func startWorkerPrior23(ctx context.Context, id int) {
	group.Add(1)
	go func() {
		defer group.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop() // 确保 ticker 会被终止
	loop:
		for {
			select {
			case <-ctx.Done():
				// done 已关闭
				break loop
			case <-ticker.C:
				doWork(id)
			}
		}
		// 执行善后
		doCleanup(id)
		// 正常退出
	}()
}

func doCleanup(id int) {
	fmt.Printf("#%v: cleaning up...\n", id)
}

func doWork(id int) {
	fmt.Printf("#%v working...\n", id)
}
