package channel

import (
	"testing"
	"time"
)

func TestSelectCase(t *testing.T) {
	c1 := Make[int](0)
	c2 := Make[int](1)

	var (
		recv, send bool
	)

	go c1.Send(1)
	go c2.Recv(nil)

	Select(
		[]*Recv[int]{{Chan: c1, F: func(v int, ok bool) {
			recv = true
			if v != 1 && ok != true {
				t.Fatal(v, ok)
			}
		}}},
		[]*Send[int]{{Chan: c2, Value: 1, F: func() {
			send = true
		}}},
		nil,
	)

	if !recv && !send {
		t.Fatal("no case selected")
	}
}

func TestSelectWaitSend(t *testing.T) {
	c2 := Make[int](0)

	var (
		recv, send bool
	)

	go func() {
		time.Sleep(time.Millisecond * 100)
		c2.Recv(nil)
	}()

	Select(
		[]*Recv[int]{{Chan: nil, F: func(v int, ok bool) {
			recv = true
		}}},
		[]*Send[int]{{Chan: c2, Value: 1, F: func() {
			send = true
		}}},
		nil,
	)

	if recv || !send {
		t.Fatal("wrong branch")
	}
}

func TestSelectWaitRecv(t *testing.T) {
	c2 := Make[int](0)

	var (
		recv, send bool
	)

	go func() {
		time.Sleep(time.Millisecond * 100)
		c2.Send(1)
	}()

	Select(
		[]*Recv[int]{{Chan: c2, F: func(v int, ok bool) {
			recv = true
		}}},
		[]*Send[int]{{Chan: nil, Value: 1, F: func() {
			send = true
		}}},
		nil,
	)

	if !recv || send {
		t.Fatal("wrong branch")
	}
}

func TestSelectDefault(t *testing.T) {
	var (
		def, recv bool
	)
	Select(
		[]*Recv[int]{{Chan: nil, F: func(v int, ok bool) {
			recv = true
		}}},
		nil,
		func() { def = true },
	)
	if !def || recv {
		t.Fatal("wrong branch")
	}
}

func TestRandBranch(t *testing.T) {
	c1 := Make[int](0)
	c2 := Make[int](2)
	c3 := Make[int](0)
	c4 := Make[int](10)

	var result = make(map[int]bool)

	for range 10 {
		go c1.Send(1)
		go c2.Send(2)
		go func() {
			var ok bool
			if v := c3.Recv(&ok); !ok || v != 3 {
				panic(v)
			}
		}()
		go func() {
			var ok bool
			if v := c4.Recv(&ok); !ok || v != 4 {
				panic(v)
			}
		}()

		var selected int = -1

		Select(
			[]*Recv[int]{
				{Chan: c1, F: func(v int, ok bool) {
					if v != 1 || !ok {
						t.Fatal(v, ok)
					}
					selected = 1
				}},
				{Chan: c2, F: func(v int, ok bool) {
					if v != 2 || !ok {
						t.Fatal(v, ok)
					}
					selected = 2
				}},
			},
			[]*Send[int]{
				{Chan: c3, Value: 3, F: func() { selected = 3 }},
				{Chan: c4, Value: 4, F: func() { selected = 4 }},
			},
			nil)

		result[selected] = true

		for i, f := range []func(){
			func() {
				var ok bool
				if v := c1.Recv(&ok); v != 1 || !ok {
					t.Fatal(v, ok)
				}
			},
			func() {
				var ok bool
				if v := c2.Recv(&ok); v != 2 || !ok {
					t.Fatal(v, ok)
				}
			},
			func() { c3.Send(3) },
			func() { c4.Send(4) },
		} {
			if i+1 != selected {
				f()
			}
		}
	}

	if len(result) < 2 {
		t.Fatal("not random")
	}
}
