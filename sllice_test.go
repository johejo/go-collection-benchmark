package main

import (
	"testing"
)

func BenchmarkEmptySlice(b *testing.B) {
	b.ResetTimer()
	var s []string
	for i := 0; i < b.N; i++ {
		s = append(s, "Hello world!!")
	}
}

func Benchmark0Slice(b *testing.B) {
	b.ResetTimer()
	s := make([]string, 0)
	for i := 0; i < b.N; i++ {
		s = append(s, "Hello World!!")
	}
}

func BenchmarkNSlice(b *testing.B) {
	b.ResetTimer()
	s := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, "Hello World!!")
	}
}

func Benchmark0NSlice(b *testing.B) {
	b.ResetTimer()
	s := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, "Hello World!!")
	}
}
