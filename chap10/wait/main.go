package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup()
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Sum(s))
	fmt.Println(Sum2(s))
}

func waitGroup() {
	var group sync.WaitGroup

	group.Add(1)
	go func() {
		fmt.Println("goroutine 1")
		group.Done()
	}()

	group.Add(1)
	go func() {
		fmt.Println("goroutine 2")
		group.Done()
	}()

	group.Wait()
}

// Sum 返回参数 s 中所有元素的和
func Sum(s []int) int {
	if len(s) < 2 {
		return sumImpl(s)
	}
	var sum1 int // 前一半的和
	var sum2 int // 后一半的和
	var group sync.WaitGroup
	group.Add(2)
	go func() {
		sum1 = sumImpl(s[:len(s)/2])
		group.Done()
	}()
	go func() {
		sum2 = sumImpl(s[len(s)/2:])
		group.Done()
	}()
	group.Wait() // 等待两半都计算完毕
	return sum1 + sum2
}

// Sum 返回参数 s 中所有元素的和
func Sum2(s []int) int {
	if len(s) < 2 {
		return sumImpl(s)
	}
	var sum = make(chan int) // 1
	//var sum = make(chan int, 2)
	go func() {
		sum <- sumImpl(s[:len(s)/2]) // 把前一半的计算结果写入 sum
	}()
	go func() {
		sum <- sumImpl(s[len(s)/2:]) // 把后一半的计算结果写入 sum
	}()

	return <-sum + <-sum // 依次读出两个值并相加
}

// sumImpl 使用循环累加的方式计算 s 中所有元素的和.
func sumImpl(s []int) (ret int) {
	for _, n := range s {
		ret += n
	}
	return
}
