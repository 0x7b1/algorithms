package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	small := increasingBST(root.Left)
	large := increasingBST(root.Right)

	root.Left = nil
	root.Right = large

	curr := small
	if curr != nil {
		for curr.Right != nil {
			curr = curr.Right
		}

		curr.Right = root
		return small
	}

	return root
}

func printInOrder(r *TreeNode) {
	if r != nil {
		printInOrder(r.Left)
		fmt.Println(r.Val)
		printInOrder(r.Right)
	}
}

func main() {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n7 := &TreeNode{7, nil, nil}
	n8 := &TreeNode{8, nil, nil}
	n9 := &TreeNode{9, nil, nil}

	n2.Left = n1
	n3.Right = n4
	n3.Left = n2
	n8.Left = n7
	n8.Right = n9
	n6.Right = n8
	n5.Left = n3
	n5.Right = n6

	//n5.PrintInOrder(n5)
	increasingBST(n5)
	printInOrder(n1)
}
