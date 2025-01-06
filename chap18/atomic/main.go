package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Counter 是一个并发安全的计数器。
type Counter struct {
	value atomic.Uintptr
}

// Inc 增加计数值。
func (c *Counter) Inc() {
	c.value.Add(1)
}

// Value 返回计数值。
func (c *Counter) Value() uintptr {
	return c.value.Load()
}

// LockFreeStack 是一个并发安全的栈。
type LockFreeStack struct {
	top atomic.Pointer[stackItem]
}

type stackItem struct {
	Value int
	Next  *stackItem
}

// Pop 弹出栈顶元素到 t。
func (s *LockFreeStack) Pop() (t int, ok bool) {
	for {
		top := s.top.Load()
		if top == nil {
			return 0, false
		}
		if s.top.CompareAndSwap(top, top.Next) {
			return top.Value, true
		}
	}
}

// Push 向栈顶压入元素 t。
func (s *LockFreeStack) Push(t int) {
	var newItem = stackItem{Value: t}
	for {
		top := s.top.Load()
		newItem.Next = top
		if s.top.CompareAndSwap(top, &newItem) {
			return
		}
	}
}

func main() {
	var s LockFreeStack
	var group sync.WaitGroup
	group.Add(3)
	go func() {
		defer group.Done()
		time.Sleep(time.Millisecond * 10)
		s.Push(1)
	}()
	go func() {
		defer group.Done()
		time.Sleep(time.Millisecond * 10)
		s.Push(2)
		s.Push(3)
	}()
	go func() {
		defer group.Done()
		time.Sleep(time.Millisecond * 10)
		fmt.Println(s.Pop())
		s.Push(4)
	}()

	group.Wait()
	for v, ok := s.Pop(); ok; v, ok = s.Pop() {
		fmt.Println(v)
	}
}
