package hashtable

import (
	"container/list"
)

const (
	cMaxLoadFactor = 0.67
)

type cTable struct {
	buckets  []*list.List
	size     int
	m        int
	capacity int // capacity = 2^m
}

// NewCTable : creates a new hash table for which collision is resolved
// by chaining
func NewCTable() Table {
	return &cTable{
		buckets:  make([]*list.List, 2),
		m:        1,
		capacity: 2,
	}
}

func (t *cTable) Len() int { return t.size }
func (t *cTable) Put(key, val interface{}) {
	if t.size > int(float64(t.capacity)*cMaxLoadFactor) {
		t.resize(t.m + 1)
	}
	index := hash(encode(key), t.capacity)
	bucket := t.buckets[index]
	if bucket == nil {
		bucket = list.New()
		t.buckets[index] = bucket
	}
	var ptr *list.Element = bucket.Front()
	for ; ptr != nil && ptr.Value.(*entry).key != key; ptr = ptr.Next() {
	}

	if ptr == nil {
		bucket.PushFront(&entry{key: key, val: val})
		t.size++
	} else {
		ptr.Value.(*entry).val = val
	}
}

func (t *cTable) Get(key interface{}) interface{} {
	index := hash(encode(key), t.capacity)
	bucket := t.buckets[index]
	if bucket == nil {
		return nil
	}
	var ptr *list.Element = bucket.Front()
	for ; ptr != nil; ptr = ptr.Next() {
		if ptr.Value.(*entry).key == key {
			return ptr.Value.(*entry).val
		}
	}
	return nil
}

func (t *cTable) Delete(key interface{}) {
	index := hash(encode(key), t.capacity)
	bucket := t.buckets[index]
	if bucket == nil {
		return
	}
	var ptr *list.Element = bucket.Front()
	for ; ptr != nil; ptr = ptr.Next() {
		if ptr.Value.(*entry).key == key {
			bucket.Remove(ptr)
			t.size--
			if t.size <= t.capacity>>2 {
				t.resize(t.m - 1)
			}
			return
		}
	}
}

func (t *cTable) Exists(key interface{}) bool {
	index := hash(encode(key), t.capacity)
	bucket := t.buckets[index]
	if bucket == nil {
		return false
	}
	var ptr *list.Element = bucket.Front()
	for ; ptr != nil; ptr = ptr.Next() {
		if ptr.Value.(*entry).key == key {
			return true
		}
	}
	return false
}

func (t *cTable) resize(newM int) {
	t.m = newM
	t.capacity = 1 << newM
	newBuckets := make([]*list.List, t.capacity)
	var index int
	var ptr *list.Element
	for _, v := range t.buckets {
		if v != nil {
			ptr = v.Front()
			for ; ptr != nil; ptr = ptr.Next() {
				e := ptr.Value.(*entry)
				index = hash(encode(e.key), t.capacity)
				bucket := newBuckets[index]
				if bucket == nil {
					bucket = list.New()
				}
				bucket.PushFront(e)
			}
		}
	}
	t.buckets = newBuckets
}
