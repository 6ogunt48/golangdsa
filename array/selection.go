package main

import "fmt"

func main() {
	arr := []int{64, 32, 25, 12, 22, 11, 90}
	selectionSort(arr)
	fmt.Printf("Sorted array: %v\n", arr)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Assume the minimum is the first element
		lowest := i

		// Test against elements after i to find the smallest
		for j := i + 1; j < n; j++ {
			// if this element is less, then it is the new minimum
			if arr[j] < arr[lowest] {
				lowest = j
			}
		}
		//	if the minimum isn't in the position, swap it
		if i != lowest {
			arr[i], arr[lowest] = arr[lowest], arr[i]
		}

	}
}
