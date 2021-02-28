package queue

import "errors"

// Bound represents bounded queue
// which has max capacity limit
type Bound interface {
	Push(interface{}) error
	Pop() interface{}
	Top() interface{}
	Len() int
	Empty() bool
	Full() bool
}

// ErrQueueFull : returned when queue is full
var ErrQueueFull = errors.New("queue is full")

type bound struct {
	arr   []interface{}
	size  int
	cap   int
	read  int
	write int
}

// NewBound : returns a bounded queue with given capacity
func NewBound(cap int) Bound {
	return &bound{
		arr: make([]interface{}, cap),
		cap: cap,
	}
}
func (b *bound) Len() int    { return b.size }
func (b *bound) Empty() bool { return b.size == 0 }
func (b *bound) Full() bool  { return b.size == b.cap }

func (b *bound) Push(val interface{}) error {
	if b.Full() {
		return ErrQueueFull
	}
	b.arr[b.write] = val
	b.size++
	b.write = (b.write + 1) % b.cap
	return nil
}

func (b *bound) Pop() interface{} {
	if b.Empty() {
		return nil
	}
	first := b.arr[b.read]
	b.arr[b.read] = nil
	b.size--
	b.read = (b.read + 1) % b.cap
	return first
}

func (b *bound) Top() interface{} {
	if b.Empty() {
		return nil
	}
	return b.arr[b.read]
}
