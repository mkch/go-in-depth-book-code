package example_test

import "a.com/example"

func myadd(a, b int) int {
	return a + b
}

func ExampleAdderFunc() {
	var adder example.Adder = example.AdderFunc(myadd)
	adder.Add(0, 1)
}
