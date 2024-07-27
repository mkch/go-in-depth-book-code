package main

import (
	"fmt"
	"syscall"
	"time"
)

// 加载 Windows 系统 DLL, 定位 Win32 API 函数地址
var lzKernel32 = syscall.MustLoadDLL("kernel32")
var lzGetCurrentThreadId = lzKernel32.MustFindProc("GetCurrentThreadId")

// GetCurrentThreadId 直接调用 Win32 API 获取当前线程 Id.
// 此函数只能用于 Windows 系统!
func GetCurrentThreadId() uint32 {
	id, _, _ := lzGetCurrentThreadId.Call()
	return uint32(id)
}

func main() {
	// goroutine 1
	go func() {
		fmt.Println(GetCurrentThreadId()) // 1
		time.Sleep(time.Second * 1)
		fmt.Println(GetCurrentThreadId()) // 2
	}()

	// goroutine 2
	go func() {
		time.Sleep(time.Second * 1)
	}()

	time.Sleep(time.Second * 2) // 强制等待 2 秒. 以后会有更好的实现
}
