package stack

import (
	"errors"
	"fmt"
)

// Stack represents a stack that holds integers
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes the top item from thr stack and returns it.
// Returns an error if the stack is empty.
func (s *Stack) pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

// Peek returns the top item without removing it.
// Returns an error if the stack is empty

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	myStack := Stack{}
	myStack.Push(1)
	myStack.Push(7)

	topItem, _ := myStack.Peek()
	fmt.Println("Top item:", topItem)

	poppedItem, _ := myStack.pop()
	fmt.Println("Popped item:", poppedItem)

	fmt.Println("is stack empty?", myStack.IsEmpty())
}
