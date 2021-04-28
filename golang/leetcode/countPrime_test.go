package leetcode_test

import (
	"learn/golang/leetcode"
	"testing"
)

func BenchmarkCountPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CountPrime(10000)
	}
}

func BenchmarkCountPrimeEratothenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.CountPrimeEratothenes(10000)
	}
}
