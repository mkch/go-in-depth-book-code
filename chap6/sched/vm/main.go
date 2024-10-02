package main

import (
	"fmt"
	"time"
)

// Instruction 为 VM 中的一条指令.
//
// 此 VM 为一个无栈 VM, 其中函数无参数也无返回值.
// 一个函数被编译为若干条指令加一个 nil.
// 因为没有栈, 所以函数调用只能出现在最顶层 [NewG] 中.
type Instruction func(*Thread)

var Code []Instruction

// Thread 为 VM 中的一个本地线程.
type Thread struct {
	id int // 此线程的 ID
	pc int // 程序计数器, 指向 [Code] 中要执行的指令
}

// NewThread 启动一个 VM 本地线程.
// proc 为线程要执行的函数地址, 对应 Code 中的位置.
func NewThread(id int, proc int) {
	t := &Thread{id: id, pc: proc}
	// 在此 VM 中, 本地线程即 goroutine
	go func() {
		for {
			for inst := Code[t.pc]; inst != nil; inst = Code[t.pc] {
				t.pc++  // 指向下一条指令
				inst(t) // 执行当前指令
			}
		}
	}()
	return
}

func main() {
	Code = []Instruction{
		// 第一个函数 func1
		// 相当于如下等效 go 函数, 其中 tid 为本地线程 ID
		// func func1() {
		// 	fmt.Printf("Thread#%d func1 A\n", tid)
		// 	fmt.Printf("Thread#%d func1 B\n", tid)
		// 	fmt.Printf("Thread#%d func1 C\n", tid)
		// }
		/*0*/ func(t *Thread) { fmt.Printf("Thread#%d func1 A\n", t.id) },
		/*1*/ func(t *Thread) { fmt.Printf("Thread#%d func1 B\n", t.id) },
		/*2*/ func(t *Thread) { fmt.Printf("Thread#%d func1 C\n", t.id) },
		/*3*/ nil,

		// 第二个函数 func2
		// 相当于如下等效 go 函数, 其中 tid 为本地线程 ID
		// func func2() {
		// 	fmt.Printf("Thread#%d func2 D\n", tid)
		// 	fmt.Printf("Thread#%d func2 E\n", tid)
		// 	fmt.Printf("Thread#%d func2 F\n", tid)
		// 	fmt.Printf("Thread#%d func2 G\n", tid)
		// }
		/*4*/ func(t *Thread) { fmt.Printf("Thread#%d func2 D\n", t.id) },
		/*5*/ func(t *Thread) { fmt.Printf("Thread#%d func2 E\n", t.id) },
		/*6*/ func(t *Thread) { fmt.Printf("Thread#%d func2 F\n", t.id) },
		/*7*/ func(t *Thread) { fmt.Printf("Thread#%d func2 G\n", t.id) },
		/*8*/ nil,
	}

	// 使用函数 func1 (地址为 0) 启动一个线程
	NewThread(1, 0)
	// 使用函数 func2 (地址为 4) 启动一个线程
	NewThread(2, 4)

	time.Sleep(time.Second * 1)
}
