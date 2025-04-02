package mathext

import (
	"math"

	"github.com/tanveerprottoy/stdlib-ext/typesext"
)

func Add[T typesext.Numeric](a, b T) T {
	return a + b
}

func Subract[T typesext.Numeric](a, b T) T {
	return a - b
}

func Multiply[T typesext.Numeric](a, b T) T {
	return a * b
}

func Divide[T typesext.Numeric](a, b T) T {
	return a / b
}

func Percentage(part, total int64) float64 {
	if part == 0 || total == 0 || part > total {
		return 0
	}

	return float64(part) / float64(total) * 100
}

// RoundFloat rounds a float to a given precision
func RoundFloat(f float64, p int) float64 {
	pow := math.Pow(10, float64(p))
	return math.Round(f*pow) / pow
}

// The factorial of n is denoted by n! and calculated by multiplying the integer
// numbers from 1 to n. The formula for n factorial is n! = n × (n - 1)!.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}
