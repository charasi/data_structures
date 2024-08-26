package linkedList

// represent node of a linked list
type node struct {
	value any
	next  *node
}

// represent
type LinkedList struct {
	head *node
}

// InsertAfter insert element after the specified key
func (l *LinkedList) InsertAfter(key any, element any) {
	// return if data is null
	if element == nil {
		return
	}
	temp := l.head
	// insert first if list is nill
	if temp == nil {
		l.AddFirst(element)
		return
	}
	// loop through list until key is found, or to end of list
	for temp != nil && temp.value != key {
		temp = temp.next
	}
	// key was not found
	if temp == nil {
		return
	}
	// add element before key
	newNode := &node{value: element, next: nil}
	newNode.next = temp.next
	temp.next = newNode
}

// InsertBefore insert element before the specified key
func (l *LinkedList) InsertBefore(key any, element any) {
	// return if data is null
	if element == nil {
		return
	}

	// insert first if list is nil
	if l.head == nil {
		l.AddFirst(element)
		return
	}

	// loop through list until key is found, or to end of list
	prev, temp := l.head, l.head
	for temp != nil && temp.value != key {
		prev = temp
		temp = temp.next
	}
	// key was not found
	if temp == nil {
		return
	}
	// add element before key
	newNode := &node{value: element, next: nil}
	newNode.next = temp
	prev.next = newNode

}

// Remove an element from the list
func (l *LinkedList) Remove(element any) {
	// return if list or element is nil
	if l.head == nil {
		return
	}
	if element == nil {
		return
	}
	prev, temp := l.head, l.head
	// loop through list until key is found, or to end of list
	for temp != nil && temp.value != element {
		prev = temp
		temp = temp.next
	}
	// key was not found
	if temp == nil {
		return
	}
	prev.next = temp.next

}

// Get an element from list
func (l *LinkedList) Get(element any) any {
	// return if list or element is nil
	if l.head == nil {
		return nil
	}
	if element == nil {
		return nil
	}
	temp := l.head
	for temp != nil && temp.value != element {
		temp = temp.next
	}
	if temp == nil {
		return nil
	}
	return temp.value
}

// AddFirst adds a node to the beginning of a list
func (l *LinkedList) AddFirst(element any) {
	// return if data is null
	if element == nil {
		return
	}
	// add node to beginning of list
	temp := l.head
	newNode := &node{value: element, next: nil}
	newNode.next = temp
	l.head = newNode

}

// AddLast adds a node to the end of a list
func (l *LinkedList) AddLast(element any) {
	// return if data is null
	if element == nil {
		return
	}
	// insert first if list is nil
	if l.head == nil {
		l.AddFirst(element)
		return
	}
	// create node
	newNode := &node{value: element, next: nil}
	// add data to end of list
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// returns list as an array
func (l *LinkedList) ToArray() []any {
	if l.head == nil {
		return nil
	}
	temp := l.head
	arr := make([]any, 0)
	for temp != nil {
		arr = append(arr, temp.value)
		temp = temp.next
	}
	return arr
}
