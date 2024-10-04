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
func Select[T any](recvBranches []*Recv[T], sendBranches []*Send[T], defaultBranch Default) {
	// 持有所有分支的锁
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

	// 为所有分支添加 signalSelect. 调用时必须持有所有锁.
	var appendSignal = func(n *signalSelect) {
		for _, r := range recvBranches {
			if r.Chan != nil {
				r.Chan.signalSelect = append(r.Chan.signalSelect, n)
			}
		}
		for _, w := range sendBranches {
			if w.Chan != nil {
				w.Chan.signalSelect = append(w.Chan.signalSelect, n)
			}
		}
	}
	// 为所有分支删除 signalSelect. 调用时必须持有所有锁.
	var removeSignal = func(n *signalSelect) {
		for _, r := range recvBranches {
			if r.Chan != nil {
				r.Chan.signalSelect = slices.DeleteFunc(r.Chan.signalSelect, func(a *signalSelect) bool { return a == n })
			}
		}
		for _, w := range sendBranches {
			if w.Chan != nil {
				w.Chan.signalSelect = slices.DeleteFunc(w.Chan.signalSelect, func(a *signalSelect) bool { return a == n })
			}
		}
	}

	// 阻塞等待, 直到运行了一个不阻塞的分支
	var wait = func() {
		// caseReadyCond 保护 caseReady 的 Cond
		var caseReadyCond = sync.NewCond(&sync.Mutex{})
		var caseReady bool // 是否有分支解除阻塞
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
			r, w := randSelect(recvBranches, sendBranches)
			if r != nil {
				removeSignal(&signal)
				var ok bool
				v := r.Chan.recv(&ok)
				unlock()
				r.F(v, ok)
				return
			} else if w != nil {
				removeSignal(&signal)
				w.Chan.send(w.Value)
				unlock()
				w.F()
				return
			}
			unlock()
		}
	}

	lock()

	// 找出不阻塞的分支
	r, w := randSelect(recvBranches, sendBranches)
	if r != nil {
		var ok bool
		v := r.Chan.recv(&ok)
		unlock()
		r.F(v, ok)
	} else if w != nil {
		w.Chan.send(w.Value)
		unlock()
		w.F()
	} else {
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
}

// randSelect 从 recvCases 和 sendCases 中随机选择一个不会阻塞的分支.
// 调用时必须持有所有分支的锁.
func randSelect[T any](recvBranches []*Recv[T], sendBranches []*Send[T]) (r *Recv[T], w *Send[T]) {
	recvBranches = slices.DeleteFunc(slices.Clone(recvBranches),
		func(a *Recv[T]) bool { return a.Chan == nil || !a.Chan.canRecv() })
	sendBranches = slices.DeleteFunc(slices.Clone(sendBranches),
		func(a *Send[T]) bool { return a.Chan == nil || !a.Chan.canSend() })
	n := len(recvBranches) + len(sendBranches)
	if n == 0 {
		return
	}
	i := rand.IntN(n)
	if i < len(recvBranches) {
		r = recvBranches[i]
	} else {
		w = sendBranches[i-len(recvBranches)]
	}
	return
}
