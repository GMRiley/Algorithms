package main

import (
	"fmt"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	benchmarkFib := []struct {
		name  string
		value int
	}{
		{"Five", 5},
		{"TwentyFive", 25},
		{"Fifty", 50},
	}
	for _, bmF := range benchmarkFib {
		b.Run(bmF.name+"Loop", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibLoop(bmF.value)
			}
		})
	}
	fmt.Println("")
	for _, bmF := range benchmarkFib {
		b.Run(bmF.name+"Calc", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibCalc(float64(bmF.value))
			}
		})
	}
	fmt.Println("")
	for _, bmF := range benchmarkFib {
		b.Run(bmF.name+"Recurs", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibRecurs(bmF.value)
			}
		})
	}
}
