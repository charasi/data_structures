package stack

import (
	"log"
	"os"
	"testing"
)

var s stackInterface

func TestMain(m *testing.M) {
	s = NewStack(10)
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestStack_Push(t *testing.T) {
	err := s.Push(1)
	if err != nil {
		t.Error("Error pushing 1")
	}
}

func TestStack_Pop(t *testing.T) {
	err := s.Push(1)
	err = s.Push(2)
	err = s.Push(3)
	err = s.Push(4)
	_, err = s.Pop()
	if err != nil {
		t.Error("Error popping 1")
	}
}
