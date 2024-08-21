package example_test

import (
	"fmt"

	"a.com/example"
)

func Example_extra() {
	_ = example.Add(1, 2)
}

func ExampleAdd() {
	sum := example.Add(1, 2)
	fmt.Println(sum)
	// Output:
	// 3
}

func ExampleAdd_neg() {
	fmt.Println(example.Add(-1, -2))
	// Output:
	// -3
}
