package array

import (
	"fmt"
	"time"
)

// linear search performs linear search on an ordered array

func linearSearch(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
		if v > value {
			break
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	valueToFind := 7

	start := time.Now() //Get the current time before function call

	result := linearSearch(arr, valueToFind)

	elapsed := time.Since(start)

	if result != -1 {
		fmt.Printf("Value %d found at index %d\n", valueToFind, result)
	} else {
		fmt.Printf("Value %d not found in array \n", valueToFind)
	}

	fmt.Printf("Total execution: %s\n", elapsed)
}
