package main

import (
	"fmt"
)

// quickSort sorts an array using the Quick Sort algorithm.
func quickSortlomuto(arr []int, low, high int) {
	if low < high {
		// pi is partitioning index, arr[pi] is now at the right place
		pi := lomutoPartition(arr, low, high)

		// Separately sort elements before partition and after partition
		quickSortlomuto(arr, low, pi-1)
		quickSortlomuto(arr, pi+1, high)
	}
}

// lomutoPartition takes the last element as pivot, places the pivot element at its correct
// position in sorted array, and places all smaller (smaller than pivot) to left of pivot
// and all greater elements to right of pivot.
func lomutoPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)
	quickSortlomuto(arr, 0, n-1)
	fmt.Println("Sorted array:", arr)
}
