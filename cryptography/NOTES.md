# Notes Related to Modern cryptography

## Table of Content

- [Essential Number Theory for Public-Key Cryptography](#Essential-Number-Theory-for-Public-Key-Cryptography)

### Essential Number Theory for Public-Key Cryptography
- [Euclidean Algorithm](#Euclidean-Algorithm)
- [Extended Euclidean Algorithm](#Extended-Euclidean-Algorithm)
- [Euler's Phi Function](Euler's-Phi-Function)
- [Fermat’s Little Theorem](#Fermat’s-Little-Theorem)
- [Euler’s Theorem](#Euler’s-Theorem)

### Euclidean Algorithm

Euclidean algorithm is used for calculating ``greatest common divisor(gcd)``, algorithm says,
```
    gcd(r0,r1) = gcd(r1,r0(mod)r1) ; where r0 > r1
```
[Implementation](../math/gcd/gcd.go)

### Extended Euclidean Algorithm
Extended Euclidean Algorithm allows us to compute modular inverse of a number, mathematically we can express

``` 
    gcd(r0,r1) = (S*r0) + (T*r1)
```
Proof
``` 
    we can express, r0 = (q0)*(r1) + r2
    also for index i,
    r(i-2) = S(i-2)*r0 + T(i-2)*r1       ------- (1)
    r(i-1) = S(i-1)*r0 + T(i-1)*r1       ------- (2)

    also,
    r(i) = r(i-2) + q(i-1)*r(i-1)        ------- (3)
    from equation (1),(2) & (3)

    r(i) = [S(i-2) - q(i-1)*S(i-1)]r0 + [T(i-2) - q(i-1)*T(i-1)]r1
    the last iteration will give the value of S and T
    => S(i) = S(i-2) - q(i-1)*S(i-1)
    => T(i) = T(i-2) - q(i-1)*T(i-1)
```

This Algorithm can also be use for finding inverse in Finite fields (Galois Fields)
,r0 in case finite field will be ineducable function P(x), this will give S(x) and T(x)

[Implementation](../math/modularinverse/modularinverse.go)

### Euler's Phi Function
Its is defined as number of element in Ring **Z(m)** {0,1,2,3..m-1} which are co-prime of **m**.
```
    The number of integers in Z(m) relatively prime to m is denoted by Φ (m).
    solution;
        Let,
                m = ((p1)^(e1))*((p2)^(e2))...((pn)^(en)) 
        be a prime factorization of m
        then;
                Φ (m) = ∏ (((pi)^(ei))-((pi)^(ei-1)))
```

Property used in RSA :-

It is easy to compute phi(m) if prime factorization of m is known, and difficult if prime factorization is not know

```
    It is important to stress that we need to know the factorization of m in order to
    calculate Euler’s phi function quickly in this manner. As we will see in the next
    chapter, this property is at the heart of the RSA public-key scheme: Conversely, if
    we know the factorization of a certain number, we can compute Euler’s phi function
    and decrypt the ciphertext. If we do not know the factorization, we cannot compute
    the phi function and, hence, cannot decrypt.
                                                        ---- Understanding Cryptography (Christof Paar and Jan Pelzl)
```

### Fermat’s Little Theorem
```
    Let a be an integer and p be a prime, then:
        a^p ≡ a( mod p)
    also, 
        a^(p-1) = 1 mod p
     => a*a^(p-2) = 1 mod p
     => = a^(-1) = a^(p-2) mod p
```

### Euler’s Theorem
```
    Let a and m be integers with gcd(a, m) = 1, then:
        => a ^(Φ(m)) ≡ 1 ( mod m).
```