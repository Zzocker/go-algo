// Package rsa : is a naive RSA public key cryptography implication
// keys are uint64 (i.e k < 2^(64)-1), also not recommended, insted use math/big.Ints
// also encryption and decryption are done on a uint64 number (>_<)
package rsa

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Zzocker/go-algo/math/fastexp"
	"github.com/Zzocker/go-algo/math/modularinverse"
	"github.com/Zzocker/go-algo/math/primetest"
)

// Key : represents a key used in rsa
type Key struct {
	// public key
	pub uint64
	// private key
	prv uint64

	// n = p*q
	n uint64
}

const (
	// e public key = 3, 17 , or 1<<16 + 1
	e = 17

	// for Miller–Rabin number test
	security = 5

	primeTrail = 100
)

type PrimeComputeError error

func init() { rand.Seed(time.Now().Unix()) }

// GenerateKey : will generate a new rsa public and private key
func GenerateKey() (*Key, error) {
	// 1. choose prime number p and q
	p := genPrime()
	q := genPrime()
	if p == 0 || q == 0 {
		return nil, errors.New("failed to generate prime for keys")
	}

	// 2. compute n=p*q
	n := p * q

	// 3. Geneate phi(n) = (p-1)*(q-1)
	phi := (p - 1) * (q - 1)
	// 4. Choose e from set {2,3,..phi -1}

	// 5. Compute d, such that; e.d = 1 mod phi
	d := uint64(modularinverse.Find(int64(phi), int64(e)))

	return &Key{
		pub: e,
		prv: d,
		n:   n,
	}, nil
}

func Encrypt(x, e, n uint64) (y uint64) {
	// compute y = x^e mod n
	y = fastexp.Find(x, e, n)
	return
}

func (k *Key) Decrypt(y uint64) (x uint64) {
	// compute x = y^d mod n
	x = fastexp.Find(y, k.prv, k.n)
	return
}
func genPrime() uint64 {
	var i int
	var p uint64
	for i = 0; i < primeTrail; i++ {
		p = rand.Uint64() % (1 << 10)
		if p == 1 || (p != 2 && p&1 == 0) {
			i--
		} else {
			if primetest.Check(p, security) {
				break
			}
		}
	}
	if i == primeTrail {
		return 0 // means no prime was found
	}
	return p
}
