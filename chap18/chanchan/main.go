package main

import (
	"fmt"
	"sync"
)

func main() {
	var calc = NewSumCalc()
	defer calc.Close()

	var group sync.WaitGroup
	group.Add(3)
	go func() {
		defer group.Done()
		fmt.Println(calc.Sum([]int{1, 2, 3, 4, 5}))
	}()
	go func() {
		defer group.Done()
		fmt.Println(calc.Sum([]int{5, 4, 3, 2, 1}))
	}()
	go func() {
		defer group.Done()
		fmt.Println(calc.Sum([]int{50, 40, 30, 20, 10}))
	}()

	group.Wait()
}

// SumCalc 是一个计算器
type SumCalc struct {
	taskChan chan *task
	group    sync.WaitGroup
}

// NewSumCalc 创建一个新的 SumCalc
// 使用完毕后必须调用 Close() 方法关闭
func NewSumCalc() (calc *SumCalc) {
	calc = &SumCalc{taskChan: make(chan *task)}
	const MAX_WORKERS = 3
	calc.group.Add(MAX_WORKERS)
	for range MAX_WORKERS {
		go worker(&calc.group, calc.taskChan)
	}
	return
}

// Close 关闭 sum, 释放其所占用的资源
func (sum *SumCalc) Close() {
	close(sum.taskChan)
	sum.group.Wait()
}

// Sum 计算 s 中所有元素的和
func (sum *SumCalc) Sum(s []int) int {
	var sumTask = &task{
		param:  s,
		result: make(chan int),
	}

	sum.taskChan <- sumTask // 向 worker goroutine 提交任务
	return <-sumTask.result // 从 worker goroutine 接收任务执行结果
}

// task 是提交给 worker 的一个任务
type task struct {
	param  []int    // 欲计算其和的 slice
	result chan int // 计算的结果写入此 chan
}

// worker 不断从 taskChan 读取 task, 计算 task.s 的和并把
// 结果写入 task.result
// worker 在返回前会调用 group.Done()
func worker(group *sync.WaitGroup, taskChan <-chan *task) {
	defer group.Done()
	for {
		if t, ok := <-taskChan; ok {
			// 读到任务
			// 执行任务(求和)
			sum := sumImpl(t.param)
			// 发送结果
			t.result <- sum
		} else {
			return // taskChan 被关闭了, 退出
		}
	}
}

// sumImpl 使用循环累加的方式计算 s 中所有元素的和.
func sumImpl(s []int) (ret int) {
	for _, n := range s {
		ret += n
	}
	return
}
