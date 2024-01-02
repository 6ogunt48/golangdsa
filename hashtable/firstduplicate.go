package hashtable

import "fmt"

// First Duplicate in an array of string
func firstDuplicate(arr []string) string {
	m := make(map[string]bool)

	for _, str := range arr {
		if _, exists := m[str]; exists { //map return two values the exists is a boolean if the key exists
			return str
		}
		m[str] = true
	}
	return ""
}

func main() {
	fmt.Println("First Duplicate:", firstDuplicate([]string{"a", "b", "c", "d", "c", "e", "f"}))
}
