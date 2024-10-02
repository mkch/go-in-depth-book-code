package gotcha

import (
	"sync"
	"testing"
)

func TestF(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	var value int

	change := func(n int) {
		cond.L.Lock()
		value = n
		cond.Signal()
		cond.L.Unlock()
	}

	go func() {
		for n := 1; ; n++ {
			change(n)
		}
	}()

	listen := func(old int) int {
		cond.L.Lock()
		defer cond.L.Unlock()
		for value == old {
			cond.Wait()
		}
		return value
	}

	for {
		var old int
		new := listen(old)
		if new-old != 1 {
			t.Fatalf("new=%v old=%v", new, old)
		}
	}
}
