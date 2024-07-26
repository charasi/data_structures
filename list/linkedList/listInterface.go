package linkedList

type linkedListInterface interface {
	AddFirst(element any)
	AddLast(element any)
	InsertAfter(key any, element any)
	InsertBefore(key any, element any)
	Remove(element any)
	Get(element any) any
	toArray() []any
}
