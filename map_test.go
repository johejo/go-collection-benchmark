package main

import (
	"strconv"
	"testing"
)

func BenchmarkEmptyMap(b *testing.B) {
	b.ResetTimer()
	m := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = i
	}
}

func BenchmarkNMap(b *testing.B) {
	b.ResetTimer()
	m := make(map[string]interface{}, b.N)
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = i
	}
}
