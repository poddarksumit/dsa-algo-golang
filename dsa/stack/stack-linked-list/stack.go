package stack_linked_list

import "errors"

type (
	StackTypes interface {
		int | float64 | string | rune | byte
	}

	Node[T StackTypes] struct {
		value T
		next  *Node[T]
	}

	Stack[T StackTypes] struct {
		head   *Node[T]
		length int
	}
)

func new[T StackTypes](newValue T) *Node[T] {
	return &Node[T]{
		value: newValue,
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var result T
		return result, errors.New("empty stack")
	}

	return s.head.value, nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var result T
		return result, errors.New("empty stack")
	}

	tempHead := s.head
	s.head = s.head.next
	s.length--

	return tempHead.value, nil
}

func (s *Stack[T]) Push(newValue T) {
	newNode := new(newValue)

	if s.IsEmpty() {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}

	s.length++
}
