package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startWorker(0)
	startWorker(1)
	time.Sleep(time.Second * 3)
	shutdown()
	fmt.Println("exit")
}

var done = make(chan struct{})
var group sync.WaitGroup

// startWorker 启动一个 worker goroutine
func startWorker(id int) {
	group.Add(1)
	go func() {
		defer group.Done()
	loop:
		for {
			select {
			case <-done:
				// done 已关闭
				break loop
			default:
				// NOP
			}
			doWork(id)
			time.Sleep(time.Second)
		}
		// 执行善后
		doCleanup(id)
		// 正常退出
	}()
}

// shutdown 关闭所有 worker 并等待其退出
func shutdown() {
	close(done)  // 发出退出信号
	group.Wait() // 等待 worker 退出
}

func doCleanup(id int) {
	fmt.Printf("#%v: cleaning up...\n", id)
}

func doWork(id int) {
	fmt.Printf("#%v working...\n", id)
}
