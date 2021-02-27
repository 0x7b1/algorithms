package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func depth(sum *int, root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(sum, root.Left)
	right := depth(sum, root.Right)

	*sum = max(*sum, left+right)

	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	sum := 0
	depth(&sum, root)
	return sum
}

func main() {
	t := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil}}
	fmt.Println(diameterOfBinaryTree(t))
}
