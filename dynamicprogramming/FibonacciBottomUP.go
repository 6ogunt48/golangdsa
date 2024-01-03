package main

import "fmt"

// FibonacciBottomUp computes the nth Fibonacci number using bottom-up approach.

func FibonacciBottomUp(n int) int {
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}
	return curr
}

func main() {
	n := 10
	fmt.Printf("Fibonacci(%d) = %d\n", n, FibonacciBottomUp(n))
}
