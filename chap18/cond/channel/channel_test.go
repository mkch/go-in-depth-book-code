package channel

import (
	"context"
	"slices"
	"testing"
	"time"
)

func TestSendRecv(t *testing.T) {
	var test = func(cap int) {
		c := Make[int](cap)
		go c.Send(1)

		if v := c.Recv(nil); v != 1 {
			t.Fatal(v)
		}
	}
	t.Run("Buffered1", func(t *testing.T) { test(1) })
	t.Run("Buffered10", func(t *testing.T) { test(10) })
	t.Run("NoBuffer", func(t *testing.T) { test(0) })
}

func TestCloseOnRecv(t *testing.T) {
	var test = func(cap int) {
		c := Make[string](cap)
		go func() {
			time.Sleep(time.Millisecond * 100)
			c.Close()
		}()
		var ok bool
		if v := c.Recv(&ok); ok != false || v != "" {
			t.Fatal(v, ok)
		}
	}
	t.Run("Buffered1", func(t *testing.T) { test(1) })
	t.Run("Buffered10", func(t *testing.T) { test(10) })
	t.Run("NoBuffer", func(t *testing.T) { test(0) })
}

func TestClosed(t *testing.T) {
	var test = func(cap int) {
		c := Make[int](cap)
		c.Close()

		if v := c.Recv(nil); v != 0 {
			t.Fatal(v)
		}

		var ok bool
		if v := c.Recv(&ok); ok != false || v != 0 {
			t.Fatal(v, ok)
		}

		var panicked bool
		func() {
			defer func() {
				panicked = recover() != nil
			}()
			c.Send(1)
		}()
		if !panicked {
			t.Fatal(panicked)
		}

		panicked = false
		func() {
			defer func() {
				panicked = recover() != nil
			}()
			c.Close()
		}()
		if !panicked {
			t.Fatal(panicked)
		}
	}
	t.Run("Buffered1", func(t *testing.T) { test(1) })
	t.Run("Buffered10", func(t *testing.T) { test(10) })
	t.Run("NoBuffer", func(t *testing.T) { test(0) })
}

func TestNil(t *testing.T) {
	var test = func(op func(c *Channel[int])) {
		var c *Channel[int]
		var done = make(chan struct{})
		go func() {
			defer close(done)
			op(c)
		}()
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()
		select {
		case <-done:
			t.Fatal("should block")
		case <-ctx.Done():
		}
	}

	t.Run("Send", func(t *testing.T) {
		test(func(c *Channel[int]) { c.Send(1) })
	})
	t.Run("Recv", func(t *testing.T) {
		test(func(c *Channel[int]) { c.Recv(nil) })
	})
	t.Run("RecvOk", func(t *testing.T) {
		test(func(c *Channel[int]) { c.Recv(new(bool)) })
	})
}

func TestConcurrency(t *testing.T) {
	var test = func(cap int) {
		c := Make[int](cap)
		result := make(chan int)
		recv := func(id int) {
			_ = id
			var ok bool
			for v := c.Recv(&ok); ok; v = c.Recv(&ok) {
				//t.Logf("#%v Received %v", id, v)
				result <- v
			}
		}

		go recv(1)
		go recv(2)
		go recv(3)

		go func() {
			for i := range 10 {
				//t.Logf("Sending %v", i)
				c.Send(i)
			}
			c.Close()
		}()

		var s []int
		for range 10 {
			s = append(s, <-result)
		}
		slices.Sort(s)
		if !slices.Equal(s, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Fatal(s)
		}
	}
	t.Run("Buffered1", func(t *testing.T) { test(1) })
	t.Run("Buffered100", func(t *testing.T) { test(100) })
	t.Run("NoBuffer", func(t *testing.T) { test(0) })
}
