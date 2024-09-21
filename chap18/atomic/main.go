package main

import (
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

// Counter 是一个并发安全的计数器
type Counter struct {
	value atomic.Uintptr
}

// Inc 增加计数值
func (c *Counter) Inc() {
	c.value.Add(1)
}

// Value 返回计数值
func (c *Counter) Value() uintptr {
	return c.value.Load()
}

// LockFreeSlice 是一个无锁的 []int
type LockFreeSlice struct {
	p atomic.Pointer[[]int]
}

// NewLockFreeSlice 创建一个 LockFreeSlice, 初始内容为 s
func NewLockFreeSlice(s []int) *LockFreeSlice {
	ret := &LockFreeSlice{}
	ret.p.Store(&s)
	return ret
}

// Value 返回 l 的值
func (l *LockFreeSlice) Value() []int {
	return slices.Clone(*l.p.Load())
}

// Append 类似内建函数 append, 但是可以在并发环境中使用
func (l *LockFreeSlice) Append(s []int) {
	for {
		old := l.p.Load()
		new := append((*old)[:len(*old):len(*old)], s...)
		if l.p.CompareAndSwap(old, &new) {
			return
		}
	}
}

func main() {
	var s = NewLockFreeSlice([]int{1})
	var group sync.WaitGroup
	group.Add(3)
	go func() {
		defer group.Done()
		time.Sleep(time.Millisecond * 10)
		s.Append([]int{2})
		fmt.Println(s.Value())
	}()
	go func() {
		defer group.Done()
		time.Sleep(time.Millisecond * 10)
		s.Append([]int{3, 4})
		fmt.Println(s.Value())
	}()
	go func() {
		defer group.Done()
		time.Sleep(time.Millisecond * 10)
		s.Append([]int{5, 6, 7})
		fmt.Println(s.Value())
	}()

	group.Wait()
	fmt.Println(s.Value())
}
