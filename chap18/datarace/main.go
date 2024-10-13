package main

import (
	"fmt"
	"math"
	"sync"
)

func F1_Fixed() {
	var total int

	var group sync.WaitGroup
	group.Add(1)

	var m sync.Mutex

	go func() {
		for range 100 {
			m.Lock()
			total += 1
			m.Unlock()
		}
		group.Done()
	}()

	for range 100 {
		m.Lock()
		total -= 1
		m.Unlock()
	}

	group.Wait()
}

func F1() {
	var total int

	go func() {
		for range 100 {
			total += 1
		}
	}()

	for range 100 {
		total -= 1
	}

	// busy loop, 仅用作演示, 不推荐
	for total != 0 {
	}
}

func main() {
	for i := range math.MaxInt {
		fmt.Println(i)
		F1()
	}
}
