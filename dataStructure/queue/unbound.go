// Package queue implements Queue (FIFO)
// unbound stack is implemented using linked list
package queue

import "container/list"

// UnBound represents stack has no size limit
type UnBound interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	Len() int
	Empty() bool
}

type unbound struct {
	ld *list.List
}

// NewUnbound creates a new unbound queue
func NewUnbound() UnBound {
	return unbound{
		ld: list.New(),
	}
}
func (u unbound) Push(v interface{}) {
	u.ld.PushFront(v)
}

func (u unbound) Pop() interface{} {
	first := u.ld.Back()
	if first == nil {
		return nil
	}
	u.ld.Remove(first)
	return first.Value
}

func (u unbound) Top() interface{} {
	first := u.ld.Back()
	if first == nil {
		return nil
	}
	return first.Value
}

func (u unbound) Len() int    { return u.ld.Len() }
func (u unbound) Empty() bool { return u.ld.Len() == 0 }
