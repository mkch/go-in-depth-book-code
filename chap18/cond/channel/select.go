package channel

import (
	"math/rand/v2"
	"slices"
	"sync"
)

// Recv 为 Select 中一个接收分支.
type Recv[T any] struct {
	Chan *Channel[T]        // 操作对象
	F    func(v T, ok bool) // 读取完成后的操作
}

// Send 为 Select 中一个发送分支.
type Send[T any] struct {
	Chan  *Channel[T] // 操作对象
	Value T           // 要写入的值
	F     func()      // 发送成功后的操作
}

// Default 为 Select 中的一个 default 分支.
type Default func()

type signalSelect struct {
	f func()
}

// Select 对应 select 语句:
//
//	select {
//		case v, ok := r[N].Chan:
//			r[N].F(v, ok)
//		case w[N].Chan <- w[N].Value:
//			w[N].F()
//		default:
//			def()
//	}
//
// 如果任何分支缺失, 则无对应 case.
func Select[T any](recvBranches []*Recv[T],
	sendBranches []*Send[T],
	defaultBranch Default) {
	// 持有所有分支的锁
	// !! DEADLOCK !!
	var lock = func() {
		for _, r := range recvBranches {
			if r.Chan != nil {
				r.Chan.cond.L.Lock()
			}
		}
		for _, w := range sendBranches {
			if w.Chan != nil {
				w.Chan.cond.L.Lock()
			}
		}
	}
	// 释放所有分支的锁
	var unlock = func() {
		for _, r := range recvBranches {
			if r.Chan != nil {
				r.Chan.cond.L.Unlock()
			}
		}
		for _, w := range sendBranches {
			if w.Chan != nil {
				w.Chan.cond.L.Unlock()
			}
		}
	}

	// 为所有分支添加 signalSelect. 调用时必须持有所有分支的锁.
	var appendSignal = func(signal *signalSelect) {
		for _, r := range recvBranches {
			if r.Chan != nil {
				r.Chan.signalSelect = append(r.Chan.signalSelect,
					signal)
			}
		}
		for _, w := range sendBranches {
			if w.Chan != nil {
				w.Chan.signalSelect = append(w.Chan.signalSelect,
					signal)
			}
		}
	}
	// 为所有分支删除 signalSelect. 调用时必须持有所有分支的锁.
	var removeSignal = func(n *signalSelect) {
		for _, r := range recvBranches {
			if r.Chan != nil {
				r.Chan.signalSelect = slices.DeleteFunc(
					r.Chan.signalSelect,
					func(a *signalSelect) bool { return a == n })
			}
		}
		for _, w := range sendBranches {
			if w.Chan != nil {
				w.Chan.signalSelect = slices.DeleteFunc(
					w.Chan.signalSelect,
					func(a *signalSelect) bool { return a == n })
			}
		}
	}

	// execRandomNb 从 recvBranches 和 sendBranches 中
	// 随机选择一个不会阻塞的分支并执行.
	// 如果成功执行了一个分支, 返回 true, 否则返回 false.
	// 调用时必须持有所有分支的锁.
	// 如果有分支可执行, 则先执行 callback, 然后执行分支操作,
	// 并在调用分支的F()前调用 unlock().
	var execRandomNb = func(recvBranches []*Recv[T],
		sendBranches []*Send[T],
		callback func()) (exec bool) {
		recvBranches = slices.DeleteFunc(slices.Clone(recvBranches),
			func(a *Recv[T]) bool {
				return a.Chan == nil || !a.Chan.canRecv()
			})
		sendBranches = slices.DeleteFunc(slices.Clone(sendBranches),
			func(a *Send[T]) bool {
				return a.Chan == nil || !a.Chan.canSend()
			})
		n := len(recvBranches) + len(sendBranches)
		if n == 0 {
			return false
		}
		i := rand.IntN(n)
		if i < len(recvBranches) {
			callback()
			var ok bool
			b := recvBranches[i]
			v := b.Chan.recv(&ok)
			unlock()
			b.F(v, ok)
		} else {
			callback()
			b := sendBranches[i-len(recvBranches)]
			b.Chan.send(b.Value)
			unlock()
			b.F()
		}
		return true
	}

	// 阻塞等待, 直到运行了一个不阻塞的分支.
	// 调用时必须持有所有分支的锁.
	var wait = func() {
		var caseReadyCond = sync.NewCond(&sync.Mutex{})
		var caseReady bool // 是否有分支发生变化
		var signal = signalSelect{
			func() {
				caseReadyCond.L.Lock()
				defer caseReadyCond.L.Unlock()
				caseReady = true
				caseReadyCond.Signal()
			},
		}
		// 为所有分支添加 signal
		appendSignal(&signal)
		unlock()

		for {
			// 等待, 直到 caseReady 被通知
			caseReadyCond.L.Lock()
			for !caseReady {
				caseReadyCond.Wait()
			}
			caseReadyCond.L.Unlock()

			lock()
			// 执行一个不阻塞的分支
			if execRandomNb(recvBranches, sendBranches,
				func() { removeSignal(&signal) }) {
				return
			}
			unlock()
			// 如果没有执行成功(和其他并发竞争识别)则继续等待
		}
	}

	lock()

	// 执行一个不阻塞的分支
	if execRandomNb(recvBranches, sendBranches, func() {}) {
		return
	}
	// 所有分支阻塞
	if defaultBranch != nil {
		// 有 Default, 执行 Default 分支
		unlock()
		defaultBranch()
		return
	}
	// 阻塞等待
	wait()
}
