package stack

type stackInterface interface {
	Push(element any) error
	Pop() (any, error)
	Peek() (any, error)
	IsEmpty() bool
}
