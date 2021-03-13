package gcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	var r0 uint64 = 7
	var r1 uint64 = 2
	var want uint64 = 1
	assert.Equal(t, want, Recursive(r0, r1))
	assert.Equal(t, want, Iterative(r0, r1))
	r0 = 9
	r1 = 3
	want = 3
	assert.Equal(t, want, Recursive(r0, r1))
	assert.Equal(t, want, Iterative(r0, r1))
}
