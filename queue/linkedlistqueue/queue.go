package linkedlistqueue

import (
	"fmt"
	"strings"

	"github.com/TranThang-2804/golangds/list/linkedlist"
)

// LinkedListQueue is a struct to represent a linked list queue
type LinkedListQueue[T comparable] struct {
	linkedList *linkedlist.LinkedList[T]
}

// New creates a new empty linked list queue
func New[T comparable]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

// Enqueue adds a value to the end of the queue
func (q *LinkedListQueue[T]) Enqueue(value T) {
	q.linkedList.Append(value)
}

// Dequeue removes the first element of the queue
func (q *LinkedListQueue[T]) Dequeue() (T, bool) {
  // Get the first element of the linked list
	t, err := q.linkedList.Get(0)

	if !err {
		return t, false
	}

  // Remove the first element of the linked list
	err = q.linkedList.Remove(0)

  if !err {
    return t, false
  }

	return t, true
}

// Peek returns the first element of the queue without removing it
func (q *LinkedListQueue[T]) Peek() (T, bool) {
  return q.linkedList.Get(0)
}

// Values returns all the elements of the queue as an array
func (q *LinkedListQueue[T]) Values() []T {
  return q.linkedList.GetAllNode()
}

// Size returns the number of elements
func (q *LinkedListQueue[T]) Size() int {
  return q.linkedList.GetSize()
}

// IsEmpty returns true if the queue is empty
func (q *LinkedListQueue[T]) IsEmpty() bool {
  return q.linkedList.IsEmpty()
}

// Return the string representation of the stack
func (s *LinkedListQueue[T]) String() string {
	str := "LinkedListStack\n"
	values := []string{}
	for _, value := range s.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
