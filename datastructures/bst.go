package main

import (
	"fmt"
	"log"
)

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	key    int
}

type BST struct {
	root *Node
}

func (b *BST) Insert(k int) {
	if b.root == nil {
		b.root = &Node{key: k}
	} else {
		b.insert(b.root, b.root, k)
	}
}

func (b *BST) insert(n *Node, p *Node, k int) *Node {
	if n == nil {
		return &Node{key: k, parent: p}
	}

	if k <= n.key {
		n.left = b.insert(n.left, n, k)
	} else {
		n.right = b.insert(n.right, n, k)
	}

	return n
}

func (b *BST) Delete(k int) {
	b.delete(b.root, k)
}

func (b *BST) delete(n *Node, k int) {
	if n.left == nil {
		b.transplant(n, n.right)
	} else if n.right == nil {
		b.transplant(n, n.left)
	} else {
		y := b.min(n.right)
		if y.parent != n {
			b.transplant(y, y.right)
			y.right = n.right
			y.right.parent = y
		}

		b.transplant(n, y)
		y.left = n.left
		y.left.parent = y
	}
}

func (b *BST) transplant(u, v *Node) {
	if u.parent == nil {
		b.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func (b *BST) Min() int {
	return b.min(b.root).key
}

func (b *BST) min(n *Node) *Node {
	if n.left != nil {
		return b.min(n.left)
	} else {
		return n
	}
}

func (b *BST) Max() int {
	return b.max(b.root).key
}

func (b *BST) max(n *Node) *Node {
	if n.right != nil {
		return b.max(n.right)
	} else {
		return n
	}
}

func (b *BST) Exist(k int) bool {
	if b.search(b.root, k) == nil {
		return false
	}
	return true
}

func (b *BST) search(n *Node, k int) *Node {
	if n == nil {
		return nil
	}

	if n.key == k {
		return n
	} else if k <= n.key {
		return b.search(n.left, k)
	} else {
		return b.search(n.right, k)
	}
}

func (b *BST) Predecessor(k int) int {
	n := b.search(b.root, k)

	return n.key
}

func (b *BST) Successor(k int) int {
	n := b.search(b.root, k)

	if n == nil {
		log.Fatal("Provided key does not exist")
		return -1
	}

	if n.right == nil {
		return b.Successor(n.right.key)
	}

	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = y.parent
	}

	return y.key
}

func (b *BST) PrintInOrder() {
	b.printInOrder(b.root)
}

func (b *BST) printInOrder(n *Node) {
	if n != nil {
		b.printInOrder(n.left)
		fmt.Print(n.key, " ")
		b.printInOrder(n.right)
	}
}

func main() {
	bst := new(BST)
	bst.Insert(2)
	bst.Insert(30)
	bst.Insert(14)
	bst.Insert(5)
	bst.Insert(17)

	bst.PrintInOrder()

	fmt.Println("\nMin: ", bst.Min(), "Max: ", bst.Max())

	fmt.Println("Exists 30 =", bst.Exist(30))
	fmt.Println("Exists 15 =", bst.Exist(15))

	fmt.Println("Predecessor 30 =", bst.Predecessor(30))
	fmt.Println("Successor 14 =", bst.Successor(14))

	bst.Delete(14)
	bst.PrintInOrder()
}
