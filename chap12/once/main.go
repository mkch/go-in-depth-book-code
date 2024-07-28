package main

import (
	"sync"
	"time"
)

func main() {
	Instance()
	Instance2()
	go Instance()
	go Instance2()
	time.Sleep(time.Second)
}

type Singleton struct{}

var singletonOnce sync.Once
var instance *Singleton

func Instance() *Singleton {
	singletonOnce.Do(func() {
		instance = new(Singleton) // 此行代码只会执行一次
	})
	return instance
}

var getInstance = sync.OnceValue(func() *Singleton {
	return new(Singleton) // 此行代码只会执行一次
})

func Instance2() *Singleton {
	return getInstance()
}
