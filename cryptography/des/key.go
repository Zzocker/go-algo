package des

// in des key are required to encrypt and decrypt plaintext
// key provide two function
// 1. KeySchedule : generate 16 subkey required for encryption
// 2. ReverseKeySchedule : genrate 16 subkey required for decryption

// matrix need for in both the process
var (
	// permutation choice - 1
	// drops 8bit from the 64 bit key to give a 56bit key
	pc1Matrix = [56]byte{
		7, 15, 23, 31, 39, 47, 55, 63,
		6, 14, 22, 30, 38, 46, 54, 62,
		5, 13, 21, 29, 37, 45, 53, 61,
		4, 12, 20, 28, 1, 9, 17, 25,
		33, 41, 49, 57, 2, 10, 18, 26,
		34, 42, 50, 58, 3, 11, 19, 27,
		35, 43, 51, 59, 36, 44, 52, 60,
	}

	// permutation choice - 2
	// generates 48 bit key which will be sent to the f-function
	pc2Matrix = [48]byte{
		42, 39, 45, 32, 55, 51,
		53, 28, 41, 50, 35, 46,
		33, 37, 44, 52, 30, 48,
		40, 49, 29, 36, 43, 54,
		15, 4, 25, 19, 9, 1,
		26, 16, 5, 11, 23, 8,
		12, 7, 17, 0, 22, 3,
		10, 14, 6, 20, 27, 24,
	}
	// for i = 0,1,8,15 only one bit is rotated, and two bit for others
	keyLeftRotation = [16]uint8{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}
)

// produce 16 subkey that will be used for encryption[0-15] and decrypotion[15-0]
func keySchedule(key uint64) [16]uint64 {
	// pc1 drop 8bit from the key
	pc1 := permute(key, pc1Matrix[:])

	subkeys := [16]uint64{}
	// left part of 56bits
	c0 := uint32(pc1 >> 28)
	d0 := uint32((pc1 << 4) >> 4)

	lastLeft := c0
	lastRight := d0
	for i := range subkeys {
		lastLeft = ((lastLeft << (4 + keyLeftRotation[i])) >> 4) | (lastLeft<<4)>>(32-keyLeftRotation[i])
		lastRight = ((lastRight << (4 + keyLeftRotation[i])) >> 4) | (lastRight<<4)>>(32-keyLeftRotation[i])
		subkeys[i] = permute(uint64(lastLeft)<<28|uint64(lastRight), pc2Matrix[:])
	}
	return subkeys
}
	