package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var value int

	change := func(n int) {
		cond.L.Lock()
		defer cond.L.Unlock()
		value = n
		cond.Signal()
	}

	listen := func(old int) int {
		cond.L.Lock()
		defer cond.L.Unlock()
		for value == old {
			cond.Wait()
		}
		return value
	}

	go func() {
		var old int
		for range 10 {
			new := listen(old)
			fmt.Println(new)
			old = new
		}
	}()

	for n := range 10 {
		change(n + 1)
	}

	time.Sleep(time.Second * 2)
}
