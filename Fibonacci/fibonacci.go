package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fibRecurs(5))
	fmt.Println(fibLoop(5))
	fmt.Println(fibCalc(5))
}

// * It is exponenial: O(2^n)
func fibRecurs(counter int) int {
	if counter <= 1 {
		return counter
	}
	return fibRecurs(counter-1) + fibRecurs(counter-2)
}

// * It is linear: O(n)
func fibLoop(max int) int {
	numOne := 0
	numTwo := 1
	fib := 0
	for i := 0; i < max; i++ {
		fib = numOne + numTwo
		numOne = numTwo
		numTwo = fib
	}
	return numOne
}

// * It is constant: O(1)
func fibCalc(max float64) float64 {
	return (((1.618034) * math.Exp(max)) - ((-0.618034) * math.Exp(max))) / math.Sqrt(5)
}
