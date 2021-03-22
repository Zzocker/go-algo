package dfk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// ALice's private key
	a := GeneratePrivateKey()

	// Bob's Alice Key
	b := GeneratePrivateKey()

	// Alice generates its public key
	A := GeneratePublicKey(a)

	// Bob generates its public key
	B := GeneratePublicKey(b)

	//Exchange of public key  <---->
	// <------------------------>

	// Alice computes session key
	KAB := GenerateSessionKey(B, a)

	// Bob computes session key
	KBA := GenerateSessionKey(A, b)

	assert.Equalf(t, KBA, KAB, "session key computed by Alice and bob should be same, but found otherwise KbyALice = %d; KbyBob =%d", KAB, KBA)

}
