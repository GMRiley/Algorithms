package main

import "fmt"

func main() {
	answer := fibCalc(100)
	fmt.Println(answer)
}
func fibCalc(counter int) int {
	if counter <= 1 {
		return counter
	}
	return fibCalc(counter-1) + fibCalc(counter-2)
}
func fibLoop(max int) int {
	total := 0
	for i := 0 i < max; i++{
		
	}
}
