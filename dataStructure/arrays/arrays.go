// Package arrays : implements a dynamic array
package arrays

import "errors"

// Array represents methods that are supported by this dynamic array implantation
type Array interface {
	Len() int
	Cap() int
	// Append : insert a new element at the end of array
	Append(interface{})

	// Pop : removes last element of the array
	Pop() interface{}

	// Insert : will updated the element at given index
	// if index is wrong then invalid index error will be returned
	Insert(int, interface{}) error

	// Get : will return the element present at given index
	// nil is returned if index is invalid
	Get(int) interface{}
}

// New : creates a new dynamic array
func New() Array {
	return &a{
		base:     make([]interface{}, initCapacity),
		capacity: initCapacity,
	}
}

const (
	initCapacity int = 1
)

type a struct {
	base     []interface{}
	size     int
	capacity int
}

func (a *a) Len() int {
	return a.size
}
func (a *a) Cap() int {
	return a.capacity
}
func (a *a) Append(elm interface{}) {
	if a.size >= a.capacity {
		a.resize(2 * a.capacity)
	}
	a.base[a.size] = elm
	a.size++
}
func (a *a) Pop() interface{} {
	if a.size == 0 {
		return nil
	}
	a.size--
	out := a.base[a.size]
	a.base[a.size] = nil
	if a.size < a.capacity/4 {
		a.resize(a.capacity / 2)
	}
	return out
}
func (a *a) Insert(index int, elm interface{}) error {
	if index < 0 || index >= a.size {
		return errors.New("invalid index")
	}
	a.base[index] = elm
	return nil
}
func (a *a) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		return nil
	}
	return a.base[index]
}

func (a *a) resize(newCap int) {
	a.capacity = newCap
	newBase := make([]interface{}, newCap)
	for i := 0; i < a.size; i++ {
		newBase[i] = a.base[i]
	}
	a.base = nil
	a.base = newBase
}
