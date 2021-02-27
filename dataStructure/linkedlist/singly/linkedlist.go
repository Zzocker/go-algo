// Package singly : implements both singly linkedList
package singly

// Node : is a node of the linked list
type Node struct {
	Value interface{}
	next  *Node
}

// Next : gives node next to this current node
func (n *Node) Next() *Node {
	return n.next
}

// List represents singly linked List
type List struct {
	root Node
	len  int
}

// New create a empty List
func New() List {
	return List{}
}

// NewNode : create new node with given value
func NewNode(val interface{}) *Node {
	return &Node{Value: val}
}

// Size : current size of the linkedList
func (l *List) Size() int { return l.len }

// Empty : returns boolean value represent whether list is empty or not
func (l *List) Empty() bool { return l.len == 0 }

// NodeAt : return ith node (for 0 indexing)
// nil is returns in case node is not present
func (l *List) NodeAt(i int) *Node {
	if i < 0 || i >= l.Size() {
		return nil
	}
	var head *Node = l.root.next
	for j := 0; j < i; j, head = j+1, head.Next() {
	}
	return head
}

// PushFront : will add given node at start of the list
func (l *List) PushFront(n *Node) {
	n.next = l.root.next
	l.root.next = n
	l.len++
}

// PopFront : removes first element and return value of that node
// return nil can case of empty list
func (l *List) PopFront() interface{} {
	if l.Size() == 0 {
		return nil
	}
	out := l.root.next.Value
	l.root.next = l.root.Next().Next()
	l.len--
	return out
}

// PushBack : will add given node at end of the list
func (l *List) PushBack(n *Node) {
	if l.Size() == 0 {
		l.root.next = n
	} else {
		l.NodeAt(l.Size() - 1).next = n
	}
	l.len++
}

// PopBack : removes end element and return value of that node
// return nil can case of empty list
func (l *List) PopBack() interface{} {
	if l.Size() == 0 {
		return nil
	}
	var out interface{}
	if l.Size() == 1 {
		out = l.root.next.Value
		l.root.next = nil
	} else {
		secondLast := l.NodeAt(l.Size() - 2)
		out = secondLast.next.Value
		secondLast.next = nil
	}
	l.len--
	return out
}

// Front : returns front node
func (l *List) Front() *Node {
	return l.root.Next()
}

// Back : returns back node of the list
func (l *List) Back() *Node {
	return l.NodeAt(l.Size() - 1)
}

// Insert : inserts a given node at index
func (l *List) Insert(index int, n *Node) {
	if index < 0 || index >= l.Size() {
		return
	}

	if index == 0 {
		l.PushFront(n)
	} else {
		prev := l.NodeAt(index - 1)
		n.next = prev.next
		prev.next = n
		l.len++
	}

}

// Erase : remove the at index
func (l *List) Erase(index int) {
	if index < 0 || index >= l.Size() {
		return
	}
	if index == 0 {
		l.PopFront()
	} else {
		preNode := l.NodeAt(index - 1)
		preNode.next = preNode.next.next
		l.len--
	}
}

// FromEnd : return ith node from end (1-indexed)
func (l *List) FromEnd(i int) *Node {
	if i < 1 || i > l.Size() {
		return nil
	}
	return l.NodeAt(l.Size() - i)
}

// Reverse : will reverse list
func (l *List) Reverse() {
	var newHead *Node
	var ptr *Node = l.root.next
	var next *Node
	for ptr != nil {
		next = ptr.next
		ptr.next = newHead
		newHead = ptr
		ptr = next
	}
	l.root.next = newHead
}
