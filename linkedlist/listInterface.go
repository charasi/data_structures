package listinterface

type LinkedList interface {
	AddFirst(element interface{})
	AddLast(element interface{})
	InsertAfter(key interface{}, element interface{})
	//InsertBefore(key interface{}, element interface{})
	//remove(element interface{})
	//get(element interface{})
}
