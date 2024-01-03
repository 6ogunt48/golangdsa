package main

import "fmt"

// Find the missing letter
func missingLetter(s string) rune {
	letters := make([]bool, 26) //slice named letters of length 26 is created with false as default boolean values

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			letters[ch-'a'] = true
		} else if ch >= 'A' && ch <= 'Z' {
			letters[ch-'A'] = true
		}
	}
	for i, found := range letters {
		if !found {
			return rune('a' + i)
		}
	}

	return 0
}

func main() {
	fmt.Println("Missing Letter:", string(missingLetter("the quick brown box jumps over a lazy dog")))
}
