package linkedliststack

import "github.com/TranThang-2804/golangds/list/linkedlist"

type Stack[T comparable] struct {
	list *linkedlist.LinkedList[T]
}

func (s *Stack[T]) Push(values ...T) {
	s.list.Prepend(values...)
}

func (s *Stack[T]) Peek() (T, bool) {
	return s.list.Get(0)
}

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

func (s *Stack[T]) Empty() {
	s.list.Clear()
}

func (s *Stack[T]) Size() int {
	return s.list.GetSize()
}

func (s *Stack[T]) IsEmpty() bool {
  return s.list.IsEmpty()
}

func (s *Stack[T]) String() string {
  return s.list.String()
}
