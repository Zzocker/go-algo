package rsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	is := assert.New(t)
	// 1. Bob generate his keys
	Bob, err := GenerateKey()
	is.NoError(err)

	// 2. bob send his public key
	e, n := Bob.pub, Bob.n

	// 3. alice encrypt the message with bob's public key
	var plaintext uint64 = 569
	ciphertext := Encrypt(plaintext, e, n)
	// t.Log(ciphertext)

	// 4. Bob decrypt the message using his private key
	receivedtext := Bob.Decrypt(ciphertext)

	is.Equalf(receivedtext, plaintext, "actual text was = %d but received = %d", plaintext, receivedtext)

}
