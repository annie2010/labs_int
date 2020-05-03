package fib

// Compute computes a Fibonacci number.
func Compute(n int) int {
	if n < 2 {
		return n
	}

	return Compute(n-2) + Compute(n-1)
}
