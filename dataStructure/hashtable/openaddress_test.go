package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAddressPut(t *testing.T) {
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

func TestOpenAddressGet(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	tb.Put("key", "version1")
	is.Equal("version1", tb.Get("key"))

	is.Nil(tb.Get("key-not-found"))
}

func TestOpenAddressDelete(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	tb.Put("key", "version1")
	tb.Delete("key")
	is.Nil(tb.Get("key"))

	tb.Delete("key-no-found")
}

func TestOpenAddressLen(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	is.Zero(tb.Len())
	tb.Put("key", "version1")
	is.Equal(1, tb.Len())

	tb.Delete("key")
	is.Zero(tb.Len())
}

func TestOpenAddressExists(t *testing.T) {
	is := assert.New(t)
	tb := NewCTable()

	is.False(tb.Exists("key-not-found"))
	tb.Put("key", nil)
	is.True(tb.Exists("key"))
	tb.Delete("key")
	is.False(tb.Exists("key"))
}

func BenchmarkOpenAddress(b *testing.B) {
	tb := NewCTable()
	for i := 0; i < b.N; i++ {
		tb.Put(i, nil)
	}
}
