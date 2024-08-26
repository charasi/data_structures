package queue

import (
	"log"
	"os"
	"testing"
)

var q queueInterface

func TestMain(m *testing.M) {
	q = NewQueue(5)
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestQueue_Enqueue(t *testing.T) {
	err := q.Enqueue(7)
	err = q.Enqueue(8)
	err = q.Enqueue(9)
	err = q.Enqueue(10)
	err = q.Enqueue(11)
	if err != nil {
		t.Error(err)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	err := q.Enqueue(7)
	err = q.Enqueue(8)
	err = q.Enqueue(9)
	err = q.Enqueue(10)
	err = q.Enqueue(11)
	_, err = q.Dequeue()
	err = q.Enqueue(12)
	if err != nil {
		return
	}
	if err != nil {
		t.Error(err)
	}
}
