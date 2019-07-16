package main

import (
	"testing"
)

func BenchmarkThousand(b *testing.B) {
	for val := 0; val < b.N; val++ {
		sum(1000)
	}
}
func BenchmarkBillion(b *testing.B) {
	for val := 0; val < b.N; val++ {
		sum(1000000000)
	}
}
func BenchmarkQuintillion(b *testing.B) {
	for val := 0; val < b.N; val++ {
		sum(1000000000000000000)
	}
}
