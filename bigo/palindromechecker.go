package bigo

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	testStrings := []string{"radar", "madam", "hello"}
	for _, str := range testStrings {
		fmt.Printf("Is '%s' a palindrome? %t\n", str, isPalindrome(str))
	}
}
