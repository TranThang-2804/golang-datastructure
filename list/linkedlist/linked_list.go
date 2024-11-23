package linkedlist

import "fmt"

// Node is a single element in a linked list.
type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	last *Node[T]
	size int
}

// Create a new empty linked list
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, last: nil, size: 0}
}

// Append a new node to the end of linked list
func (l *LinkedList[T]) Append(items ...T) {
	for _, item := range items {
		newNode := &Node[T]{value: item, next: nil}
		if l.size == 0 {
			l.head = newNode
		} else {
			l.last.next = newNode
		}
		l.last = newNode
		l.size++
	}
}

// Append a new node to the beginning of linked list
func (l *LinkedList[T]) Prepend(items ...T) {
	for _, item := range items {
		newNode := &Node[T]{value: item, next: nil}
		if l.size == 0 {
			l.last = newNode
		} else {
			newNode.next = l.head
		}
		l.head = newNode
		l.size++
	}
}

// Delete the item of the link list
func (c *Node[T]) Delete(value T) {

}

// Go through the whole list and return the value
func (n *Node[T]) Traverse() {
	str := "LinkedList\n"
	for n != nil {
		str += fmt.Sprintf("%v", n.value)
	}
}
