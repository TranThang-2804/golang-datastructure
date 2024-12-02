package linkedliststack

import "github.com/TranThang-2804/golangds/list/linkedlist"

// This is the stack implementation using linked list
type Stack[T comparable] struct {
	list *linkedlist.LinkedList[T]
}

// Constructor for creating linkedlist stack
func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: linkedlist.New[T]()}
}

// Push the items into the stack
func (s *Stack[T]) Push(values ...T) {
	s.list.Prepend(values...)
}

// Get the value of the item at the top of
// the stack, return the value of the item and
// true if the item is found else return false
func (s *Stack[T]) Peek() (T, bool) {
	return s.list.Get(0)
}

// Pop the item from the top of the stack
// return the value of the item and true if
// the item is found else return false
func (s *Stack[T]) Pop() (T, bool) {
	value, err := s.Peek()
	if err {
		var zeroValue T
		return zeroValue, false
	}
	if s.list.Remove(0) {
		var zeroValue T
		return zeroValue, false
	}
	return value, true
}

// Empty the stack
func (s *Stack[T]) Empty() {
	s.list.Clear()
}

// Get the size of the stack
func (s *Stack[T]) Size() int {
	return s.list.GetSize()
}

// Check if the stack is is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Return the string representation of the stack
func (s *Stack[T]) String() string {
	return s.list.String()
}
