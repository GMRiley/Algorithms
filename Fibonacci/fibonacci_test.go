package main

import (
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	benchmarkFib := []struct {
		name  string
		value int
	}{
		{"Hundred", 50},
	}
	for _, bmF := range benchmarkFib {
		b.Run(bmF.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibCalc(bmF.value)
			}
		})
	}
}
