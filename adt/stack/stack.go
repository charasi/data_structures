package stack

import "errors"

// default array length
const defaultSize = 10

type Stack struct {
	elements []any
	top      int
}

// insert a new item onto the top of the stack.
// throw error if stack is full
func (s *Stack) Push(element any) error {
	if s.top == len(s.elements) {
		return errors.New("stack overflow")
	}
	s.top++
	s.elements[s.top] = element
	return nil
}

// removes and returns the element at the top
func (s *Stack) Pop() (any, error) {
	if s.top == -1 {
		return nil, errors.New("stack underflow")
	}

	val := s.elements[s.top]
	s.elements[s.top] = nil
	s.top--
	return val, nil
}

// get element without removing from stack
func (s *Stack) Peek() (any, error) {
	if s.top == -1 {
		return nil, errors.New("stack underflow")
	}
	return s.elements[s.top], nil
}

// check if array is empty
func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

// create a new stack
func NewStack(size int) *Stack {
	if size < 0 {
		size = defaultSize
	}
	// create slice of length size
	a := make([]any, size)
	top := -1
	//
	s := &Stack{a, top}

	return s
}
