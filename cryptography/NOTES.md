# Notes Related to Modern cryptography

## Table of Content

- [Essential Number Theory for Public-Key Cryptography](#Essential-Number-Theory-for-Public-Key-Cryptography)
- [Math](#Math)
- [Introduction To Public Key Cryptography](#Introduction-To-Public-Key-Cryptography)
- [Diffie–Hellman key exchange](#Diffie–Hellman-key-exchange)

### Essential Number Theory for Public-Key Cryptography
- [Euclidean Algorithm](#Euclidean-Algorithm)
- [Extended Euclidean Algorithm](#Extended-Euclidean-Algorithm)
- [Euler's Phi Function](Euler's-Phi-Function)
- [Fermat’s Little Theorem](#Fermat’s-Little-Theorem)
- [Euler’s Theorem](#Euler’s-Theorem)
- [Discrete Logarithm Problem](#Discrete-Logarithm-Problem)
- [The Generalized Discrete Logarithm Problem](#The-Generalized-Discrete-Logarithm-Problem)

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
### Discrete Logarithm Problem
```
    Discrete Logarithm Problem (DLP) in Zp
    Given is the finite cyclic group Zp of order p − 1 and a primitive element 
    α ∈ Zp and another element β ∈ Zp . The DLP is the problem
    of determining the integer 1 ≤ x ≤ p − 1 such that:
        α^x ≡ β mod p
```
### The Generalized Discrete Logarithm Problem
```
    Generalized Discrete Logarithm Problem
    Given is a finite cyclic group G with the group operation **o** and
    cardinality n. We consider a primitive element α ∈ G and another
    element β ∈ G. The discrete logarithm problem is finding the 
    integer x, where 1 ≤ x ≤ n, such that:
            β = α o α o . . . o α 
                 ----x times---
```
### Introduction To Public-Key Cryptography

Basic Idea behind PKC, Let say Alice wants to send a message **x** to Bob securely over an unsecured channel.
```
        ---------                                   ---------
        | Alice |                                   | Bob   |
        ---------                                   ---------
            |                                           |
            |                                           |
            |   1. bob send his Public Key (kPub)       |
            |<------------------------------------------|
            |                                           |
            |                                           |
    2. y = enc(x,kPub)                                  |
            |                                           |
            |                                           |
            |   3. Alice send cipher text               |
            |------------------------------------------>|
            |                                           |
            |                                           |
            |                                   4. x = dec(y,kPrv)
            |                                           |
            |                                           |
            |___________________________________________|

    
    x    :         plain text
    y    :         cipher text
    enc  :         encryption function
    dec  :         decryption function
    kPub :         Public Key
    kPrv :         Private Key
```
Unlike in symmetric key cryptography (DES,AES) encryption and decryption is done using a same shared key,but in un-symmetric key cryptography encryption is done using public key and decryption using a private key.

### 3 Type of PKC
1. [RSA](#RSA)

### RSA
RSA (Rivest–Shamir–Adleman) algorithm can divided into two part
#### 1. [Key Generation](#RSA-Key-Generation)
#### 2. [Encryption](#RSA-Encryption)
#### 3. [Decryption](#RSA-Decryption)

#### RSA Key Generation
1. Choose a very large prime number p,q
2. Compute n = p*q
3. Compute phi(n)=(p-1)(q-1)
4. Choose a number **e** from set {2,3,4... phi(n)-1}, such that gcd(e,phi(n)) = 1 (or e and phi(n) should be co-prime) and to ease the computation requirement for encryption, most commonly e = (1<<16)+1 is choose
5. Compute d, such that e.d = 1 mod phi(n)

Now, kPub = (e,n) and kPrv = d

##### RSA Encryption

```
    y = x^e mod n
```

##### RSA Decryption

```
    x = y^d mod n
```
[Implementation](./rsa/rsa.go)

### Math
- [Fast Modular Exponential Computation](#Fast-Modular-Exponential-Computation)
- [Prime Number Test Using Fermat’s Little Theorem](#Prime-Number-Test-Using-Fermat’s-Little-Theorem)
- [Prime Number Test Using Miller–Rabin Theorem](#Prime-Number-Test-Using-Miller–Rabin-Theorem)
- [Groups](#Groups)

### Fast Modular Exponential Computation
Computing ``y = x ^(e) mod n`` for large value of e (say 2^(2024))

```
    express e in binary form
    e = h(n)h(n-1)h(n-2)...h(0)
    INITIALIZE r = x

    for i = n-1 To 1
        r = r*r mod n
        if h(i) == 1:
            r = r*x mod n

    return r
```
[Implementation](../math/fastexp/fastExponential.go)

### Prime Number Test Using Fermat’s Little Theorem
For a prime number **p** and integer **a** in {2,3...p-2}, a^(p-1) = 1 mod p
```
    for i = 1 to S
        choose random a from {2,3...p-2}
        if a^(p-1) != 1:
            return false
    return true
```
S is a security parameter
### Prime Number Test Using Miller–Rabin Theorem
```
        Given the decomposition of an odd prime candidate p̃
                p̃ − 1 = 2^(u)*r
        where r is odd. If we can find an integer a such that
                a^(r) !≡ 1 mod p̃        and        a^(r*2^j) !≡ p̃ − 1 mod p̃
        for all j = {0, 1, . . . , u − 1}, then p̃ is composite. Otherwise, it is probably a prime.
```

Pseudo code
```
    Input   : p (prime candidate) , s (security parameter)
    Output : true (p is likely prime) , false (p is composite)
    INITIALIZE r , u ( p -1 = 2^(u)*r)

    for i = 1 to S: [security loop]
        z = a^r mod p
        if z == 1 or z == p-1:
            goto security loop
        for j=1 to u - 1
            z=(z*z) mod p
            if z == p-1:
                goto security loop
        if z!=p-1:
            return false
    return true
```
[Implementation](../math/primetest/primeTest.go)

### Groups

very engineering definition, a set **G** of element and operation **o** defined on those elements.
```
    Property of a Group
    1. Closeness : 
        for element a , b ∈ G , it holds a o b = c ∈ G
    2. Associative :
        a o (b o c) = (a o b) o c ; a,b,c ∈ G
    3. There exists a identity element 1 ∈ G; such that
        a o 1 = a
    4. There exists inverses of element
        a^(-1) ∈ G
    5. Commutative : (if satisfied then group is called abelian group)
        a o b = b o c
```
```
    Th 1
    The set Zn(star) which consists of all integers i = 0, 1, . . . , n − 1 for which
    gcd(i, n) = 1 forms an abelian group under multiplication modulo n. The identity element is e = 1.
```
```
    Finite Group
    A group (G, ◦) is finite if it has a finite number of elements. We denote the cardinality or order of the group G by |G|.
```
```
    Order of an element
    The order ord(a) of an element a of a group (G, ◦) is the smallest positive integer k such that
        a^(k) = a ◦ a ◦ . . . ◦ a = 1,
                ------K times----
        where 1 is the identity element of G.
```
```
    Cyclic Group
    A group G which contains an element α with maximum order ord( α ) = |G| is said to be cyclic. Elements with maximum order are called primitive elements or  generators.
```
```
    Property of a Finite group
    Let G be a finite group. Then for every a ∈ G it holds that:
        1. a^|G| = 1
        2. ord(a) divides |G|
    
    Property of a cyclic group
    Let G be a finite cyclic group. Then it holds that
        1. The number of primitive elements of G is Φ (|G|).
        2. If |G| is prime, then all elements a  != 1 ∈ G are primitive.
```
### Diffie–Hellman key exchange

```
        ---------                       -----------------                            ---------
        | Alice |                       | Public Domain |                            | Bob   |
        ---------                       -----------------                            ---------
        a = KprA                        α - generator                                 b = kprB
           |                            p - prime number                                  |
           |                                                                              |
           |                                                                              |
        A = α^(a) mod p                                                              B = α^(b) mod p
           |                                                                              |
           |                                                                              |
           |                                   A                                          |
           |----------------------------------------------------------------------------->|
           |                                   B                                          |
           |<-----------------------------------------------------------------------------|
           |                                                                              |
        Kab = B^(a) mod p                                                           Kab = A^(b) mod p
        => Kab = α^(ab) mod p                                                       => Kab = α^(ab) mod p

```