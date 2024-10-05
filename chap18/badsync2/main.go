package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

func F2() {
	var str string
	var group sync.WaitGroup
	group.Add(1)

	go func() {
		str = "1"
		group.Done()
	}()

	strconv.Atoi(str)
	group.Wait()
}

func main() {
	var i int
	defer func() {
		fmt.Printf("loop count: %v\n", i+1)
	}()
	for i = range math.MaxInt {
		F2()
	}
}
