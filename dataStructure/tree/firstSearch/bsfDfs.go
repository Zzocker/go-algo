package tree

import "github.com/Zzocker/go-algo/dataStructure/queue"

type node struct {
	key   int
	left  *node
	right *node
}

/*
	DFS (Depth First Search)
		1.PreOrder (self,left,right)
		2.InOrder (left,self,right)
		3.PostOrder (left,right,self)
	BFS (Breadth First Search)
	more info https://www.youtube.com/watch?v=uWL6FJhq5fM
*/

// DFS

func preOrder(tree *node, out *[]int) {
	if tree == nil {
		return
	}
	*out = append(*out, tree.key)
	preOrder(tree.left, out)
	preOrder(tree.right, out)
}

func inOrder(tree *node, out *[]int) {
	if tree == nil {
		return
	}
	inOrder(tree.left, out)
	*out = append(*out, tree.key)
	inOrder(tree.right, out)
}

func postOrder(tree *node, out *[]int) {
	if tree == nil {
		return
	}
	postOrder(tree.left, out)
	postOrder(tree.right, out)
	*out = append(*out, tree.key)
}

func bfs(tree *node) []int {
	q := queue.NewUnbound()
	var out []int
	q.Push(tree)
	for !q.Empty() {
		n := q.Pop().(*node)
		out = append(out, n.key)
		if n.left != nil {
			q.Push(n.left)
		}
		if n.right != nil {
			q.Push(n.right)
		}
	}
	return out
}
