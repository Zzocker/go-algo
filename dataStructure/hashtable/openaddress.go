package hashtable

const (
	deltedKey      = '~'
	oMaxLoadFactor = 0.67
)

type oTable struct {
	buckets  []*entry
	size     int
	m        int
	capacity int
}

// NewOTable create a new hash table implemented by openaddressing
func NewOTable() Table {
	return &oTable{
		buckets:  make([]*entry, 2),
		size:     0,
		m:        1,
		capacity: 2,
	}
}
func (t *oTable) Put(key interface{}, val interface{}) {
	if t.size >= int(float64(t.capacity)*oMaxLoadFactor) {
		t.resize(t.m + 1)
	}
	encoded := encode(key)
	hashed := hash(encoded, t.capacity)
	trail := 1
	index := hashed
	for ; t.buckets[index] != nil && t.buckets[index].key == key; trail++ {
		index = (hashed + trail*hash2(encoded, t.capacity)) % t.capacity
	}
	if t.buckets[index] == nil {
		t.buckets[index] = &entry{key: key, val: val}
		t.size++
	} else {
		t.buckets[index].val = val
	}
}
func (t *oTable) Get(key interface{}) interface{} {
	encoded := encode(key)
	hashed := hash(encoded, t.capacity)
	trail := 1
	index := hashed
	for ; t.buckets[index] != nil; trail++ {
		if t.buckets[index].key == key {
			return t.buckets[index].val
		}
		index = (hashed + trail*hash2(encoded, t.capacity)) % t.capacity
	}
	return nil
}
func (t *oTable) Delete(key interface{}) {
	encoded := encode(key)
	hashed := hash(encoded, t.capacity)
	trail := 1
	index := hashed
	for ; t.buckets[index] != nil; trail++ {
		if t.buckets[index].key == key {
			t.buckets[index].key = deltedKey
			t.buckets[index].val = nil
			t.size--
		}
		index = (hashed + trail*hash2(encoded, t.capacity)) % t.capacity
	}

	if t.size <= t.capacity/4 {
		t.resize(t.m - 1)
	}
}
func (t *oTable) Exists(key interface{}) bool {
	return t.Get(key) != nil
}
func (t *oTable) Len() int {
	return t.size
}

func (t *oTable) resize(newM int) {
	t.m = newM
	t.capacity = 1 << newM
	newBucket := make([]*entry, t.capacity)
	var encoded int
	var hashed int
	var trail int
	var index int
	for _, v := range t.buckets {
		if v != nil && v.key != deltedKey {
			encoded = encode(v.key)
			hashed = hash(encoded, t.capacity)
			index = hashed
			trail = 1
			for ; newBucket[index] != nil; trail++ {
				index = (hashed + trail*hash2(encoded, t.capacity)) % t.capacity
			}
			newBucket[index] = v
		}
	}
	t.buckets = newBucket
}
