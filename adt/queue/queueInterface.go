package queue

type queueInterface interface {
	Enqueue(element any) error
	Dequeue() (any, error)
	Peek() (any, error)
	IsEmpty() bool
}
