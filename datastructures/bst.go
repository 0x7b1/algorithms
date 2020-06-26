package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	key   int
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(k int) {
	t._insert(t.root, k)
}

func (t *Tree) _insert(n *Node, k int) *Node {
	if t.root == nil {
		t.root = &Node{key: k}
		return t.root
	}

	if n == nil {
		return &Node{key: k}
	}

	if k <= n.key {
		n.left = t._insert(n.left, k)
	} else {
		n.right = t._insert(n.right, k)
	}

	return n
}

func (t *Tree) delete(k int) {
	t._delete(t.root, k)
}

func (t *Tree) _delete(n *Node, k int) {
	if n.key < k {
		t._delete(n.right, k)
	} else if n.key > k {
		t._delete(n.left, k)
	} else {
		if n.left == nil {

		} else if n.right == nil {

		} else {

		}
	}
}

func (t *Tree) transplant(n *Node, u *Node, v *Node) {

}

func (t *Tree) minAbs() int {
	return t._min(t.root)
}

func (t *Tree) maxAbs() int {
	return t._max(t.root)
}

func (t *Tree) _min(n *Node) int {
	if n.left == nil {
		return n.key
	}
	return t._min(n.left)
}

func (t *Tree) _max(n *Node) int {
	if n.right == nil {
		return n.key
	}
	return t._max(n.right)
}

func (t *Tree) find(k int) bool {
	if t._find(t.root, k) == nil {
		return false
	}

	return true
}

func (t *Tree) _find(n *Node, k int) *Node {
	if n == nil {
		return nil
	}

	if n.key == k {
		return n
	}

	if k < n.key {
		return t._find(n.left, k)
	} else {
		return t._find(n.right, k)
	}
}

func (t *Tree) inorder() {
	t._inorder(t.root)
	fmt.Println()
}

func (t *Tree) _inorder(n *Node) {
	if n != nil {
		t._inorder(n.left)
		fmt.Print(n.key, " ")
		t._inorder(n.right)
	}
}

func main() {
	t := new(Tree)

	t.insert(2)
	t.insert(30)
	t.insert(22)
	t.insert(32)
	t.insert(10)

	fmt.Println(t.find(30))
	fmt.Println(t.find(39))

	t.inorder()

	fmt.Println(t.minAbs())
	fmt.Println(t.maxAbs())
}
