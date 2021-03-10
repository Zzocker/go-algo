// Package heap : heap implantation of max heap
// https://en.wikipedia.org/wiki/Heap_(data_structure)
package heap

// Heap represents a max heap
type Heap struct {
	bucket []int
	size   int
}

// New : Creates a new max heap with a given capacity
func New(capacity int) *Heap {
	return &Heap{
		size:   0,
		bucket: make([]int, capacity),
	}
}

// Insert will insert a new elem in heap
func (h *Heap) Insert(item int) {
	h.bucket[h.size] = item
	h.shiftUp(h.size)
	h.size++
}

func (h *Heap) shiftUp(index int) {
	i := index
	for i > 0 && h.bucket[i] > h.bucket[(i-1)/2] {
		h.bucket[i], h.bucket[(i-1)/2] = h.bucket[(i-1)/2], h.bucket[i]
		i = (i - 1) / 2
	}
}

// Max : give max element
func (h *Heap) Max() int { return h.bucket[0] }

// Size :
func (h *Heap) Size() int { return h.size }

// IsEmpty : true if heap is empty
// false otherwise
func (h *Heap) IsEmpty() bool { return h.size == 0 }

// ExtractMax : removes max element from the heap and return it
func (h *Heap) ExtractMax() int {
	out := h.bucket[0]
	h.bucket[0], h.bucket[h.size-1] = h.bucket[h.size-1], h.bucket[0]
	h.size--
	h.shiftDown(0)
	return out
}

func (h *Heap) shiftDown(index int) {
	j := index
	for j < h.size {
		var tmp int
		if 2*j+2 < h.size {
			tmp = 2*j + 1
			if h.bucket[2*j+2] > h.bucket[2*j+1] {
				tmp = 2*j + 2
			}
		} else if 2*j+1 < h.size {
			tmp = 2*j + 1
		} else {
			break
		}

		if h.bucket[tmp] > h.bucket[j] {
			h.bucket[j], h.bucket[tmp] = h.bucket[tmp], h.bucket[j]
			j = tmp
		} else {
			break
		}
	}
}

// Remove item at index i
func (h *Heap) Remove(i int) {
	h.bucket[i] = h.bucket[h.size-1]
	h.size--
	h.shiftDown(i)
}

// Heapify : converters a array into a heap
func Heapify(arr []int) *Heap {
	h := Heap{
		size:   len(arr),
		bucket: arr,
	}
	for i := h.size - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
	return &h
}

// Sort : will sort given array using heap sort
func Sort(arr []int) {
	h := Heapify(arr)
	for i := 0; i < len(arr); i++ {
		h.ExtractMax()
	}
}
