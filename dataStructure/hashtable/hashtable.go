// Package hashtable is Has table implementation using chaining and openaddressing
package hashtable

import (
	"fmt"
	"math"
)

// Table represents a has table
type Table interface {
	// Put : insert a new entry in hash table with given key
	// if key already exists then value that key is updated
	Put(interface{}, interface{})

	// Get : returns value of a given key
	// for non-existing key nil is returned
	Get(interface{}) interface{}

	Delete(interface{})
	Exists(interface{}) bool

	Len() int
}

type entry struct {
	key interface{}
	val interface{}
}

const (
	radix       int     = '~' - ' ' + 1
	goldenRatio float64 = 1.618033988749895
)

func encode(key interface{}) int {
	var out int
	for _, v := range fmt.Sprintf("%v", key) {
		out = int(v) + radix*out
	}
	return out
}

// multiplication based hashing
func hash(encoded int, cap int) int {
	ge := float64(encoded) * goldenRatio
	return int((ge - math.Floor(ge)) * float64(cap))
}
