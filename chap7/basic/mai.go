package main

import "fmt"

func main() {
	var str string = "abc中文"
	var length int = len(str) // 1: length 的值为 9
	var b byte = str[0]       // 1: string 索引操作的结果类型为 byte
	var str2 string = str[3:] // 2: string 切片操作的结果类型为 string
	fmt.Println(length, b, str2)

	conv()
	raw()
	interpreted()
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

func raw() {
	str := `a
b\n中`
	fmt.Printf("%X\n", str)
}

func interpreted() {
	str := "中文"                        // 字面量
	str2 := `中文`                       // 原始字面量
	str3 := "\u4E2D\u6587"             // UTF-8 码位
	str4 := "\U00004E2D\U00006587"     // UTF-8 码位
	str5 := "\xE4\xB8\xAD\xE6\x96\x87" // 16 进制 UTF-8 字节序列
	str6 := "\344\270\255\346\226\207" // 8 进制 UTF-8 字节序列
	fmt.Println(str, str2, str3, str4, str5, str6)
}
