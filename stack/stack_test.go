package stack

import "testing"

func TestIsFull(t *testing.T) {
	s := New(2)
	s.Push(10)
	s.Push(20)
	if !s.IsFull() {
		t.Errorf("Expected true, got false")
	}
}

func TestIsEmpty(t *testing.T) {
	s := New(1)
	if !s.IsEmpty() {
		t.Errorf("Expected true, got false")
	}
}

func TestPush(t *testing.T) {
	expectedLastElement := 10
	s := New(1)
	s.Push(expectedLastElement)

	lastElement := s.Peek()

	if *lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, *lastElement)
	}
}

func TestPushFullStack(t *testing.T) {
	s := New(1)
	s.Push(10)
	err := s.Push(20)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPop(t *testing.T) {
	expectedLastElement := 10
	initialCapacity := 1
	expectedCapacity := initialCapacity - 1
	s := New(initialCapacity)
	s.Push(expectedLastElement)

	lastElement, _ := s.Pop()

	if *lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, *lastElement)
	} else if s.Capacity != expectedCapacity {
		t.Errorf("Expected capacity %d, got %d", expectedCapacity, s.Size())
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := New(1)
	_, err := s.Pop()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSize(t *testing.T) {
	expectedSize := 2
	s := New(expectedSize)
	s.Push(10)
	s.Push(20)

	if s.Size() != expectedSize {
		t.Errorf("Expected %d, got %d", expectedSize, s.Size())
	}
}

func TestPeek(t *testing.T) {
	expectedLastElement := 10
	s := New(1)
	s.Push(expectedLastElement)

	lastElement := s.Peek()

	if *lastElement != expectedLastElement {
		t.Errorf("Expected %v, got %v", expectedLastElement, *lastElement)
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := New(1)
	lastElement := s.Peek()

	if lastElement != nil {
		t.Errorf("Expected nil, got %v", *lastElement)
	}
}