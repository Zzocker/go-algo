package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteLLRotation(t *testing.T) {
	is := assert.New(t)

	var root *node

	insert(&root, 5)
	insert(&root, 4)
	insert(&root, 3)
	insert(&root, 2)

	delete(&root, 5)

	is.Equal(3, root.data)
	is.Equal(2, root.height)

	is.Equal(2, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(4, root.right.data)
	is.Equal(1, root.right.height)
}


func TestDeleteLRRotation(t *testing.T) {
	is := assert.New(t)

	var root *node

	insert(&root, 10)
	insert(&root, 8)
	insert(&root, 6)
	insert(&root, 7)

	delete(&root, 10)

	is.Equal(7, root.data)
	is.Equal(2, root.height)

	is.Equal(6, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(8, root.right.data)
	is.Equal(1, root.right.height)
}


func TestDeleteRRRotation(t *testing.T) {
	is := assert.New(t)

	var root *node

	insert(&root, 2)
	insert(&root, 3)
	insert(&root, 4)
	insert(&root, 5)

	delete(&root, 2)

	is.Equal(4, root.data)
	is.Equal(2, root.height)

	is.Equal(3, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(5, root.right.data)
	is.Equal(1, root.right.height)
}

func TestDeleteRLRotation(t *testing.T) {
	is := assert.New(t)

	var root *node

	insert(&root, 7)
	insert(&root, 6)
	insert(&root, 9)
	insert(&root, 8)

	delete(&root, 6)

	is.Equal(8, root.data)
	is.Equal(2, root.height)

	is.Equal(7, root.left.data)
	is.Equal(1, root.left.height)

	is.Equal(9, root.right.data)
	is.Equal(1, root.right.height)
}