package fastexp

import (
	"math/bits"
)

// Find : returns y = x^(e) mod n
func Find(x, e, n uint64) uint64 {
	r := x
	for i := bits.Len64(e) - 2; i >= 0; i-- {
		r = (r * r) % n
		if (e>>i)&1 == 1 {
			r = (r * x) % n
		}
	}
	return r
}
