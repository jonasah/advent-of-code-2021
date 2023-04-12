package container

import (
	"fmt"
)

type Stack[T any] interface {
	Push(val T)
	Pop() (*T, error)
	Len() int
}

type stack[T any] struct {
	values []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{values: make([]T, 0)}
}

func (s *stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *stack[T]) Pop() (*T, error) {
	l := len(s.values)

	if l == 0 {
		return nil, fmt.Errorf("empty stack")
	}

	val := s.values[l-1]
	s.values = s.values[:l-1]
	return &val, nil
}

func (s *stack[T]) Len() int {
	return len(s.values)
}
