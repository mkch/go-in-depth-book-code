package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var group sync.WaitGroup
	group.Add(1)

	// stop 是一个并发安全的信号
	var stop = make(chan struct{})

	// 启动 goroutine 1
	go func() {
		defer group.Done()
		for i := range 100 {
			select {
			case <-stop:
				// 收到退出信号, 退出
				return
			default:
				// 无需退出
			}
			fmt.Println(i)
			time.Sleep(time.Millisecond * 20)
		}
	}()

	time.Sleep(time.Second * 1)
	close(stop) // 发出退出信号

	// 等待 goroutine 1 正常结束
	group.Wait()
}
