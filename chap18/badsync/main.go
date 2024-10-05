package main

import (
	"fmt"
	"math"
	"sync"
)

func F1() {
	var total int
	var group sync.WaitGroup
	group.Add(1)

	go func() {
		total += 1
		group.Done()
	}()

	total -= 1

	group.Wait()
	if total != 0 {
		panic(total)
	}
}

func main() {
	var i int
	defer func() {
		fmt.Printf("loop count: %v\n", i+1)
	}()
	for i = range math.MaxInt {
		F1()
	}
}
