package main

import "fmt"

// intersection of two arrays where a is the larger array
func intersection(a, b []int) []int {
	m := make(map[int]bool)
	var inter []int

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			inter = append(inter, item)
		}

	}
	return inter
}

func main() {
	fmt.Println("Intersection:", intersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8}))
}
