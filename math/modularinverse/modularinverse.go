// package modularinverse demonastrate Extended Euclidean Algorithm
// for finding modular inverse of number
package modularinverse

// Find : will give modular inverse of r1
// r1 * inverse(r1) = 1 mod r0
// from EEA, gcd(r0,r1) = s*r0 + t*r1
// taking mod r0 both side and inverse exists iff gcd(r0,r1) == 1 (i.e r0 and r1 are co-prime)
// => 1 mode r0 = 0 + t*r1
// => modular inverse of r1 = t
func Find(r0, r1 int64) int64 {
	// check if r0 and r1 are co-prime, using math.GCD(r1,r0) == 1
	_, t := EEA(r0, r1)
	return (t + r0) % r0 // if value of t is -ve
}

// EEA : Extended Euclidean Algorithm
// gcd(r0,r1) = s*r0 + t*r1
// returns : s,t
func EEA(r0, r1 int64) (int64, int64) {
	// initial values for s and t
	var s0, t0, s1, t1 int64 = 1, 0, 0, 1
	for r1 != 0 {
		ri := r0 % r1
		q := (r0 - ri) / r1
		s0, s1 = s1, s0-(q*s1)
		t0, t1 = t1, t0-(q*t1)

		r0, r1 = r1, ri
	}
	return s0, t0
}
