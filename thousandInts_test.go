package main

import (
	"testing"
)

func BenchmarkCustomarks(b *testing.B) {
	benchmarksFormula := []struct {
		name  string
		value int
	}{
		{"Thousand Formula", 1000},
		{"Billion Formula", 1000000000},
		{"Quintillion Formula", 1000000000000000000},
	}
	benchmarksLoop := []struct {
		name  string
		value int
	}{
		{"Thousand Loop", 1000},
		{"Billion Loop", 1000000000},
		{"Quintillion Loop", 1000000000000000000},
	}
	for _, bmF := range benchmarksFormula {
		b.Run(bmF.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumFormula(bmF.value)
			}
		})
	}
	for _, bmL := range benchmarksLoop {
		b.Run(bmL.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumLoop(bmL.value)
			}
		})
	}
}
