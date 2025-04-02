package mathext_test

import (
	"math/rand"
	"testing"

	"github.com/tanveerprottoy/stdlib-ext/mathext"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		val0 int
		val1 int
		exp  int
	}{
		{"2 + 5", 2, 5, 7},
		{"9 + 5", 9, 5, 14},
		{"27 + 45", 27, 45, 72},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := mathext.Add(tc.val0, tc.val1)
			if actual != tc.exp {
				t.Errorf("Add(%d, %d) = %v; want %v", tc.val0, tc.val1, actual, tc.exp)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		val0 int
		val1 int
		exp  int
	}{
		{"5 - 3", 5, 3, 2},
		{"9 - 5", 9, 5, 4},
		{"7 - 5", 7, 5, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := mathext.Subract(tc.val0, tc.val1)

			if actual != tc.exp {
				t.Errorf("Subract(%d, %d) = %v; want %v", tc.val0, tc.val1, actual, tc.exp)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name  string
		num   int64
		denom int64
		exp   float64
	}{
		{"50 percent", 50, 100, 50.0},
		{"25 percent", 1, 4, 25.0},
		{"0 percent", 0, 100, 0.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := mathext.Percentage(tc.num, tc.denom)

			if actual != tc.exp {
				t.Errorf("Percentage(%d, %d) = %v; want %v", tc.num, tc.denom, actual, tc.exp)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		in   int
		exp  int
	}{
		{"2!", 2, 2},
		{"3!", 3, 6},
		{"4!", 4, 24},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := mathext.Factorial(tc.in)

			if actual != tc.exp {
				t.Errorf("Factorial(%d) = %v; want %v", tc.in, actual, tc.exp)
			}
		})
	}
}

// benchmarks
func BenchmarkAdd(b *testing.B) {
	for i := range b.N {
		_ = mathext.Add(rand.Intn(i), rand.Intn(i+2))
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := range b.N {
		_ = mathext.Subract(rand.Intn(i), rand.Intn(i+2))
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := range b.N {
		_ = mathext.Factorial(rand.Intn(i))
	}
}
