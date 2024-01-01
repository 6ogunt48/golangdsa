package bigo

import "fmt"

/*
The following function accepts an array of strings and returns a new array that only contains the strings that start with the character "a".
Use Big O Notation to describe the time complexity of the function:
*/

func selectAString(array []string) []string {
	var newArray []string
	for _, str := range array {
		if len(str) > 0 && str[0] == 'a' {
			newArray = append(newArray, str)
		}
	}
	return newArray
}

func main() {
	array := []string{"bola", "shade", "apprentice", "araujo", "amala"}
	newarray := selectAString(array)
	fmt.Println(newarray)
}
