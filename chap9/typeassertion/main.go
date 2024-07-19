package main

import (
	"io"
	"net"
	"os"
)

func main() {
	var rwc io.Writer = os.Stdin // 1
	r := rwc.(io.Reader)         // 2
	c := rwc.(io.Closer)         // 3
	conn, ok := rwc.(net.Conn)   // 4
	_ = conn
	//conn = rwc.(net.Conn) // 5: panic
	_, _, _, _ = r, c, conn, ok

	var x io.Writer
	_, ok = x.(io.Writer) // ok 为 false
	_ = x.(io.Writer)     // 6: panic
}
