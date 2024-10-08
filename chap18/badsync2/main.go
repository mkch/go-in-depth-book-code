package main

import "fmt"

var done bool

func main() {
	go func() { done = true }()
	for !done {
	}
	fmt.Println("done")
}
