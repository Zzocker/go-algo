// Package dfk is simple demonstration of Diffie–Hellman key exchange
// over a prime cyclic group (Zp,*)
package dfk

import (
	"math/rand"
	"time"

	"github.com/Zzocker/go-algo/math/fastexp"
)

const (
	//public domain paramenters
	// p : prime number
	p uint64 = 11

	// a : generator or primitive element of the group
	a uint64 = 2
)

func init() { rand.Seed(time.Now().UnixNano()) }

// GeneratePrivateKey : returns a randome number from the set {2,3,..p-1}
func GeneratePrivateKey() uint64 {
	var key uint64
	for {
		key = rand.Uint64() % p
		if p > 1 {
			break
		}
	}
	return key
}

// GeneratePublicKey : alpha^(prvKey) mod p
func GeneratePublicKey(prvKey uint64) uint64 {
	return fastexp.Find(a, prvKey, p)
}

// GenerateSessionKey : A^(prvKey) mod p
func GenerateSessionKey(A, prvKey uint64) uint64 {
	return fastexp.Find(A, prvKey, p)
}
