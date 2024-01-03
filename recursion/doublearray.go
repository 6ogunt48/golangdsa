package main

import "fmt"

func DoubleArray(arr []int, index int) {
	// Base Case: if index is equal to the length of the array, return
	if index == len(arr) {
		return
	}

	// double the value at current index
	arr[index] *= 2

	//Recursive call for the next index
	DoubleArray(arr, index+1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	DoubleArray(arr, 0)
	fmt.Println("Double Array:", arr)
}
