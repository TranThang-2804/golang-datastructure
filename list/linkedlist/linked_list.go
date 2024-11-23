package linkedlist

import "fmt"

// Node is a single element in a linked list.
type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	last *Node[T]
	size int
}

// Create a new empty linked list
func New[T comparable]() *LinkedList[T] {
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

func (l *LinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var t T
		return t, false
	}

	node := l.head
	for i := 0; i != index; i, node = i+1, node.next {
	}

	return node.value, true
}

// Remove the item of the link list
func (l *LinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}

	// current node starts at index 1
	prevNode := l.head
	currentNode := l.head.next

	// If the index is 0, then we need to remove the head
	if index == 0 {
		l.head = prevNode.next
		prevNode = nil
	} else {
		for i := 1; i < index; i++ {
			prevNode = prevNode.next
			currentNode = currentNode.next
		}

		prevNode.next = currentNode.next
		currentNode = nil
	}

	l.size--
	return true
}

func (l *LinkedList[T]) Constains(item T) bool {
	for current := l.head; current != nil; current = current.next {
		if current.value == item {
			return true
		}
	}
	return false
}

// Go through the whole list and return the value
func (n *Node[T]) Traverse() {
	str := "LinkedList\n"
	for n != nil {
		str += fmt.Sprintf("%v", n.value)
	}
}
