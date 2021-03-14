package modularinverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEEA(t *testing.T) {
	S, T := EEA(973, 301)
	var wantS, wantT int64 = 13, -42
	assert.Equal(t, wantS, S)
	assert.Equal(t, wantT, T)
}

func TestFind(t *testing.T) {
	var want int64 = 6
	assert.Equal(t, want, Find(29, 5)) // 5*6 = 1 mod 29
}
