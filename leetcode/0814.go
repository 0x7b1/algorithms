package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Val == 0 && root.Left == nil && root.Right == nil {
		root = nil
	}

	return root
}

func printPreorder(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Val)
		printPreorder(t.Left)
		printPreorder(t.Right)
	}
}

func main() {
	t := &TreeNode{
		1,
		nil,
		&TreeNode{
			0,
			&TreeNode{0, nil, nil},
			&TreeNode{1, nil, nil},
		},
	}

	printPreorder(pruneTree(t))
}
