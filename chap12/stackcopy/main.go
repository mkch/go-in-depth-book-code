package main

import (
	"os"
	"strconv"
	"strings"
	"unsafe"
)

// n 的初始地址值.
var pn0 uintptr

// 调用深度.
var depth = 0

// f 栈帧的大致大小.
const FRAME_SIZE = 1024

type Frame [FRAME_SIZE]byte

func main() {
	var n int
	pn0 = uintptr(unsafe.Pointer(&n))
	f(&n, Frame{})
}

func f(pn *int, frame Frame) {
	depth++
	if uintptr(unsafe.Pointer(pn)) != pn0 {
		print(pn, depth)
		os.Exit(0)
	}
	f(pn, frame)
}

// print 输出 pn 和 depth.
// 不能使用 fmt 系列函数，因为它们会导致 n 分配到堆上.
func print(pn *int, depth int) {
	var buf strings.Builder
	buf.WriteString("pn0=0x")
	buf.WriteString(strconv.FormatUint(uint64(pn0), 16))
	buf.WriteString(" pn=0x")
	buf.WriteString(strconv.FormatUint(uint64(uintptr(unsafe.Pointer(pn))), 16))
	buf.WriteString(" depth=")
	buf.WriteString(strconv.Itoa(depth))
	buf.WriteRune('\n')
	os.Stdout.WriteString(buf.String())
}
