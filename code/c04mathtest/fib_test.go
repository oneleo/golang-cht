package c04mathtest

import (
	"testing"
)

func Test_FibTestFunc(t *testing.T) { // OMIT START
	var fibTests = []struct {
		in, want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34},
	}
	for _, tt := range fibTests {
		if out := Fib(tt.in); out != tt.want {
			t.Errorf("Error on Fib(%d) = %d; want = %d\n", tt.in, out, tt.want) // OMIT END
		}
	}
}

func Test_FibTestFunc2(t *testing.T) { // OMIT START
	var fibTests = []struct {
		in, want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34},
	}
	for _, tt := range fibTests {
		if out := Fib(tt.in); out != tt.want {
			t.Errorf("Error on Fib(%d) = %d; want = %d\n", tt.in, out, tt.want) // OMIT END
		}
	}
}

func Benchmark_FibBenchFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
