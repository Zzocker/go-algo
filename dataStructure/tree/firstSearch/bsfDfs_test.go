package tree

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFSPreOrder(t *testing.T) {
	got := []int{}

	tree := prepareTestTree()
	preOrder(tree, &got)
	want := []int{10, 8, 6, 7, 9, 5, 11}
	assert.True(t, reflect.DeepEqual(got, want))
}

func TestDFSInOrder(t *testing.T) {
	got := []int{}

	tree := prepareTestTree()
	inOrder(tree, &got)
	want := []int{6, 8, 7, 10, 5, 9, 11}
	assert.True(t, reflect.DeepEqual(got, want))
}

func TestDFSPostOrder(t *testing.T) {
	got := []int{}

	tree := prepareTestTree()
	postOrder(tree, &got)
	want := []int{6, 7, 8, 5, 11, 9, 10}
	assert.True(t, reflect.DeepEqual(got, want))
}

func TestBFS(t *testing.T) {

	tree := prepareTestTree()
	got := bfs(tree)
	want := []int{10, 8, 9, 6, 7, 5, 11}
	assert.True(t, reflect.DeepEqual(got, want))
}
func prepareTestTree() *node {
	/*
	   		  10
	   	8			9
	   6		7	5		11
	*/
	n10 := &node{key: 10}
	n8 := &node{key: 8}
	n9 := &node{key: 9}
	n6 := &node{key: 6}
	n7 := &node{key: 7}
	n5 := &node{key: 5}
	n11 := &node{key: 11}

	n10.left = n8
	n10.right = n9

	n8.left = n6
	n8.right = n7

	n9.left = n5
	n9.right = n11

	return n10
}
