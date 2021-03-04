package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	is := assert.New(t)
	var root *node

	insert(&root, 5)
	is.NotNil(root)
	is.Equal(5, root.data)

	root = nil

	prepare := []int{5, 6, 7, 3, 4, 2, 1}
	for _, v := range prepare {
		insert(&root, v)
	}

	is.Equal(5, root.data)
	is.Equal(6, root.right.data)
	is.Equal(7, root.right.right.data)

	is.Equal(3, root.left.data)
	is.Equal(4, root.left.right.data)
	is.Equal(2, root.left.left.data)
	is.Equal(1, root.left.left.left.data)

}

func TestSorted(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7}
	got := []int{}
	sorted(prepareTest(), &got)
	assert.Equal(t, want, got)
}

func TestDelete(t *testing.T) {
	is := assert.New(t)
	want := []int{1, 2, 4, 5, 6, 7}
	root := prepareTest()

	delete(&root, 3)
	got := []int{}
	sorted(root, &got)

	is.Equal(want, got)
}

func TestIsBST(t *testing.T) {
	is := assert.New(t)
	is.True(isBST(prepareTest(), -((1 << 32) - 1), (1<<32)-11))

	n1 := &node{data: 1}
	n2 := &node{data: 2}
	n3 := &node{data: 3}
	n1.left = n3
	n1.right = n2

	is.False(isBST(n1, -((1 << 32) - 1), (1<<32)-11))

}

func TestHeight(t *testing.T) {
	is := assert.New(t)
	root := prepareTest()

	is.Equal(3, height(root))
	is.Equal(2, height(root.left))
	is.Equal(1, height(root.right))
}

func TestMinMax(t *testing.T) {
	is := assert.New(t)
	root := prepareTest()

	is.Equal(1, min(root))
	is.Equal(7, max(root))
}
func TestIsInTree(t *testing.T) {
	is := assert.New(t)

	root := prepareTest()

	is.True(isInTree(root, 5))
	is.False(isInTree(root, 96))
}

func TestDeleteTree(t *testing.T) {
	root := prepareTest()
	deleteTree(&root)
	assert.Nil(t, root)
}

func TestSuccessor(t *testing.T) {
	root := prepareTest()
	is := assert.New(t)

	is.Equal(5, successor(root, 4))
	is.Equal(-1, successor(root, 9))

}
func prepareTest() *node {
	var root *node
	prepare := []int{5, 6, 7, 3, 4, 2, 1}
	for _, v := range prepare {
		insert(&root, v)
	}
	return root
}
