package stack

import ()

type Stack[T any] interface {
	Push(value ...T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
