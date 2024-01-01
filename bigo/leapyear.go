package bigo

import "fmt"

// Use Big O Notation to describe the time complexity of the following function that determines whether a given year is a leap year:

// Time complexity: 0(1)
func isLeapYear(year int) bool {
	// check if the year is divisible by 100
	if year%100 == 0 {
		// if yes, check if this also divisible by 100
		return year%400 == 0
	}
	// if not divisible by 100, check if its divisible 4
	return year%4 == 0

}

func main() {
	testYears := []int{1996, 1900, 2000, 2023}
	for _, year := range testYears {
		fmt.Printf("Year %d is a leap year: %t\n", year, isLeapYear(year))
	}
}
