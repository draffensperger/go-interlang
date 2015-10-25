package main

import "testing"

func BenchmarkGoToCNoOp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoToCNoOp()
	}
}

func BenchmarkGoOnlyNoOp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoOnlyNoOp()
	}
}
