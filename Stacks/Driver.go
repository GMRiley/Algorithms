package main

import "fmt"

func main() {
	stack := data{}
	for counter := 0; counter < 500; counter++ {
		stack.push(10)
	}
	fmt.Println(stack.size())
	fmt.Println(stack.pop())
}
