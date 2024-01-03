package main

import (
	"fmt"
)

func main() {
	arr := []int{64, 25, 12, 22, 11, 31, 99, 73}
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		currentValue := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > currentValue {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = currentValue
	}
}
