package main

import (
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(conn, "Hello\n")
	if err != nil {
		panic(err)
	}
}
