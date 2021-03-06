package des

// Expension of the right hand side bit(R_i-1) to get a 48 bit message
// XOR with subkey ki
// feed this 48 bit value to 8 S-Boxes
// combain all the 4 bit out from those S-Boxes to get 32bit
// Permutation

var (
	expensionMatrix = []byte{
		0, 31, 30, 29, 28, 27, 28, 27,
		26, 25, 24, 23, 24, 23, 22, 21,
		20, 19, 20, 19, 18, 17, 16, 15,
		16, 15, 14, 13, 12, 11, 12, 11,
		10, 9, 8, 7, 8, 7, 6, 5,
		4, 3, 4, 3, 2, 1, 0, 3,
	}
	sBox = [8][4][16]uint8{
		{
			{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
			{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
			{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
			{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
		},
		{
			{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
			{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
			{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
			{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
		},
		{
			{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
			{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
			{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
			{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
		},
		{
			{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
			{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
			{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
			{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
		},
		{
			{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
			{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
			{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
			{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
		},
		{
			{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
			{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
			{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
			{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
		},
		{
			{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
			{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
			{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
			{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
		},
		{
			{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
			{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
			{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
			{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
		},
	}
	permutation = [32]byte{
		16, 25, 12, 11, 3, 20, 4, 15,
		31, 17, 9, 6, 27, 14, 1, 22,
		30, 24, 8, 18, 0, 5, 29, 23,
		13, 19, 2, 26, 10, 21, 28, 7,
	}
)

// generate a 32 bit which will XOR with l0
//  takes right side of from feistel network and k1
func fFunction(r0 uint32, k1 uint64) uint32 {
	// expension of r0
	// from 32bit  to 48 bit
	var internal uint64 = permute(uint64(r0), expensionMatrix[:])

	// xor with key
	internal = internal ^ k1
	var out uint32
	// pass it through S-Box
	for i, v := range sBox {
		sixBit := internal & (63 << (6 * i)) >> (6 * i)
		column := sixBit & 30 >> 1
		row := ((sixBit & 1) | (((sixBit) >> 5) << 1))
		out |= uint32(v[row][column] << (4 * i))
	}
	return uint32(permute(uint64(out), permutation[:]))
}

func encrypt(plaintext uint64, key uint64) uint64 {
	// IP not required
	// also IP-1 not required
	// generate subkeys
	subkeys := keySchedule(key)
	l0 := uint32(plaintext >> 32)
	r0 := uint32(plaintext) // will drop last 32 bits
	for _, k := range subkeys {
		// l1,r1 = r0, f-F(r0,k)
		l0, r0 = r0, fFunction(r0, k)
	}
	return uint64(r0) | uint64(l0)<<32
}

func decrypt(cipher uint64, key uint64) uint64 {
	// IP not required
	// also IP-1 not required
	// generate subkeys
	subkeys := keySchedule(key)
	l0 := uint32(cipher >> 32)
	r0 := uint32(cipher) // will drop last 32 bits
	for i := 15; i >= 0; i-- {
		k := subkeys[i]
		// l1,r1 = r0, f-F(r0,k)
		l0, r0 = r0, fFunction(r0, k)
	}
	return uint64(r0) | uint64(l0)<<32
}
