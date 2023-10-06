package main

/*
	To run the benchmarks
	go test -bench=.
*/
import "testing"

func IsPrime(no int) bool {
	for i := 2; i < (no); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime2(no int) bool {
	end := (no / 2)
	for i := 2; i <= end; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func Benchmark_IsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(99973)
	}
}

func Benchmark_IsPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime2(99973)
	}
}
