package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minHeight(root.Right) + 1
	}

	if root.Right == nil {
		return minHeight(root.Left) + 1
	}

	left := minHeight(root.Left)
	right := minHeight(root.Right)

	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxHeight(root.Left)
	right := maxHeight(root.Right)

	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := maxHeight(root.Left)
	right := maxHeight(root.Right)

	return abs(left-right) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	tree := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	fmt.Println(isBalanced(tree))
}
