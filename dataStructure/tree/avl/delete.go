package avl

func delete(root **node, val int) {
	if *root == nil {
		return
	}
	// for deletion of any node in BST, there are three case
	// 1. Node is leaf
	// 2. Node only had one child
	// 3. Node is intermediate node
	if (*root).data < val {
		delete(&(*root).right, val)
	} else if (*root).data > val {
		delete(&(*root).left, val)
	} else {
		if (*root).left == nil && (*root).right == nil {
			*root = nil
		} else if (*root).right == nil {
			*root = (*root).left
		} else if (*root).left == nil {
			*root = (*root).right
		} else {
			minVal := min((*root).right)
			(*root).data = minVal
			delete(&(*root).right, minVal)
		}
		return
	}
	(*root).height = height(*root)

	if balanceFactor(*root) == 2 { // L
		if balanceFactor((*root).left) == 1 { // LL
			llRotation(root)
		} else if balanceFactor((*root).left) == -1 { // LR
			lrRotation(root)
		} else if balanceFactor((*root).left) == 0 { // LL | LR
			llRotation(root)
		}
	} else if balanceFactor(*root) == -2 { // R
		if balanceFactor((*root).right) == 1 { // RL
			rlRotation(root)
		} else if balanceFactor((*root).right) == -1 { // RR
			rrRotation(root)
		} else if balanceFactor((*root).right) == 0 { // LL | LR
			rrRotation(root)
		}
	}
}

func min(root *node) int {
	if root.left == nil {
		return root.data
	}
	return min(root.left)
}
