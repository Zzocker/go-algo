package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	l := New()
	is := assert.New(t)

	is.Zero(l.Size())
	n1 := NewNode(1)
	l.PushFront(n1)
	is.Equal(1, l.Size())

	n2 := NewNode(1)
	l.PushFront(n2)
	is.Equal(2, l.Size())

	l.PopBack()
	is.Equal(1, l.Size())
}

func TestEmpty(t *testing.T) {
	l := New()
	is := assert.New(t)

	is.True(l.Empty())

	n1 := NewNode(1)
	l.PushFront(n1)
	is.False(l.Empty())

	l.PopBack()
	is.True(l.Empty())

}

func TestNodeAt(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	l.PushFront(n1)
	l.PushFront(n2)
	l.PushFront(n3)
	is.Equal(n2.Value, l.NodeAt(1).Value)
	is.Equal(n3.Value, l.NodeAt(0).Value)
	is.Equal(n1.Value, l.NodeAt(2).Value)
	is.Nil(l.NodeAt(5))

}

func TestPushFront(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushFront(n1)
	is.Equal(n1.Value, l.NodeAt(0).Value)

	n2 := NewNode(2)
	l.PushFront(n2)

	is.Equal(n2.Value, l.NodeAt(0).Value)

	n3 := NewNode(3)
	l.PushFront(n3)
	is.Equal(n3.Value, l.NodeAt(0).Value)

}

func TestPopFront(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushFront(n1)
	is.Equal(n1.Value, l.PopFront())
	is.Nil(l.NodeAt(0))
	is.Zero(l.Size())
}

func TestPushBack(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushBack(n1)
	is.Equal(n1.Value, l.NodeAt(0).Value)

	n2 := NewNode(2)
	l.PushBack(n2)

	is.Equal(n2.Value, l.NodeAt(1).Value)

	n3 := NewNode(3)
	l.PushBack(n3)
	is.Equal(n3.Value, l.NodeAt(2).Value)

}

func TestPopBack(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushFront(n1)
	n2 := NewNode(2)
	l.PushFront(n2)
	is.Equal(n1.Value, l.PopBack())
	is.Equal(1, l.Size())

	is.Equal(n2.Value, l.PopBack())
	is.Zero(l.Size())
}

func TestFront(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushFront(n1)
	is.Equal(n1.Value, l.Front().Value)
	n2 := NewNode(2)
	l.PushFront(n2)
	is.Equal(n2.Value, l.Front().Value)
	n3 := NewNode(3)
	l.PushBack(n3)
	is.Equal(n2.Value, l.Front().Value)

}

func TestBack(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushFront(n1)
	is.Equal(n1.Value, l.Back().Value)
	n2 := NewNode(2)
	l.PushFront(n2)
	is.Equal(n1.Value, l.Back().Value)

	n3 := NewNode(3)
	l.PushBack(n3)
	is.Equal(n3.Value, l.Back().Value)
}

func TestInsert(t *testing.T) {
	l := New()
	is := assert.New(t)

	l.Insert(1, nil)
	is.Zero(l.Size())
	size := 10
	notFound := 5
	for i := size - 1; i >= 0; i-- {
		if i != notFound {
			l.PushFront(NewNode(i))
		}
	}
	l.Insert(notFound, NewNode(notFound))
	is.Equal(size, l.Size())
	is.Equal(notFound, l.NodeAt(notFound).Value)
	is.Equal(notFound-1, l.NodeAt(notFound-1).Value)
	is.Equal(notFound+1, l.NodeAt(notFound+1).Value)

}

func TestErase(t *testing.T) {
	l := New()
	is := assert.New(t)

	l.Insert(1, nil)
	is.Zero(l.Size())
	size := 10
	for i := size - 1; i >= 0; i-- {
		l.PushFront(NewNode(i))
	} // 0 1 2 3 4 6 7 8 9
	removeAt := 5
	l.Erase(removeAt)
	is.Equal(size-1, l.Size())
	is.Equal(removeAt+1, l.NodeAt(removeAt).Value)
	is.Equal(removeAt-1, l.NodeAt(removeAt-1).Value)
}

func TestFromEnd(t *testing.T) {
	l := New()
	is := assert.New(t)

	n1 := NewNode(1)
	l.PushFront(n1)
	n2 := NewNode(2)
	l.PushFront(n2)
	n3 := NewNode(3)
	l.PushBack(n3)
	// 2 1 3
	is.Equal(n3.Value, l.FromEnd(1).Value)
	is.Equal(n1.Value, l.FromEnd(2).Value)
	is.Equal(n2.Value, l.FromEnd(3).Value)

}

func TestReverse(t *testing.T) {
	l := New()
	is := assert.New(t)

	size := 10
	for i := 0; i < size; i++ {
		l.PushFront(NewNode(i))
	}
	l.Reverse()
	ptr := l.NodeAt(0)
	for i := 0; i < size; i++ {
		is.Equal(i, ptr.Value)
		ptr = ptr.Next()
	}
}
