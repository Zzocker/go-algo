package des

import (
	"testing"
)

func TestKeySchedule(t *testing.T) {
	var plaintext uint64 = 56
	var key uint64 = 43
	cipher := encrypt(plaintext, key)
	decipher := decrypt(cipher, key)
	_ = decipher
}
