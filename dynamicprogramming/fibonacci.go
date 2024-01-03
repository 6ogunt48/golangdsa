package main

import (
	"fmt"
)

// Fibonacci computes the nth Fibonacci number.
func Fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	// check if a result is already in memo
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	// Store the result in memo and return it
	memo[n] = Fibonacci(n-1, memo) + Fibonacci(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int)
	fmt.Println("Fibonacci 10:", Fibonacci(10, memo))
}
