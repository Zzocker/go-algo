package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	is := assert.New(t)
	s := NewUnBound()

	s.Push(1)
	is.Equal(1, s.Len())
	is.Equal(1, s.Top())

	s = NewUnBound()
	size := 16
	for i := 0; i < size; i++ {
		s.Push(i)
		is.Equal(i+1, s.Len())
		is.Equal(i, s.Top())
	}
}

func TestPop(t *testing.T) {
	is := assert.New(t)
	s := NewUnBound()

	is.Nil(s.Pop())

	s.Push(1)
	is.Equal(1, s.Pop())
	is.Zero(s.Len())

	size := 16
	for i := 0; i < size; i++ {
		s.Push(i)
	}
	for i := size - 1; i >= 0; i-- {
		is.Equal(i, s.Pop())
		is.Equal(i, s.Len())
	}
}

func TestTop(t *testing.T) {
	is := assert.New(t)
	s := NewUnBound()

	is.Nil(s.Top())

	s.Push(1)
	is.Equal(1, s.Top())
}

func TestLen(t *testing.T) {
	is := assert.New(t)
	s := NewUnBound()

	is.Zero(s.Len())

	size := 16
	for i := 0; i < size; i++ {
		s.Push(5)
		is.Equal(i+1, s.Len())
	}
}

func TestEmpty(t *testing.T) {
	is := assert.New(t)
	s := NewUnBound()

	is.True(s.Empty())

	s.Push(1)
	is.False(s.Empty())

	s.Pop()
	is.True(s.Empty())
}

func BenchmarkPush(b *testing.B) {
	s := NewUnBound()
	for i := 0; i < b.N; i++ {
		s.Push(1)
	}
}
