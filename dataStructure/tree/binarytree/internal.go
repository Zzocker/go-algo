package binarytree

/*
	binary tree with int data
	- insert : insert data into the tree
	- getNodeCount
	- printValue : from min to max
	- deleteTree
	- isInTree  : returns true if given number exists in tree
	- getHeight : return height of a node
	- getMin : minium value in tree
	- getMax : return max value of the tree
	- isBinaryTree : return true if given tree is binary tree
	- deleteValue
	- getSuccessor() : return next-height value

	internal contains all non-exported function which can as util function for exported functions
*/

type node struct {
	data        int
	left, right *node
}

func deleteTree(root **node) { *root = nil }
func insert(root **node, data int) {
	if *root == nil {
		*root = &node{data: data}
		return
	}
	if (*root).data < data {
		insert(&(*root).right, data)
	} else {
		insert(&(*root).left, data)
	}
}

func sorted(root *node, out *[]int) {
	if root == nil {
		return
	}
	sorted(root.left, out)
	*out = append(*out, root.data)
	sorted(root.right, out)
}

func delete(root **node, data int) {
	if *root == nil {
		return
	}

	if (*root).data < data {
		delete(&(*root).right, data)
	} else if (*root).data > data {
		delete(&(*root).left, data)
	} else {
		// found the node
		// case 1 : node is a leaf
		// case 2 : node has one child
		// case 3 : node is internal node
		if (*root).left == nil && (*root).right == nil {
			*root = nil
		} else if (*root).left == nil {
			*root = (*root).right
		} else if (*root).right == nil {
			*root = (*root).left
		} else {
			rightMin := min((*root).right)
			(*root).data = rightMin
			delete(&(*root).right, rightMin)
		}
	}
}

func isBST(root *node, min, max int) bool {
	if root == nil {
		return true
	}
	if root.data >= min && root.data < max && isBST(root.left, min, root.data) && isBST(root.right, root.data+1, max) {
		return true
	}
	return false
}

func height(root *node) int {
	if root == nil {
		return -1
	}
	leftHeight := height(root.left)
	rightHeight := height(root.right)
	if rightHeight > leftHeight {
		leftHeight = rightHeight
	}
	return 1 + leftHeight
}
func min(root *node) int {
	if root.left == nil {
		return root.data
	}
	return min(root.left)
}

func isInTree(root *node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	} else if root.data < data {
		root = root.right
	} else {
		root = root.left
	}
	return isInTree(root, data)
}

func getNode(root *node, data int) *node {
	if root == nil {
		return nil
	}
	if root.data < data {
		root = root.right
	} else if root.data > data {
		root = root.left
	} else {
		return root
	}
	return getNode(root, data)
}

func successor(root *node, data int) int {
	n := getNode(root, data)
	if n == nil {
		return -1
	}
	if n.right != nil {
		return min(n.right)
	}

	// deapest successor for which the node is at left
	var successor *node
	ancestor := root
	for successor != ancestor {
		if data < ancestor.data {
			successor = ancestor
			ancestor = ancestor.left
		} else if data > ancestor.data {
			ancestor = ancestor.right
		} else {
			break
		}
	}
	if successor == nil {
		return -1
	}
	return successor.data
}
func max(root *node) int {
	if root.right == nil {
		return root.data
	}
	return max(root.right)
}
