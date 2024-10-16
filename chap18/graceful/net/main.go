package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown := func() {
		cancel()
		group.Wait()
	}

	startServer(ctx)
	time.Sleep(time.Second * 50)
	shutdown()
	fmt.Println("exit")
}

var group sync.WaitGroup

func startServer(ctx context.Context) {
	group.Add(1)
	go func() {
		defer group.Done()
		var config net.ListenConfig
		listener, err := config.Listen(ctx, "tcp", ":1234")
		if err != nil {
			if errors.Is(err, context.Canceled) {
				// 收到退出信号
				fmt.Println("interrupted: Listen")
				return
			}
			panic(err) // 忽略错误处理
		}
		// 在 ctx 完成时关闭 listener
		context.AfterFunc(ctx, func() { listener.Close() })
		for {
			conn, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					// 收到退出信号
					fmt.Println("interrupted: Accept")
					return
				}
				panic(err) // 忽略错误处理
			}
			group.Add(1)
			go serveClient(ctx, conn)
		}
	}()
}

func serveClient(ctx context.Context, conn net.Conn) {
	defer group.Done()
	defer conn.Close()
	// 在 ctx 完成时关闭 conn
	context.AfterFunc(ctx, func() { conn.Close() })
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		if errors.Is(err, net.ErrClosed) {
			// 收到退出信号
			fmt.Println("interrupted: Read")
			return
		}
		panic(err) // 忽略错误处理
	}
	fmt.Printf("Recv from %v: %v\n", conn.RemoteAddr(), msg)
}
