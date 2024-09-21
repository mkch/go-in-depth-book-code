package main

type Request struct{}

func handle(*Request) {}

func Serve(q chan *Request) {
	for r := range q {
		go handle(r)
	}
}

func MaxGoServer(q chan *Request) {
	const MAX_GO = 16 // 最大并发数
	var sem = make(chan struct{}, MAX_GO)
	for r := range q {
		sem <- struct{}{} // 向 sem 写入一个值 (获取资源)
		go func() {
			handle(r)
			<-sem // 从 sem 读出一个值 (释放资源)
		}()
	}
}

func MaxGoServer2(q chan *Request, quit chan struct{}) {
	const MAX_GO = 16 // 最大并发数
	// 一次性创建固定个数的 goroutine
	for range MAX_GO {
		go process(q)
	}
	<-quit // 等待退出
}

func process(q chan *Request) {
	for r := range q {
		handle(r)
	}
}
