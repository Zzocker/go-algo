package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainingPut(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	tb.Put("key", "version1")
	is.Equal(1, tb.Len())
	is.Equal("version1", tb.Get("key"))

	tb.Put("key", "version2")
	is.Equal(1, tb.Len())
	is.Equal("version2", tb.Get("key"))

	tb.Put("key1", "version1")
	is.Equal(2, tb.Len())
	is.Equal("version1", tb.Get("key1"))

}

func TestChainingGet(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	tb.Put("key", "version1")
	is.Equal("version1", tb.Get("key"))

	is.Nil(tb.Get("key-not-found"))
}

func TestChainingDelete(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	tb.Put("key", "version1")
	tb.Delete("key")
	is.Nil(tb.Get("key"))

	tb.Delete("key-no-found")
}

func TestChainingLen(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	is.Zero(tb.Len())
	tb.Put("key", "version1")
	is.Equal(1, tb.Len())

	tb.Delete("key")
	is.Zero(tb.Len())
}

func TestChainingExists(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	is.False(tb.Exists("key-not-found"))
	tb.Put("key", nil)
	is.True(tb.Exists("key"))
	tb.Delete("key")
	is.False(tb.Exists("key"))
}

func BenchmarkChaining(b *testing.B) {
	tb := NewCTable()
	for i := 0; i < b.N; i++ {
		tb.Put(i, nil)
	}
}

func BenchmarkGoMap(b *testing.B) {
	m := make(map[interface{}]interface{}, 2)
	for i := 0; i < b.N; i++ {
		m[i] = nil
	}
}
