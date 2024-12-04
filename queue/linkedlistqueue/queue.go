package linkedlistqueue

import "github.com/TranThang-2804/golangds/list/linkedlist"

type LinkedListQueue[T comparable] struct {
  linkedList *linkedlist.LinkedList[T]
}

func New[T comparable]() *LinkedListQueue[T] {
  return &LinkedListQueue[T]{}
}
