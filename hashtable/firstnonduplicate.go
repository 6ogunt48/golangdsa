package hashtable

import "fmt"

// First non-duplicated character
func firstNonDuplicate(s string) rune {
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range s {
		if count[ch] == 1 {
			return ch
		}
	}
	return 0
}

func main() {
	fmt.Println("First Non-Duplicate:", string(firstNonDuplicate("minimum")))
}
