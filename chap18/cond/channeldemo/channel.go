package channel

import (
	"sync"
)

type Channel[T any] struct {
	cond sync.Cond // 条件变量
	buf  []T       // 缓冲区
	cap  int       // 容量
}

// Make 创建 Channel. 相当于 make(chan T).
func Make[T any](capacity int) *Channel[T] {
	if capacity <= 0 {
		panic("invalid argument: cap must be positive")
	}
	c := &Channel[T]{cap: capacity}
	c.cond.L = &sync.Mutex{}
	return c
}

// Send 向 c 中写入一个值 v.
// 相当于 c <- v
func (c *Channel[T]) Send(v T) {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	c.send(v)
}

// send 向 c 中写入 v. 调用时必须持有 c.cond.L.
func (c *Channel[T]) send(v T) {
	// 等待 buf 有剩余空间
	for !(len(c.buf) < c.cap) {
		c.cond.Wait()
	}
	// 写出 v
	c.buf = append(c.buf, v)
	// 通知
	c.cond.Broadcast()
}

// Recv 从 c 中读取一个值
// 相当于 v = <- c
func (c *Channel[T]) Recv() (v T) {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	return c.recv()
}

// recv 从 c 中读取一个值. 调用时必须持有 c.cond.L.
func (c *Channel[T]) recv() (v T) {
	// 缓冲区有数据可读
	for !(len(c.buf) > 0) {
		c.cond.Wait()
	}
	// 读取
	v = c.buf[0]
	c.buf = c.buf[1:]
	// 通知
	c.cond.Broadcast()
	return
}
