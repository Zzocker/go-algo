package fastexp

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// [x,e,n,ans]
var testCase = [][]uint64{
	{4, 1084, 326, 218},
	{4, 4098, 232, 168},
	{986, (1 << 20) - 1, 1<<10 + 1, 501},
}

func TestFastExp(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for _, v := range testCase {
		got := Find(v[0], v[1], v[2])
		assert.Equalf(t, v[3], Find(v[0], v[1], v[2]), "for [x = %d,e = %d and n = %d] want %d but got %d", v[0], v[1], v[2], v[3], got)
	}
}
