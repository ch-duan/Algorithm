package main

import (
	"strconv"
	"testing"
)

func BenchmarkB(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
