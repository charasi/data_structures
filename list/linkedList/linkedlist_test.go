package linkedList

import (
	"log"
	"os"
	"testing"
)

type ll struct {
	list linkedListInterface
}

var list ll

func TestMain(m *testing.M) {
	list = ll{
		list: &linkedlist{},
	}
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestLinkedlist_AddFirst(t *testing.T) {

	list.list.AddFirst(6)
	val := list.list.Get(6)
	if val != 6 {
		t.Error("Expected 6, got ", val)
	}

}

func TestLinkedlist_AddLast(t *testing.T) {
	list.list.AddLast(6)
	list.list.AddLast(7)
	a := list.list.toArray()
	arr := []int{6, 7}
	for i := 0; i < len(arr); i++ {
		if arr[i] != a[i] {
			t.Error("Expected", arr[i], ", got ", a[i])
		}
	}
}
