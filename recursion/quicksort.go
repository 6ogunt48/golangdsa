package main

import "fmt"

// quickSort using the Hoare partition scheme.
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := hoarePartition(arr, low, high)

		quickSort(arr, low, pi)
		quickSort(arr, pi+1, high)
	}
}

// hoarePartition function for the Hoare partition scheme.
func hoarePartition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		// Move i to the right at least once and until arr[i] >= pivot
		i++
		for arr[i] < pivot {
			i++
		}

		// Move j to the left at least once and until arr[j] <= pivot
		j--
		for arr[j] > pivot {
			j--
		}

		//if two pointers met.
		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)
	quickSort(arr, 0, n-1)
	fmt.Println("Sorted array: ", arr)
}
