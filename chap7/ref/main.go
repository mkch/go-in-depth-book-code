package main

func main() {

}

// Push 把 value 压入 stack 中
// 此处 stack 为引用语义
func Push(stack *[]int, value int) {
	*stack = append(*stack, value)
}
