package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	is := assert.New(t)
	a := []int{20, 15, 30, 10, 15, 13, 40}
	h := New(len(a))

	is.Zero(h.Size())
	for _, v := range a {
		h.Insert(v)
	}

	is.Equal(len(a), h.Size())
	is.Equal(40, h.Max())
}

func TestExtractMax(t *testing.T) {
	is := assert.New(t)
	a := []int{20, 15, 30, 10, 15, 13, 40}
	h := New(len(a))

	is.Zero(h.Size())
	for _, v := range a {
		h.Insert(v)
	}

	is.Equal(40, h.ExtractMax())
	is.Equal(len(a)-1, h.Size())

	is.Equal(30, h.ExtractMax())
	is.Equal(len(a)-2, h.Size())
}

func TestRemove(t *testing.T) {
	is := assert.New(t)
	a := []int{20, 15, 30, 10, 15, 13, 40}
	h := New(len(a))

	is.Zero(h.Size())
	for _, v := range a {
		h.Insert(v)
	}

	// removing element 30 which is at index 2
	// [40 15 30 10 15 13 20]
	h.Remove(2)
	is.Equal(len(a)-1, h.Size())
	h.ExtractMax()

	is.Equal(20, h.Max())
}

func TestHeapify(t *testing.T) {
	is := assert.New(t)
	a := []int{20, 15, 30, 10, 15, 13, 40}
	sorted := []int{40, 30, 20, 15, 15, 13, 10}
	h := Heapify(a)

	for _, v := range sorted {
		is.Equal(v, h.ExtractMax())
	}
}

func TestSort(t *testing.T) {
	is := assert.New(t)
	a := []int{20, 15, 30, 10, 15, 13, 40}
	want := []int{10, 13, 15, 15, 20, 30, 40}

	Sort(a)
	is.Equal(want, a)
}
