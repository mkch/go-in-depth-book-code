/*
Hello 输出欢迎信息。
本程序会在控制台上输出“Hello <名字>!”，其中 <名字> 为命令行参数。
如果没有提供命令行参数，则默认名字为 World。

使用方法：

	hello [flags] [name]

可用的 flags 有：

	-q 在名字前后添加引号。
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var q bool
	flag.BoolVar(&q, "q", false, "是否为名字添加引号")
	flag.Parse()
	var name = "World"
	if flag.NArg() > 0 {
		if flag.NArg() > 1 {
			fmt.Fprintln(os.Stderr, "参数太多")
			os.Exit(1)
		}
		name = flag.Arg(0)
	}
	if q {
		name = fmt.Sprintf("%q", name)
	}
	fmt.Printf("Hello %v!\n", name)
}
