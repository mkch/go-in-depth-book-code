package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	Chan()
	Ctx()
}

func Chan() {
	// cancellation 是一个取消信号，关闭此 chan 表示取消任务
	cancellation := make(chan struct{})
	// 500 毫秒后取消任务
	time.AfterFunc(time.Millisecond*500, func() { close(cancellation) })
	// 开始任务
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("任务开始")
		for i := range 10 { // 假设任务需要分 10 步进行
			select {
			case <-cancellation: // chan 被关闭，任务取消
				fmt.Println("任务取消")
				return
			default:
			}
			// 模拟每步操作
			fmt.Printf("第 %v 步...\n", i)
			time.Sleep(time.Millisecond * 150)
		}
	}()

	wg.Wait()
	fmt.Println("程序结束")
}

func Ctx() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	// 开始任务
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("任务开始")
		for i := range 10 { // 假设任务需要分 10 步进行
			select {
			case <-ctx.Done(): // 任务取消
				fmt.Println("任务取消")
				return
			default:
			}
			// 模拟每步操作
			fmt.Printf("第 %v 步...\n", i)
			time.Sleep(time.Millisecond * 150)
		}
	}()

	wg.Wait()
	fmt.Println("程序结束")
}
