package main

import (
	"testing"
	"fmt"
)

func BenchmarkFibonacci(b *testing.B) {
	benchmarkFib := []struct {
		name  string
		value int
	}{
		{"Five", 5},
		{"Ten", 10},
		{"Fifty", 50},
	}
	fmt.println("Loop Results")
	for _, bmF := range benchmarkFib{
		b.Run(bmF.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibLoop(bmF.value)
			}
		})
	}
	fmt.println("Calculation Results")
	for _, bmF := range benchmarkFib {
		b.Run(bmF.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibCalc(float64(bmF.value))
			}
		})
	}
	fmt.println("Recursion Results")
	for _, bmF := range benchmarkFib{
		b.Run(bmF.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibRecurs(bmF.value)
			}
		})
	}
}
