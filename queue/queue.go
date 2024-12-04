package queues

type Queue[T comparable] interface {
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)
}
