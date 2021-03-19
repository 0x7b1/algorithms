package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	if root.Left == nil && root.Right == nil {
		return -1
	}

	left := root.Left.Val
	right := root.Right.Val

	if root.Left.Val == root.Val {
		left = findSecondMinimumValue(root.Left)
	}

	if root.Right.Val == root.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left != -1 && right != -1 {
		return min(left, right)
	} else if left != -1 {
		return left
	} else {
		return right
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	t1 := &TreeNode{2,
		&TreeNode{2, nil, nil},
		&TreeNode{5,
			&TreeNode{5, nil, nil},
			&TreeNode{7, nil, nil}}}

	fmt.Println(findSecondMinimumValue(t1))

	t2 := &TreeNode{5,
		&TreeNode{5, nil, nil},
		&TreeNode{6, nil, nil}}

	fmt.Println(findSecondMinimumValue(t2))

	t3 := &TreeNode{2,
		&TreeNode{2, nil, nil},
		&TreeNode{2, nil, nil}}

	fmt.Println(findSecondMinimumValue(t3))
}
