package queue

import "errors"

// default array length
const defaultSize = 10

type Queue struct {
	elements []any
	// back of the queue
	back int
	// front of the queue
	front int
	// keep track of current number of items
	nItems int
}

func NewQueue(size int) *Queue {
	if size < 0 {
		size = defaultSize
	}
	// create slice of length size
	a := make([]any, size)
	b, f, n := -1, 0, 0
	q := &Queue{a, b, f, n}
	return q
}

// Enqueue insert a new item into the back of the queue
func (q *Queue) Enqueue(element any) error {
	if q.nItems == len(q.elements) {
		return errors.New("queue is full")
	}
	// increment back when new element is added to queue
	q.back++
	// perform modulo operation to circle back to
	// beginning of array
	index := q.back % len(q.elements)
	q.elements[index] = element
	q.nItems++
	return nil
}

// Dequeue returns and removes the item at front
func (q *Queue) Dequeue() (any, error) {
	if q.nItems == -1 {
		return nil, errors.New("queue is empty")
	}
	// modulo operation to get right index
	index := q.front % len(q.elements)
	value := q.elements[index]
	q.elements[index] = nil
	q.front++
	q.nItems--
	return value, nil
}

// Peek returns the first item in the queue without removing it
func (q *Queue) Peek() (any, error) {
	if q.nItems == -1 {
		return nil, errors.New("queue is empty")
	}
	return q.elements[q.front], nil
}

// IsEmpty checks if array is empty
func (q *Queue) IsEmpty() bool {
	if q.nItems == -1 {
		return true
	}
	return false
}
