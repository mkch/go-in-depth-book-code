package main

import (
	"fmt"
	"math"
	"strconv"
)

func F2_Fixed() {
	var c = make(chan string)

	go func() {
		c <- "1"
	}()

	strconv.Atoi(<-c)
}

func F2() {
	var str string

	go func() {
		str = "1"
	}()

	strconv.Atoi(str)
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
