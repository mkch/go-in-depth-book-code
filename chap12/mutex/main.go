package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var group sync.WaitGroup
	group.Add(1)

	var stop bool
	var stopLock sync.Mutex // stopLock 保护 stop　变量
	// 在 stopLock 的保护下读取 stop 的值
	var readStop = func() bool {
		stopLock.Lock()
		defer stopLock.Unlock()
		return stop
	}

	// 在 stopLock 的保护下设置 stop 的值
	var writeStop = func(newValue bool) {
		stopLock.Lock()
		defer stopLock.Unlock()
		stop = newValue
	}

	// 启动 goroutine 1
	go func() {
		defer group.Done()
		for i := range 100 {
			if readStop() {
				return
			}
			// stopLock.Lock() // 加锁
			// if stop {       // 访问 stop
			// 	stopLock.Unlock() // 解锁
			// 	return
			// }
			stopLock.Unlock() // 解锁
			fmt.Println(i)
			time.Sleep(time.Millisecond * 20)
		}
	}()

	time.Sleep(time.Second * 1)

	// 设置 stop 为 true
	// stopLock.Lock()
	// stop = true
	// stopLock.Unlock()
	writeStop(true)

	// 等待 goroutine 1 正常结束
	group.Wait()
}

// Counter 是一个并发安全的计数器
type Counter struct {
	value uintptr
	lock  sync.RWMutex
}

// Inc 增加计数值
func (c *Counter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
}

// Value 返回计数值
func (c *Counter) Value() uintptr {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.value
}
