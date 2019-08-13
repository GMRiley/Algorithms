package main

import "fmt"

func main() {
	stack := data{}
	stack.push(10)
	stack.push(11)
	stack.push(1)
	stack.push(0)
	stack.push(5)

	fmt.Println(stack.size())
	fmt.Println(stack.pop())
}
