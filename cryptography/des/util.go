package des

// will be used for permute using various matrix
func permute(in uint64, matrix []uint8) uint64 {
	var block uint64
	for i, v := range matrix {
		block |= ((in >> v) & 1) << (len(matrix) - 1 - i)
	}
	return block
}

