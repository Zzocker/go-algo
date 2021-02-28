package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundLen(t *testing.T) {
	is := assert.New(t)
	q := NewBound(4)

	is.Zero(q.Len())
	q.Push(1)
	is.Equal(1, q.Len())
	q.Push(1)
	is.Equal(2, q.Len())

	q.Pop()
	is.Equal(1, q.Len())
}

func TestBoundEmpty(t *testing.T) {
	is := assert.New(t)
	q := NewBound(4)
	is.True(q.Empty())

	q.Push(1)
	is.False(q.Empty())

	q.Pop()
	is.True(q.Empty())
}

func TestBoundFull(t *testing.T) {
	is := assert.New(t)
	q := NewBound(3)

	is.False(q.Full())

	q.Push(1)
	q.Push(2)
	q.Push(3)
	is.True(q.Full())

	q.Pop()
	is.False(q.Full())
}

func TestBoundPush(t *testing.T) {
	is := assert.New(t)
	q := NewBound(4)

	is.NoError(q.Push(1))
	is.Equal(1, q.Top())
	is.Equal(1, q.Len())

	is.NoError(q.Push(2))
	is.Equal(1, q.Top())
	is.Equal(2, q.Len())

	is.NoError(q.Push(3))
	is.NoError(q.Push(4))

	is.Error(q.Push(5))

	q.Pop()
	is.Equal(2, q.Top())

	is.NoError(q.Push(5))
}

func TestBoundPop(t *testing.T) {
	is := assert.New(t)
	q := NewBound(2)

	is.Nil(q.Pop())

	q.Push(1)
	is.Equal(1, q.Pop())
	is.Zero(q.Len())

	q.Push(1)
	q.Push(2)

	is.Equal(1, q.Pop())
	is.Equal(1, q.Len())
	is.Equal(2, q.Pop())
	is.Zero(q.Len())
}

func TestBoundTop(t *testing.T) {
	is := assert.New(t)
	q := NewBound(4)

	is.Nil(q.Top())

	q.Push(1)
	is.Equal(1, q.Top())

	q.Push(2)
	q.Pop()
	is.Equal(2, q.Top())
}
