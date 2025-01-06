package main

import (
	"io"
	"net"
	"os"
)

func main() {
	var rwc io.Writer = os.Stdin
	f := rwc.(*os.File)
	r := rwc.(io.Reader)
	c := rwc.(io.Closer)
	conn, ok := rwc.(net.Conn)
	_ = conn
	//conn = rwc.(net.Conn) // 5: panic
	_, _, _, _, _ = f, r, c, conn, ok

	var x io.Writer
	_, ok = x.(io.Writer) // ok 为 false
	_ = x.(io.Writer)     // 6: panic
}
