package main

import (
	"fmt"
	"sync"
)

func main() {
	var g sync.WaitGroup
	g.Add(2) // 共有两个并发需要等待

	go func() {
		fmt.Println(1)
		g.Done() // 完成一个
	}()
	go func() {
		fmt.Println(2)
		g.Done() // 完成一个
	}()

	g.Wait() // 等待两个并发都完成
	fmt.Println(3)
}
