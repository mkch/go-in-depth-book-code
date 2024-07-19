package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {
	selectDemo()
	gen, cancel := NewIdGenerator()
	defer cancel()
	fmt.Println(gen())
	valBackupDefault()
	fmt.Println(readTimeout(nil))
}

func selectDemo() {
	c1 := make(chan int, 1)
	c2 := make(chan int)
	close(c2)

	select {
	case v := <-c1: // 阻塞
		fmt.Println("read c1", v)
	case c1 <- 1: // 不阻塞
		fmt.Println("write c1")
	case v, ok := <-c2: // 不阻塞
		fmt.Println("read c2", v, ok)
	default:
		fmt.Println("default")
	}
}

// NewIdGenerator 返回一个并发安全的顺序 ID 生成器 g,
// 和一个可以停止 ID 生成的函数 cancel
func NewIdGenerator() (g func() uint64, cancel func()) {
	c := make(chan uint64)
	c1 := make(chan struct{}) // 1
	producer := func(w chan<- uint64, exit <-chan struct{}) {
		for i := range uint64(math.MaxUint64) {
			select { // 2
			case w <- i:
			case <-exit:
				return
			}
		}
		panic("out of Id")
	}
	go producer(c, c1)
	g = func() uint64 { return <-c }
	cancel = func() {
		close(c1) // 3
	}
	return
}

// valBackupDefault 演示了 val, backup, def 三级 取值.
func valBackupDefault() {
	val := make(chan int, 1)
	backup := make(chan int, 1)
	// 把 val 和 backup 交给其他 goroutine 去生成值
	// ...

	def := -1
	var v int
	// 从 val 或 backup 读取值到 v
	// 如果此时二者均无值, 则使用 def
	select {
	case v = <-val:
	case v = <-backup:
	default:
		v = def
	}
	// 使用 v
	// ...
	fmt.Println(v)

	// 保证 val 优先级大于 backup
	select {
	case v = <-val:
	default:
		select {
		case v = <-backup:
		default:
			v = def
		}
	}
	fmt.Println(v)
}

// readTimeout 从 c 中读取一个值, 如果阻塞时间过长则返回非 nil error.
func readTimeout(c <-chan int) (int, error) {
	// timer 控制读取的超时
	timer := time.NewTimer(time.Millisecond * 20)
	defer timer.Stop() // 防止资源泄露
	select {
	case v := <-c:
		return v, nil
	case <-timer.C:
		return 0, errors.New("timed out")
	}
}

// sendWithDrop 试图把 item 写入 ch 中,
// 如果此写入会阻塞, 则抛弃该值
func sendWithDrop(c chan<- int, item int) {
	select {
	case c <- item:
		// 发送成功
	default:
		// 丢弃 item
	}
}

// sendWithDrop 试图把 item 写入 ch 中,
// 如果此写入会阻塞, 则抛弃已缓冲的值而写入 item.
// 要求 c 必须为一个容量为 1 的 channel
func sendKeepLatest(c chan int, item int) {
	select {
	case <-c:
		// 如果缓冲中有未读的值, 则先消耗掉
	default:
		// 如果缓冲中没有未读值, Nop
	}
	// 发送最新值
	c <- item
}
