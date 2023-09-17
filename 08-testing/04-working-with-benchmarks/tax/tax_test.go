package tax

import "testing"

// We can run benchmarks with the command `go test -bench=.`
// Some helpfull flags
//   * -benchmem: Benchmark the memory allocation
//   * -run=<selector>: Run only the tests match the selector
//   * -count=<number>: Define the number of times the banchmark will run for
//   * -benchtime=<duration>: Define the duration for each benchmark

func BenchmarkCalculateTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
