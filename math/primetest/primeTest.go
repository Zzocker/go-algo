package primetest

import (
	"math/rand"
	"time"

	"github.com/Zzocker/go-algo/math/fastexp"
)

// Check : check if input p is prime or not
// Input : p (prime-candidate), s security parameter
func Check(p uint64, s int) bool {
	if p != 2 && p&1 == 0 {
		return false
	}
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	u, r := decompose(p - 1)
	for i := 0; i < s; i++ { // pass
		// choose a from [2,p-2]
		a := ran.Uint64() % (p - 1) // use crypto/rand
		if a > 1 {
			z := fastexp.Find(a, r, p)
			if z == 1 || z == p-1 {
				continue
			}
			var flag bool
			for j := uint64(0); j < u-1; j++ {
				z = (z * z) % p
				if z == p-1 {
					flag = true
					break
				}
			}
			if !flag {
				return false
			}
		}
	}
	return true
}

// decompose p in form p = 2^(u)*r
// here r is a mod number
func decompose(p uint64) (u uint64, r uint64) {
	for ; p&1 == 0; u++ {
		p = p >> 1
	}
	r = p
	return
}
