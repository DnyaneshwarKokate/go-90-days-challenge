package bench_test

import (
	"strconv"
	"testing"

	"day-43/bench"
)

func generateSlice(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = "item_" + strconv.Itoa(i)
	}
	return slice
}

func BenchmarkConcatPlus(b *testing.B) {
	items := generateSlice(100)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bench.ConcatPlus(items)
	}
}

func BenchmarkConcatSprintf(b *testing.B) {
	items := generateSlice(100)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bench.ConcatSprintf(items)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	items := generateSlice(100)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bench.ConcatBuilder(items)
	}
}
