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
	var group sync.WaitGroup

	var receiveEvent = func(n int) {
		defer group.Done()
		for {
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("goroutine %v received %#v\n", n, event)
			if event == "exit" {
				cond.L.Unlock()
				return
			}
			cond.L.Unlock()
		}
	}

	group.Add(2)
	go receiveEvent(1)
	go receiveEvent(2)

	time.Sleep(time.Second)

	lock.Lock()
	event = "broadcast event"
	lock.Unlock()
	cond.Broadcast() // 唤醒所有 Wait

	time.Sleep(time.Second)

	lock.Lock()
	event = "signal event"
	lock.Unlock()
	cond.Signal() // 唤醒一个 Wait

	time.Sleep(time.Second)

	lock.Lock()
	event = "exit"
	lock.Unlock()
	cond.Broadcast()

	group.Wait()
}
