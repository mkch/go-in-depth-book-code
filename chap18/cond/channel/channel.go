package channel

import (
	"sync"
)

type Channel[T any] struct {
	cond         sync.Cond // 条件变量
	buf          []T       // 缓冲区
	cap          int       // 容量
	closed       bool      // 是否已关闭
	readBlocking bool      // 是否有阻塞的读, 在 cap == 0 时 canSend 用到

	// 如果一个 Select 需要等待此 Channel, 会向此切片中添加一个 signalSelect.
	// 当条件改变时, 会调用 signalSelect 通知 Select.
	signalSelect []*signalSelect
}

// Make 创建 Channel. 相当于 make(chan T).
func Make[T any](capacity int) *Channel[T] {
	if capacity < 0 {
		panic("invalid argument: cap must not be negative")
	}
	c := &Channel[T]{cap: capacity}
	c.cond.L = &sync.Mutex{}
	return c
}

// Close 关闭 c. 相当于 close(c).
func (c *Channel[T]) Close() {
	// Close nil Channel 应该 panic
	if c == nil {
		panic("close nil Channel")
	}
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	// Close 已关闭的 Channel 应该 panic
	if c.closed {
		panic("already closed")
	}
	c.closed = true
	c.cond.Broadcast()

	// 通知正在等待读的 Select
	c.callSignalSelect()
}

// locked 为一个已上锁的 *sync.Mutext
var locked sync.Mutex

func init() {
	locked.Lock()
}

// blockForever 永远阻塞
// 用于实现读写 nil Channel 时的行为.
func blockForever() {
	locked.Lock()
}

// Send 向 c 中写入一个值 v
// 相当于 c <- v
func (c *Channel[T]) Send(v T) {
	// 写 nil Channel 永远阻塞
	if c == nil {
		blockForever()
	}
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	c.send(v)
}

// callSignalSelect 调用 signalSelect,
// 必须持有锁.
func (c *Channel[T]) callSignalSelect() {
	for _, n := range c.signalSelect {
		n.f()
	}
}

// send 向 c 中写入 v. 调用时必须持有 c.cond.L.
func (c *Channel[T]) send(v T) {
	var sendOnClosed = func() { panic("send on closed channel") }
	// 等待 buf 有剩余空间
	// 特例: cap == 0 时, len(buf) 最大为 1
	for !(c.closed || len(c.buf) < max(1, c.cap)) {
		c.cond.Wait()
	}
	if c.closed {
		sendOnClosed()
	}
	// 写出 v
	c.buf = append(c.buf, v)
	if c.cap == 0 {
		c.readBlocking = false
	}
	// 通知
	c.cond.Broadcast()
	// 通知 Select
	c.callSignalSelect()
	// 处理 cap == 0 的特殊情况
	if c.cap == 0 {
		// 当 cap == 0, 必须等读方操作完成
		for !(c.closed || len(c.buf) == 0) {
			c.cond.Wait()
		}
		if c.closed {
			sendOnClosed()
		}
	}
}

// Recv 从 c 中读取一个值.
// 如果 ok 不为 nil, 相当于 v, *ok = <- c.
// 如果 ok 为 nil, 相当于 v = <- c.
func (c *Channel[T]) Recv(ok *bool) (v T) {
	// 写 nil Channel 永远阻塞
	if c == nil {
		blockForever()
	}
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	// 等待被关闭, 或就绪
	return c.recv(ok)
}

// recv 从 c 中读取一个值. 调用时必须持有 c.cond.L.
func (c *Channel[T]) recv(ok *bool) (v T) {
	var setOk = func(val bool) {
		if ok != nil {
			*ok = val
		}
	}
	// 等待被关闭或缓冲区有数据可读
	for !(c.closed || len(c.buf) > 0) {
		if c.cap == 0 {
			c.readBlocking = true
			// 通知 Select
			c.callSignalSelect()
		}
		c.cond.Wait()
	}
	// 已关闭, 且缓冲为空
	if c.closed && len(c.buf) == 0 {
		setOk(false)
		return
	}
	// 读取
	v = c.buf[0]
	c.buf = c.buf[1:]
	// 通知
	c.cond.Broadcast()
	// 通知 Select
	c.callSignalSelect()
	setOk(true)
	return
}

// canRead 返回 c 当前是否可不阻塞地执行一个 Send().
// 调用时必须持有 c.L.
func (c *Channel[T]) canSend() bool {
	return c != nil && // nil Channel 永远阻塞
		(c.closed || // 已关闭也可以不阻塞地发送(panic)
			len(c.buf) < c.cap || // 有可用容量
			(c.cap == 0 && c.readBlocking)) // 容量为 0, 但有阻塞读
}

// canRead 返回 c 当前是否可不阻塞地执行一个 Recv().
// 调用时必须持有 c.L.
func (c *Channel[T]) canRecv() bool {
	return c != nil && (c.closed || len(c.buf) > 0)
}
