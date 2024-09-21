package main

import (
	"bufio"
	"context"
	"errors"
	"log"
	"net"
	"os"
	"time"
)

func StartServer(ctx context.Context, port string) {
	// 1: 使用 ctx 作为参数
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", net.JoinHostPort("", port))
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		log.Panic(err)
	}

	// 2: 当 ctx 被取消时，调用 l.Close
	context.AfterFunc(ctx, func() { l.Close() })

	for {
		conn, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				// l 已经因为 ctx 被取消而被关闭了
				break
			}
			log.Printf("Server Error: %v\n", err)
			continue
		}

		// 开一个 Worker goroutine 服务 conn
		go func(clientConn net.Conn) {
			defer clientConn.Close()
			// 3: 当 ctx 被取消时，让 clientConn 立即到期（超时）
			context.AfterFunc(ctx, func() { clientConn.SetDeadline(time.Now()) })
			// 和客户端通讯
			msg, err := bufio.NewReader(clientConn).ReadString('\n')
			if err != nil {
				if errors.Is(err, os.ErrDeadlineExceeded) {
					// clientConn 已经因为 ctx 被取消而到期了
					return
				}
				log.Printf("Server Error: %v\n", err)
				return
			}
			_ = msg
			// ...
		}(conn)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown := make(chan struct{})
	go func() {
		defer close(shutdown)
		StartServer(ctx, "8888")
	}()

	// 需要退出 server 时
	cancel()

	<-shutdown
}
