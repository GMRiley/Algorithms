package main

import (
	"fmt"
	"math"
)

func main() {
	answer := fibCalc(100)
	fmt.Println(answer)
}
func fibRecurs(counter int) int {
	if counter <= 1 {
		return counter
	}
	return fibRecurs(counter-1) + fibRecurs(counter-2)
}
func fibLoop(max int) int {
	numOne := 0 
	numTwo := 1
	fib := 0
	for i := 0; i < max; i++ {
		fib = numOne + numTwo
		numOne = numTwo
		numTwo = fib
	}
	return fib
}
func fibCalc(max float64) float64 {
	return (((1.618034)* math.Exp(max)) - ((-0.618034)* math.Exp(max))) / math.Sqrt(5)
}
