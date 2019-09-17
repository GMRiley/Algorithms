package main

import "fmt"

func main() {
	stack := data{}
	for counter := 0; counter < 5000000; counter++ {
		stack.push(10)
	}
	fmt.Println(stack.size())
	fmt.Println(stack.pop())
}
