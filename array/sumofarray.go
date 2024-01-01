package array

import "fmt"

// Use Big O Notation to describe the time complexity of the following function that sums up all the numbers from a given array:
// Time Complexity O(N)
func arraySum(array []int) int {
	var sum int

	for _, v := range array {
		sum += v
	}

	return sum
}

func main() {
	array := []int{200, 300, 400, 500, 600, 700, 800}
	sum := arraySum(array)
	fmt.Println(sum)
}
