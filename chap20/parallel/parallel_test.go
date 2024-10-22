package parallel

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	t.Parallel()
	for i := range 4 {
		time.Sleep(time.Millisecond * 20)
		t.Log(i)
	}
}
func Test11(t *testing.T) {
	t.Parallel()
	t.Run("A1", func(t *testing.T) {
		t.Parallel()
		for i := range 2 {
			time.Sleep(time.Millisecond * 20)
			t.Log(i)
		}
	})
	t.Run("A11", func(t *testing.T) {
		t.Parallel()
		for i := 2; i < 4; i++ {
			time.Sleep(time.Millisecond * 20)
			t.Log(i)
		}
	})
	t.Run("B", func(t *testing.T) {
		for i := range 2 {
			time.Sleep(time.Millisecond * 20)
			t.Log(i)
		}
	})
}

func Test2(t *testing.T) {
	time.Sleep(time.Millisecond * 200)
	for i := range 4 {
		time.Sleep(time.Millisecond * 20)
		t.Log(i)
	}
}

func Benchmark1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			time.Sleep(time.Nanosecond * 50000)
		}
	})
}
