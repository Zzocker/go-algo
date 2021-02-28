package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnboundPush(t *testing.T) {
	is := assert.New(t)
	q := NewUnbound()

	q.Push(1)
	is.Equal(1, q.Len())
	is.Equal(1, q.Top())

	q.Push(2)
	is.Equal(1, q.Top())
}

func TestUnboundPop(t *testing.T) {
	is := assert.New(t)
	q := NewUnbound()

	is.Nil(q.Pop())

	q.Push(1)
	q.Push(2)

	is.Equal(1, q.Pop())
	is.Equal(1, q.Len())

	is.Equal(2, q.Pop())
	is.Zero(q.Top())
}

func TestUnboundTop(t *testing.T) {
	is := assert.New(t)
	q := NewUnbound()

	is.Nil(q.Top())

	q.Push(1)
	q.Push(2)
	is.Equal(1, q.Top())

	q.Pop()
	is.Equal(2, q.Top())
}

func TestUnboundLen(t *testing.T) {
	is := assert.New(t)
	q := NewUnbound()

	is.Zero(q.Len())

	q.Push(1)
	is.Equal(1, q.Len())

	q.Push(1)
	is.Equal(2, q.Len())

	q.Pop()
	is.Equal(1, q.Len())
}

func TestUnboundEmpty(t *testing.T) {
	is := assert.New(t)
	q := NewUnbound()

	is.True(q.Empty())
	q.Push(1)
	is.False(q.Empty())
	q.Pop()
	is.True(q.Empty())
}

func BenchmarkUnboundPush(b *testing.B) {
	s := NewUnbound()
	for i := 0; i < b.N; i++ {
		s.Push(1)
	}
}
