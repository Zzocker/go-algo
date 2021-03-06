package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLLRotaion(t *testing.T) {
	is := assert.New(t)

	var root *node
	insert(&root, 5)
	insert(&root, 4)
	insert(&root, 3)

	is.Equal(4, root.data)
	is.Equal(2, root.height)

	is.Equal(3, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(5, root.right.data)
	is.Equal(1, root.right.height)
}

func TestLRRotation(t *testing.T) {

	is := assert.New(t)

	var root *node
	insert(&root, 5)
	insert(&root, 3)
	insert(&root, 4)

	is.Equal(4, root.data)
	is.Equal(2, root.height)

	is.Equal(3, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(5, root.right.data)
	is.Equal(1, root.right.height)
}

func TestRRRotation(t *testing.T) {
	is := assert.New(t)

	var root *node

	insert(&root, 3)
	insert(&root, 4)
	insert(&root, 5)

	is.Equal(4, root.data)
	is.Equal(2, root.height)

	is.Equal(3, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(5, root.right.data)
	is.Equal(1, root.right.height)
}

func TestRLRotation(t *testing.T) {
	is := assert.New(t)

	var root *node

	insert(&root, 3)
	insert(&root, 5)
	insert(&root, 4)

	is.Equal(4, root.data)
	is.Equal(2, root.height)

	is.Equal(3, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(5, root.right.data)
	is.Equal(1, root.right.height)
}
