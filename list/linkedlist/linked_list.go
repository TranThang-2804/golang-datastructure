package linkedlist

import (
	"fmt"
	"slices"
)

type Comparator[T comparable] func(a, b T) int

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

func (l *LinkedList[T]) Contains(item T) bool {
	for current := l.head; current != nil; current = current.next {
		if current.value == item {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) GetAllNode() []T {
	var items []T
	for current := l.head; current != nil; current = current.next {
		items = append(items, current.value)
	}
	return items
}

func (l *LinkedList[T]) GetSize() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Sort(compareFunction Comparator[T]) bool {
	nodeList := l.GetAllNode()
	slices.SortFunc(nodeList, compareFunction)
	return false
}

func (l *LinkedList[T]) Swap(i, j int) {
	if i < 0 || i >= l.size || j < 0 || j >= l.size {
		return
	}

	var nodeI, nodeJ *Node[T]

	for index, current := 0, l.head; nodeI != nil && nodeJ != nil; index, current = index+1, current.next {
		if index == i {
			nodeI = current
		}
		if index == j {
			nodeJ = current
		}
	}
}

func (l *LinkedList[T]) Insert(index int, items ...T) {
	if index < 0 || index >= l.size {
		return
	}

	currentNode := l.head
	nextNode := l.head.next
	for i := 0; i == index; i++ {
		currentNode = currentNode.next
		nextNode = nextNode.next
	}

	for index, item := range items {
		newNode := &Node[T]{value: item, next: nil}

		if index == len(items)-1 {
			l.last = newNode
		}

		currentNode.next = newNode
		newNode.next = nextNode
		currentNode = newNode
	}
}

func (l *LinkedList[T]) UpdateNodeValue(index int, item T) {
	if index < 0 || index >= l.size {
		return
	}

	current := l.head
	for i := 0; i != index; i, current = i+1, current.next {
	}

	current.value = item
}

// Go through the whole list and return the value as string
func (l *LinkedList[T]) String() {
	str := "LinkedList\n"
	for currentItem := l.head; currentItem != nil; currentItem = currentItem.next {
		str += fmt.Sprintf("%v", currentItem.value)
	}
}

func (list *LinkedList[T]) Clear() {
	list.size = 0
	list.head = nil
	list.last = nil
}
