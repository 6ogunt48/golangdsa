package main

import "fmt"

// binarySearch performs a binary search on an ordered array

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		// Check if the target is present at mid
		if arr[mid] == target {
			return mid
		}

		// if the target is greater, ignore the left half
		if arr[mid] < target {
			low = mid + 1
		} else {
			// if the target is smaller, ignore the right half
			high = mid - 1
		}

	}
	// the target was not found
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	target := 10

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Println("Element not found in array")
	}
}
