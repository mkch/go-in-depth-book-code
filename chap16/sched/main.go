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

// Thread 为 VM 中的一个本地线程.
type Thread struct {
	id int // 此线程的 ID
	g  *G  // 在此线程上运行的 G
	pc int // 程序计数器, 指向 [Code] 中要执行的指令
}

// NewThread 启动一个 VM 本地线程.
func NewThread(id int) {
	t := &Thread{id: id}
	// 在此 VM 中, 本地线程即 goroutine
	go func() {
		for {
			t.g = <-gq    // 从 gq 中取出一个待执行的 G
			t.pc = t.g.pc // 使用该 G 的 pc 继续执行
			for inst := Code[t.pc]; inst != nil; inst = Code[t.pc] {
				t.pc++  // 指向下一条指令
				inst(t) // 执行当前指令
			}
			// 当前 G 执行完毕, 此 Thread 空闲, 释放对 G 的引用
			t.g = nil
		}
	}()
	return
}

// G 为 VM 中的一个并发单元.
type G struct {
	pc int // 指向下一条要指向的指令
}

// gq 为等待执行的 G 队列.
var gq = make(chan *G, 1024)

// NewG 创建一个并发单元.
func NewG(pc int) {
	gq <- &G{pc: pc}
}

// Schedule 为调度指令, 此指令执行一次 G 调度.
// 本函数返回后, m 所关联的 G 会让出执行时间, 使得 gq 中的 G 有机会执行.
func Schedule(t *Thread) {
	select {
	case g := <-gq:
		out := t.g    // out 为被调度出去的 G
		out.pc = t.pc // 记录 out 的 pc 寄存器
		// 准备运行新 G
		t.g = g
		t.pc = t.g.pc // 把新 G 的 pc 作为当前 pc 寄存器
		// 把 out 加入等待队列
		gq <- out
	default:
		// 没有待执行的 G
	}
}

// Code 为 VM 的代码段.
var Code []Instruction

func main() {
	Code = []Instruction{
		// 第一个函数 func1
		/*0*/ func(t *Thread) { fmt.Printf("Thread#%d func1 A\n", t.id) },
		/*1*/ Schedule, // 编译器自动添加的调度指令
		/*2*/ func(t *Thread) { fmt.Printf("Thread#%d func1 B\n", t.id) },
		/*3*/ func(t *Thread) { fmt.Printf("Thread#%d func1 C\n", t.id) },
		/*4*/ nil,
		// 第二个函数 func2
		/*5*/ func(t *Thread) { fmt.Printf("Thread#%d func2 D\n", t.id) },
		/*6*/ func(t *Thread) { fmt.Printf("Thread#%d func2 E\n", t.id) },
		/*7*/ func(t *Thread) { fmt.Printf("Thread#%d func2 F\n", t.id) },
		/*8*/ Schedule, // 编译器自动添加的调度指令
		/*9*/ func(t *Thread) { fmt.Printf("Thread#%d func2 G\n", t.id) },
		/*10*/ nil,
		// 第三个函数 func3
		/*11*/ func(t *Thread) { fmt.Printf("Thread#%d func3 H\n", t.id) },
		/*12*/ Schedule, // 编译器自动添加的调度指令
		/*13*/ func(t *Thread) { fmt.Printf("Thread#%d func3 I\n", t.id) },
		/*14*/ nil,
	}

	NewThread(1)
	NewThread(2)

	// NewG(0) 代表使用函数 func1 (地址为 0) 启动一个 G
	// 此函数包含 5 条指令 (包括结尾的nil)
	NewG(0)
	// NewG(5) 代表使用函数 func2 (地址为 5), 启动一个 G
	// 此函数包含 6 条指令 (包括结尾的nil)
	NewG(5)
	// NewG(11) 代表使用函数 func3 (地址为 11), 启动一个 G
	// 此函数包含 3 条指令 (包括结尾的nil)
	NewG(11)

	time.Sleep(time.Second * 1)
}
