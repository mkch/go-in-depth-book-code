package channel_test

import (
	"example/channel"
	"fmt"
)

func ExampleChannel() {
	c := channel.Make[int](0)

	go c.Send(1)

	var ok bool
	v := c.Recv(&ok)
	fmt.Println(v, ok)
	// Output:
	// 1 true
}

func ExampleSelect() {
	c1 := channel.Make[string](0)
	c2 := channel.Make[string](1)

	channel.Select(
		[]*channel.Recv[string]{{Chan: c1,
			F: func(v string, ok bool) { fmt.Println("Recv selected") }}},
		[]*channel.Send[string]{{Chan: c2,
			F: func() { fmt.Println("Send selected") }}},
		nil,
	)

	// Output:
	// Send selected
}
