// Package stack implements bounded and ub-bounded stack
// bounded stack is implemented using array
// un-bounded is implemented using singly linkedlist
package stack

import (
	"container/list"
)

// UnBound represents linked list implementation of stack
type UnBound interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	Len() int
	Empty() bool
}

type stk struct {
	ld *list.List
}

// NewUnBound create unbounded stack
func NewUnBound() UnBound {
	return stk{
		ld: list.New(),
	}
}

func (s stk) Push(val interface{}) {
	s.ld.PushFront(val)
}

func (s stk) Pop() interface{} {
	front := s.ld.Front()
	if front == nil {
		return nil
	}
	s.ld.Remove(front)
	return front.Value
}

func (s stk) Top() interface{} {
	front := s.ld.Front()
	if front == nil {
		return nil
	}
	return front.Value
}

func (s stk) Len() int {
	return s.ld.Len()
}

func (s stk) Empty() bool {
	return s.ld.Len() == 0
}
