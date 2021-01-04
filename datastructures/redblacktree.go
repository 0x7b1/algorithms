package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	p     *Node
	color bool
	key   int
}

type RBTree struct {
	root *Node
}

func (t *RBTree) insert(k int) {
	if t.root == nil {
		t.root = &Node{key: k}
	} else {
		t._insert(t.root, k)
	}
}

func (t *RBTree) _insert(n *Node, k int) *Node {
	// TODO
	return nil
}

func (t *RBTree) printInOrder() {
	t._inOrder(t.root)
}

func (t *RBTree) _inOrder(n *Node) {
	if n != nil {
		t._inOrder(n.left)
		fmt.Println(n.key)
		t._inOrder(n.right)
	}
}

func main() {
	t := new(RBTree)
	t.insert(10)
	t.insert(6)
	t.insert(53)
	t.insert(7)

	t.printInOrder()
}
