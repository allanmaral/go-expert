package math

import "testing"

func BenchmarkAddUsingValue(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		m := New(i, i+1)
		_ = m.Add()
	}
}
