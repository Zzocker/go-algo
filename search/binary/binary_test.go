package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoopBinary(t *testing.T) {
	size := 100
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	for i := 0; i < size; i++ {
		assert.Equal(t, i, LoopSearch(arr, i))
	}
}

func TestRecursionBinary(t *testing.T) {
	size := 100
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	for i := 0; i < size; i++ {
		assert.Equal(t, i, RecursionSearch(arr, i))
	}
}

var (
	size = 100000000
	arr  = make([]int, size)
)

func BenchmarkLoopBinary(b *testing.B) {
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	for i := 0; i < b.N; i++ {
		LoopSearch(arr, i)
	}
}

func BenchmarkRecursionBinary(b *testing.B) {
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	for i := 0; i < b.N; i++ {
		RecursionSearch(arr, i)
	}
}
