package main

import "fmt"

func main() {
	var str string = "abc中文"
	var length int = len(str) // 1: length 的值为 9
	var b byte = str[0]       // 1: string 索引操作的结果类型为 byte
	var str2 string = str[3:] // 2: string 切片操作的结果类型为 string
	fmt.Println(length, b, str2)

	conv()
}

func conv() {
	var str string = "abc中文"
	var chars = []rune(str)      // 1: chars 存放上述 5 个字符的码位
	var nChar = len(chars)       // 2: nChar 为字符个数: 5
	var str2 = string(chars[4:]) // 3: str2 的内容为 "文"
	fmt.Println(nChar, str2)

	for i, r := range str {
		fmt.Printf("%v: %v(%[2]T) %v\n", i, r, string(r))
	}

}
