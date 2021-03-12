package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	size := 10
	want := make([]int, 0, size)
	a := make([]int, 0, size)
	for i := size - 1; i >= 0; i-- {
		a = append(a, i)
		want = append(want, size-i-1)
	}
	HeapSort(a)
	assert.Equal(t, want, a)
}

func TestMergeSort(t *testing.T) {
	size := 10
	want := make([]int, 0, size)
	a := make([]int, 0, size)
	for i := size - 1; i >= 0; i-- {
		a = append(a, i)
		want = append(want, size-i-1)
	}
	MergeSort(a)
	assert.Equal(t, want, a)
}

func TestBottomUpMergeSort(t *testing.T) {
	size := 10
	want := make([]int, 0, size)
	a := make([]int, 0, size)
	for i := size - 1; i >= 0; i-- {
		a = append(a, i)
		want = append(want, size-i-1)
	}
	BottomUpMergeSort(a)
	assert.Equal(t, want, a)
}
