package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	a := New()
	assert.Zero(t, a.Len())
	a.Append("12")
	a.Append("56")
	assert.Equal(t, 2, a.Len())
}

func TestCap(t *testing.T) {
	a := New()

	assert.Equal(t, 1, a.Cap())
	a.Append("123")
	assert.Equal(t, 1, a.Cap())
	a.Append("236")
	assert.Equal(t, 2, a.Cap())
	a.Append("26")
	assert.Equal(t, 4, a.Cap())
	a.Append("246")
	assert.Equal(t, 4, a.Cap())
}

func TestAppend(t *testing.T) {
	a := New()

	a.Append("15")
	assert.Equal(t, "15", a.Get(0))
	a.Append("16")
	a.Append("17")
	a.Append("18")
	assert.Equal(t, "18", a.Get(a.Len()-1))
}

func TestPop(t *testing.T) {
	a := New()

	assert.Nil(t, a.Pop())

	a.Append(1)
	a.Append(2)
	assert.Equal(t, 2, a.Pop())
	assert.Equal(t, 1, a.Pop())

}

func TestInsert(t *testing.T) {
	a := New()

	assert.Error(t, a.Insert(1, "15"))

	a.Append(1)
	a.Append(2)
	assert.NoError(t, a.Insert(1, 456))
	assert.Equal(t, 456, a.Get(1))
}

func TestGet(t *testing.T) {
	a := New()

	size := 50
	for i := 0; i < size; i++ {
		a.Append(i)
	}

	for i := 0; i < size; i++ {
		assert.Equal(t, i, a.Get(i))
	}

}

func BenchmarkArrays(b *testing.B) {
	a := New()

	for i := 0; i < b.N; i++ {
		a.Append(i)
	}
}

func BenchmarkGo(b *testing.B) {
	a := make([]interface{}, 1)

	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
	_ = a
}
