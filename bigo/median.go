package bigo

/* The following function accepts an array of strings and returns a new array that only contains the strings that start with the character "a".
Use Big O Notation to describe the time complexity of the function:
*/

// median calculates the median from an ordered array of integers.
func median(array []int) float64 {
	middle := len(array) / 2

	// If the array has an even number of elements, return the average of the middle two elements.
	if len(array)%2 == 0 {
		return float64(array[middle-1]+array[middle]) / 2.0
	}

	// If the array has an odd number of elements, return the middle element.
	return float64(array[middle]) // Conversion to float64 for consistency with even count case
}

// Time complexity -- Constant time
