package linkedlistqueue

import "github.com/TranThang-2804/golangds/list/linkedlist"

type LinkedListQueue[T comparable] struct {
	linkedList *linkedlist.LinkedList[T]
}

func New[T comparable]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

func (q *LinkedListQueue[T]) Enqueue(value T) {
	q.linkedList.Append(value)
}

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

func (q *LinkedListQueue[T]) Peek() (T, bool) {
  return q.linkedList.Get(0)
}
