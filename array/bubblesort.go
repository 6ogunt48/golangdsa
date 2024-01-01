package array

import "fmt"

// BubbleSort sorts an array using the bubble sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// swapped to check if any swap happened in this pass
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// if no two elements were swapped, the array is sorted
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{64, 32, 25, 12, 22, 11, 90}
	BubbleSort(arr)
	fmt.Printf("Sorted array: %v\n", arr)
}
