// Package avl is a self balanceing tree
// for more information watch
// https://www.youtube.com/watch?v=jDM6_TnYIqE&t=39s
package avl

type node struct {
	data        int
	left, right *node
	height      int
}

func insert(root **node, data int) {
	if *root == nil {
		*root = &node{data: data, height: 1} // height of newly created node is zero
		// height is number of edges from the root to the deepest leaf
		return
	}

	if (*root).data < data {
		insert(&(*root).right, data)
	} else {
		insert(&(*root).left, data)
	}

	// while returns update height
	(*root).height = height(*root)

	// 4 kind of rotations
	// LL - Rotation
	// LR - Rotation
	// RR - Rotation
	// RL - Rotation

	if balanceFactor(*root) == 2 { // L
		if balanceFactor((*root).left) == 1 { // LL
			llRotation(root)
		} else if balanceFactor((*root).left) == -1 { // LR
			lrRotation(root)
		}
	} else if balanceFactor(*root) == -2 { // R
		if balanceFactor((*root).right) == -1 { // RR
			rrRotation(root)
		} else if balanceFactor((*root).right) == 1 { // RL
			rlRotation(root)
		}
	}
}

func llRotation(root **node) {
	b := (*root).left
	br := (*root).left.right
	(*root).left = br
	b.right = *root

	(*root).height = height(*root)
	b.height = height(b)
	*root = b
}

func lrRotation(root **node) {
	c := (*root).left.right
	cl := c.left
	cr := c.right
	b := (*root).left
	b.right = cl
	(*root).left = cr
	c.left = b
	c.right = (*root)

	// heights
	b.height = height(b)
	(*root).height = height(*root)
	c.height = height(c)
	*root = c
}

func rrRotation(root **node) {
	b := (*root).right
	bl := b.left
	(*root).right = bl
	b.left = *root

	(*root).height = height(*root)
	b.height = height(b)
	*root = b
}

func rlRotation(root **node) {
	c := (*root).right.left
	b := (*root).right
	cl := c.left
	cr := c.right
	(*root).right = cl
	b.left = cr
	c.left = *root
	c.right = b

	// heights
	b.height = height(b)
	(*root).height = height(*root)
	c.height = height(c)
	*root = c
}

// height of left child - height of right child
// 1 : L
// -1 : R
func balanceFactor(root *node) int {
	var hl, hr int
	if root.left != nil {
		hl = root.left.height
	}
	if root.right != nil {
		hr = root.right.height
	}
	return hl - hr
}

func height(root *node) int {
	var hl, hr int
	if root.left != nil {
		hl = root.left.height
	}
	if root.right != nil {
		hr = root.right.height
	}
	if hr > hl {
		hl = hr
	}
	return 1 + hl
}
