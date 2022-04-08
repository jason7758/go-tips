package main

import "testing"


func BenchmarkUseMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseMutex()
	}
}

func BenchmarkUseChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseChan()

	}
}
