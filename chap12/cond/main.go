package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var event string
	var lock sync.RWMutex
	var cond = sync.Cond{L: lock.RLocker()} // 接收者只读

	var receiveEvent = func(n int) {
		cond.L.Lock()
		defer cond.L.Unlock()
		for event == "" {
			cond.Wait()
		}
		fmt.Printf("goroutine %v received %#v\n", n, event)
	}

	go receiveEvent(1)
	go receiveEvent(2)

	time.Sleep(time.Second)

	lock.Lock()
	event = "event 1"
	lock.Unlock()
	cond.Broadcast() // 唤醒所有 Wait

	time.Sleep(time.Second * 2)
}
