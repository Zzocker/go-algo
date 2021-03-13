// Package gcd : demostrate gcd implementation using Euclidean algorithm
package gcd

// Recursive solve the problem using programme stack
func Recursive(r0, r1 uint64) uint64 {
	if r1 == 0 {
		return r0
	}
	return Recursive(r1, r0%r1)
}

// Iterative solve the problem using loop
func Iterative(r0, r1 uint64) uint64 {
	for r1 != 0 {
		r0, r1 = r1, r0%r1
	}
	return r0
}
