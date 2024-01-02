package main

import (
	"errors"
	"fmt"
)

// Queue represents a FIFO (first in, first out) data structure.
type Queue struct {
	items []int
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes an item from the front of the queue and returns it.
// Returns an error if the queue is empty.
func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)

	item, _ := myQueue.Dequeue()
	fmt.Println("dequeued item:", item)

	fmt.Println("Is queue empty?", myQueue.IsEmpty())
}
